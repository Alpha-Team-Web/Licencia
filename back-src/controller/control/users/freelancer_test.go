package users

import (
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"testing"
)

/*{
"username":"ashkan",
"password": "a12345",
"firstname": "ashkan",
"lastname": "ashkan",
"email": "aaaaa@gmail.com"
"website": "github.com"
}*/

func TestUpdateFreelancer(t *testing.T) {
	db := database.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	frl := existence.Freelancer{}
	frl.Username = "ashkan"
	frl.Password = "fjfjfj"
	frl.FirstName = "ashkan"
	frl.LastName = "ashkan"
	frl.Email = "bbbb@gmail.com"
	if err := db.UpdateFreelancerProfile(frl.Username, frl); err != nil {
		t.Error(err)
	}
	frl2, err := db.GetFreelancer("ashkan")
	if err != nil {
		t.Error(err)
	}
	if frl2.Email != "bbbb@gmail.com" {
		t.Errorf("%s %v", "Fail : ", frl2)
	}
	if frl2.Password == "fjfjfj" {
		t.Errorf("%s %v", "Fail : ", frl2)
	}
	if err := db.UpdateFreelancerPassword(frl.Username, "dasdsa", "sadasdas"); err == nil {
		t.Error("Old pass must be the same. Fail.")
	}
	if err := db.UpdateFreelancerPassword(frl.Username, "a12345", "sadasdas"); err != nil {
		t.Error(err)
	}
	frl3, err := db.GetFreelancer("ashkan")
	if frl3.Password != "sadasdas" {
		t.Error("%s %v", "Fail : ", frl3)
	}
	frl.Website = "ffff.com"
	if err := db.UpdateFreelancerLinks("ashkan", frl); err != nil {
		t.Error(err)
	}
	frl4, err := db.GetFreelancer("ashkan")
	if frl4.Website != frl.Website {
		t.Errorf("%s %v", err, frl4)
	}
}

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
}
