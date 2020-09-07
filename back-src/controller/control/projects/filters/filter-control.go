package filters

import (
	"back-src/controller/utils/libs/sets"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"math"
)

var Inv invertedEngine

func Filter(filter data.Filter, db *database.Database) ([]notifications.ListicProject, error) {
	if resultSet, err := filterByPriceAndStat(filter, db); err == nil {
		if filter.IsFilterBySkill {
			resultSet = sets.IntersectSets(resultSet, filterBySkills(filter))
		}
		return getListicProjectsByIds(resultSet.GetMembers(), db), nil
	} else {
		return []notifications.ListicProject{}, err
	}
}

func getListicProjectsByIds(ids []string, db *database.Database) []notifications.ListicProject {
	listicProjects := []notifications.ListicProject{}
	for _, id := range ids {
		if project, err := db.ProjectTable.GetProjectDefinedColumns(id, "id", "name", "description", "start_date", "employer_username", "freelancer_requests_with_description", "fields_with_skills"); err == nil {
			listicProjects = append(listicProjects, getListicProjectFromProject(project, db))
		}
	}
	return listicProjects
}

func getListicProjectFromProject(project existence.Project, db *database.Database) notifications.ListicProject {
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

func filterByPriceAndStat(filter data.Filter, db *database.Database) (sets.Set, error) {
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
