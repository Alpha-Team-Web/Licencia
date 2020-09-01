package libs

func Xor(bool1, bool2 bool) bool {
	if (bool1 && bool2) || (!bool1 && !bool2) {
		return false
	}
	return true
}

func XNor(bool1, bool2 bool) bool {
	return !Xor(bool1, bool2)
}
