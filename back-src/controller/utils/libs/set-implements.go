package libs

type void struct{}

type Set struct {
	member void
	keys   map[string]void
}

type SetImplementation interface {
	Add(element string) SetImplementation
	Delete(element string) SetImplementation
	AddAll(elements SetImplementation) SetImplementation
	DeleteAll(elements SetImplementation) SetImplementation
	GetMemebers() []string
}

func (set *Set) Add(element string) *Set {
	set.keys[element] = set.member
	return set
}

func (set *Set) Delete(element string) *Set {
	delete(set.keys, element)
	return set
}

func (set *Set) AddAll(elements Set) *Set {
	for s, _ := range elements.keys {
		set.Add(s)
	}
	return set
}

func (set *Set) DeleteAll(elements Set) *Set {
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
