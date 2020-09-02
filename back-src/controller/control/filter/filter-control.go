package filter

import (
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs/sets"
	"back-src/model/database"
	"back-src/model/existence"
)

type invertedIndex struct {
	invertedMap map[string]sets.Set //skills -> set of projectIds
}

var inv = &invertedIndex{map[string]sets.Set{}}

func (inv *invertedIndex) addKeyToMap(skill string, projectId string) {
	if set, ok := inv.invertedMap[skill]; ok {
		set.Add(projectId)
	} else {
		set2 := sets.NewSet(projectId)
		inv.invertedMap[skill] = set2
	}
}

func (inv *invertedIndex) removeKey(skill string) {
	delete(inv.invertedMap, skill)
}

func AddSkillToProject(skill string, projectId string) {
	inv.addKeyToMap(skill, projectId)
}

func Filter(filter data.Filter, db *database.Database) ([]existence.ListicProject, error) {
	return nil, nil
}
