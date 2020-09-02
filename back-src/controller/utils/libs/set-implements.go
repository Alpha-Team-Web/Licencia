package libs

type void struct{}

type Set struct {
	member void
	keys   map[string]void
}

func (set *Set) Add(element string) *Set {
	set.keys[element] = set.member
	return set
}

func (set *Set) Delete(element string) *Set {
	delete(set.keys, element)
	return set
}

func (set *Set) UnionWith(elements *Set) *Set {
	for s, _ := range elements.keys {
		set.Add(s)
	}
	return set
}

func (set *Set) SubtractFrom(elements *Set) *Set {
	for s, _ := range elements.keys {
		set.Delete(s)
	}
	return set
}

func (set *Set) GetMembers() []string {
	members := make([]string, 0)
	for s, _ := range set.keys {
		members = append(members, s)
	}
	return members
}

func (set *Set) AddAll(members ...string) {
	for _, s := range members {
		set.Add(s)
	}
}

func IntersectSets(sets ...Set) *Set {
	mother := Set{}
	for _, s := range sets {
		mother.UnionWith(&s)
	}
	complement := &Set{}
	for _, s := range sets {
		temp := mother
		complement.UnionWith(temp.SubtractFrom(&s))
	}
	return mother.SubtractFrom(complement)
}
