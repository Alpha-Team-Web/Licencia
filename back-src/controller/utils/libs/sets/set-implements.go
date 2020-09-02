package sets

func IntersectSets(sets ...Set) Set {
	mother := NewSet()
	for _, s := range sets {
		mother.UnionWith(s)
	}
	complement := NewSet()
	for _, s := range sets {
		temp := NewSet(mother.GetMembers()...)
		complement.UnionWith(temp.SubtractFrom(s))
	}
	return mother.SubtractFrom(complement)
}
