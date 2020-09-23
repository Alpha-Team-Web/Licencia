package filters

import (
	"back-src/controller/utils/libs"
	"back-src/controller/utils/libs/sets"
	"back-src/model"
	"back-src/model/existence"
	"back-src/model/sql"
	"back-src/view/data"
	"back-src/view/notifications"
	"math"
)

var Inv invertedEngine

const PageSize = 20

func Filter(auth existence.AuthToken, filter data.Filter, dbApi model.DbApi) ([]notifications.ListicProject, error) {
	projectIds := []string{}
	userWithRole := libs.Ternary(auth.IsFreelancer, "frl-"+auth.Username, "emp-"+auth.Username).(string)
	//redis
	if dbApi.RedisDb.FilterDb.IsThereFilter(userWithRole) {
		if redisFilter, projectIds, err := dbApi.RedisDb.FilterDb.GetFilter(userWithRole); err == nil {
			if areFiltersEqual(filter, redisFilter) {
				dbApi.RedisDb.FilterDb.ExtendFilterExpiry(userWithRole)
				return getListicProjectsByIds(projectIds, filter.PageNumber, dbApi.SqlDb), nil
			}
		}
	}
	//sql
	if resultSet, err := filterByPriceAndStat(filter, dbApi.SqlDb); err == nil {
		if filter.IsFilterBySkill {
			resultSet = sets.IntersectSets(resultSet, filterBySkills(filter))
		}
		projectIds = append(projectIds, resultSet.GetMembers()...)
		dbApi.RedisDb.FilterDb.AddFilterToUserWithRole(
			userWithRole,
			filter,
			projectIds,
		)
	} else {
		return []notifications.ListicProject{}, err
	}
	return getListicProjectsByIds(projectIds, filter.PageNumber, dbApi.SqlDb), nil
}

func areFiltersEqual(first data.Filter, second data.Filter) bool {
	result := true
	result = result && (first.Status == second.Status)
	result = result && (first.MinPrice == second.MinPrice)
	result = result && (first.MaxPrice == second.MaxPrice)
	result = result && (first.IsFilterBySkill == second.IsFilterBySkill)
	result = result && libs.AreStringSetsEqual(first.MustIncludeSkills, second.MustIncludeSkills)
	result = result && libs.AreStringSetsEqual(first.IncludeSkills, second.IncludeSkills)
	result = result && libs.AreStringSetsEqual(first.ExcludeSkills, second.ExcludeSkills)
	return result
}

func getListicProjectsByIds(ids []string, pageNumber int, db *sql.Database) []notifications.ListicProject {
	listicProjects := []notifications.ListicProject{}
	start, end := (pageNumber-1)*PageSize, pageNumber*PageSize-1

	for i := start; i < end && i < len(ids); i++ {
		if project, err := db.ProjectTable.GetProjectDefinedColumns(ids[i], "id", "name", "description", "start_date", "employer_username", "freelancer_requests_with_description", "fields_with_skills"); err == nil {
			listicProjects = append(listicProjects, getListicProjectFromProject(project, db))
		}
	}
	return listicProjects
}

func getListicProjectFromProject(project existence.Project, db *sql.Database) notifications.ListicProject {
	listicProject := notifications.ListicProject{
		Id:                  project.Id,
		Name:                project.Name,
		Description:         project.Description,
		EmployerUsername:    project.EmployerUsername,
		StartDate:           project.StartDate,
		NumberOfSuggestions: len(project.FreelancerRequestsWithDescription),
		Skills:              []string{},
	}
	for _, skills := range project.FieldsWithSkills {
		listicProject.Skills = append(listicProject.Skills, skills...)
	}
	if shownName, err := db.EmployerTable.GetEmployerShownNameByUsername(project.EmployerUsername); err == nil {
		listicProject.EmployerShownName = shownName
	}
	return listicProject
}

func filterByPriceAndStat(filter data.Filter, db *sql.Database) (sets.Set, error) {
	max := filter.MaxPrice
	min := filter.MinPrice
	if max == 0 {
		max = math.MaxFloat64
	}
	if max < min {
		return sets.NewSet(), nil
	}
	if ids, err := db.ProjectTable.GetProjectIdsByStatusAndMaxBudget(filter.Status, max, min); err != nil {
		return sets.NewSet(), err
	} else {
		return sets.NewSet(ids...), nil
	}
}

func filterBySkills(filter data.Filter) sets.Set {
	set := filterByMustInclude(filter.MustIncludeSkills).UnionWith(filterByIncludes(filter.IncludeSkills))
	filterByExcludes(set, filter.ExcludeSkills)
	return set
}

func filterByMustInclude(mustIncludes []string) sets.Set {
	var resultSets []sets.Set
	for _, include := range mustIncludes {
		resultSets = append(resultSets, Inv.invertedMap[include])
	}
	return sets.IntersectSets(resultSets...)
}

func filterByIncludes(includes []string) sets.Set {
	var set = sets.NewSet()
	for _, include := range includes {
		set.UnionWith(Inv.invertedMap[include])
	}
	return set
}

func filterByExcludes(set sets.Set, excludes []string) {
	for _, exclude := range excludes {
		set.SubtractFrom(Inv.invertedMap[exclude])
	}
}
