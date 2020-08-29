package users

import (
	"back-src/controller/control/utils/libs"
	"back-src/model/database"
	"fmt"
	"testing"
)

func TestChooseFreelancerSkills(t *testing.T) {
	db := database.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	frlSkills := []string{"CSharp", "C", "Django", "Flask", "Cpp"}
	if err := ChooseFreelancerSkills("ashkan", "1", frlSkills, db); err != nil {
		t.Error(err)
	}
	skills, err := db.GetFieldSkills("1")
	if err != nil {
		t.Error(err)
	}
	newSkills := []string{"ASP.NET", "CSharp", "Django", "Flask", "Go", "JavaScript", "Python", "Ruby", "React", "html", "css", "C", "Cpp"}
	if !libs.AreStringSetsEqual(skills, newSkills) {
		t.Errorf("%v %v", skills, newSkills)
	}
	if frl, err := db.GetFreelancerByUsername("ashkan"); err != nil {
		t.Error(err)
	} else {
		if !libs.AreStringSetsEqual(frl.ChosenFieldWithSkills["1"], frlSkills) {
			t.Errorf("%v %v", skills, newSkills)
		}
	}
	fmt.Println("6")
}
