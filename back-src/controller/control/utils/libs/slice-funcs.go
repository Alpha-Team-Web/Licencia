package libs

func ContainsString(fieldSkills []string, skill string) bool {
	for _, fieldSkill := range fieldSkills {
		if fieldSkill == skill {
			return true
		}
	}
	return false
}
