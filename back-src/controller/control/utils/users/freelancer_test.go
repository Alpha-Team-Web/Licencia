package users

import (
	"back-src/model/database"
	"testing"
)

func TestChooseFreelancerSkills(t *testing.T) {
	db := database.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	if err := ChooseFreelancerSkills("ashkan", "1", []string{"CSharp", "C", "Django", "Flask", "Cpp"}, db); err != nil {
		t.Error(err)
	}
	skills, err := db.GetFieldSkills("1")
	if err != nil {
		t.Error(err)
	}

}
