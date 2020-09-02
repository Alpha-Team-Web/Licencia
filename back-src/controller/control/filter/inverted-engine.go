package filter

import "back-src/controller/utils/libs/sets"

type invertedEngine struct {
	invertedMap map[string]sets.Set //skills -> Set of projectIds
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
	inv.addKeyToMap(skill, projectId)
}
