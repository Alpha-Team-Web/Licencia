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

func RemoveStringElement(slice []string, index int) []string {
	return append(slice[index:], slice[index+1:]...)
}

func GetKeySet(givenMap map[string]string) []string {
	keys := make([]string, 0, len(givenMap))
	for k := range givenMap {
		keys = append(keys, k)
	}
	return keys
}

func ContainsKey(givenMap map[string]string, word string) bool {
	for key, _ := range givenMap {
		if key == word {
			return true
		}
	}
	return false
}
