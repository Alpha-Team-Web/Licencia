package libs

func ContainsString(fieldSkills []string, skill string) bool {
	for _, fieldSkill := range fieldSkills {
		if fieldSkill == skill {
			return true
		}
	}
	return false
}

func AreStringSetsEqual(set1 []string, set2 []string) (boolean bool) {
	boolean = false
	if len(set1) == len(set2) {
		for _, str := range set1 {
			if !ContainsString(set2, str) {
				return
			}
		}
		boolean = true
	}
	return
}
