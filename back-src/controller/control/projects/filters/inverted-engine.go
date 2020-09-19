package filters

import (
	"back-src/controller/utils/libs/sets"
	"back-src/model/sql"
)

type invertedEngine struct {
	invertedMap map[string]sets.Set //skills -> Set of projectIds
	FailedInit  bool
}

func NewEngine(db *sql.Database) invertedEngine {
	inv := invertedEngine{invertedMap: map[string]sets.Set{}, FailedInit: false}

	projects, err := db.ProjectTable.GetAllProjects()
	if err != nil {
		inv.FailedInit = true
	}

	for _, project := range projects {
		for _, skills := range project.FieldsWithSkills {
			for _, skill := range skills {
				inv.addKeyToMap(skill, project.Id)
			}
		}
	}
	return inv
}

func (inv *invertedEngine) addKeyToMap(skill string, projectId string) {
	if set, ok := inv.invertedMap[skill]; ok {
		set.Add(projectId)
	} else {
		set2 := sets.NewSet(projectId)
		inv.invertedMap[skill] = set2
	}
}

func (inv *invertedEngine) removeKey(skill string) {
	delete(inv.invertedMap, skill)
}

func AddSkillToProject(skill string, projectId string) {
	Inv.addKeyToMap(skill, projectId)
}
