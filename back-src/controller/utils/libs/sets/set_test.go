package sets

import (
	"back-src/controller/utils/libs"
	"testing"
)

func TestMainMethods(t *testing.T) {
	set1 := NewSet("shit", "hello", "sex")
	if !set1.Equals(NewSet("sex", "hello", "shit")) {
		t.Error()
	}
	set2 := NewSet()
	set2.Add("oh").AddAll("God", "Loves").RemoveAll("Oh", "oh", "fuck").UnionWith(NewSet("Sex"))
	if !libs.AreStringSetsEqual([]string{"God", "Loves", "Sex"}, set2.GetMembers()) {
		t.Error()
	}
	set3 := NewSet("Girl", "Sky", "On", "No", "Train")
	set3.Add("Shit").SubtractFrom(NewSet("Shit", "Sky", "No")).Add("The")
	if !libs.AreStringSetsEqual([]string{"Girl", "On", "The", "Train"}, set3.GetMembers()) {
		t.Error()
	}
}

func TestIntersectSets(t *testing.T) {
	set1 := NewSet("Killing", "Monsters", "Geralt", "Ciri")
	set2 := NewSet("Ciri", "Having", "Sex", "Geralt", "Yen", "Killing")
	set3 := NewSet("Geralt", "Triss", "Killing", "Ciri")
	set4 := NewSet("Geralt", "Triss", "Ciri")
	if !IntersectSets(set1, set2, set3).Equals(NewSet(set3.GetMembers()...).Remove("Triss")) {
		t.Errorf("%v, %v", IntersectSets(set1, set2, set3), NewSet(set2.GetMembers()...).Remove("Triss"))
	}
	if !IntersectSets(set1, set2, set3, set4).Equals(NewSet(set4.GetMembers()...).Remove("Triss")) {
		t.Errorf("%v, %v", IntersectSets(set1, set2, set3), NewSet(set2.GetMembers()...).Remove("Triss"))
	}
	if !IntersectSets(set1).Equals(set1) {
		t.Errorf("%v", IntersectSets(set1))
	}
	if !IntersectSets().Equals(NewSet()) {
		t.Error()
	}
}
