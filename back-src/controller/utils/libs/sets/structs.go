package sets

import "back-src/controller/utils/libs"

type void struct{}

type Set struct {
	member void
	keys   map[string]void
}

/*
set3: 223
map: 110
*/

/*
set4: 566
map: 110
*/

func NewSet(firstElements ...string) Set {
	set := Set{keys: map[string]void{}}
	set.AddAll(firstElements...)
	return set
}

func (set Set) Add(element string) Set {
	set.keys[element] = set.member
	return set
}

func (set Set) Remove(element string) Set {
	delete(set.keys, element)
	return set
}

func (set Set) AddAll(members ...string) Set {
	for _, s := range members {
		set.Add(s)
	}
	return set
}

func (set Set) RemoveAll(members ...string) Set {
	for _, s := range members {
		set.Remove(s)
	}
	return set
}

func (set Set) UnionWith(elements Set) Set {
	for s, _ := range elements.keys {
		set.Add(s)
	}
	return set
}

func (set Set) SubtractFrom(elements Set) Set {
	for s, _ := range elements.keys {
		set.Remove(s)
	}
	return set
}

func (set Set) GetMembers() []string {
	members := make([]string, 0)
	for s, _ := range set.keys {
		members = append(members, s)
	}
	return members
}

func (set Set) Equals(set2 Set) bool {
	return libs.AreStringSetsEqual(set.GetMembers(), set2.GetMembers())
}
