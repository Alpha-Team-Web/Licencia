package users

import (
	"back-src/model/database"
	"back-src/model/existence"
	"testing"
)

func TestUpdateEmployer(t *testing.T) {
	db := database.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	emp := existence.Employer{}
	emp.Username = "ashkan"
	emp.Password = "fjfjfj"
	emp.FirstName = "ashkan"
	emp.LastName = "ashkan"
	emp.Email = "bbbb@gmail.com"
	if err := EditEmployerProfile(emp, db); err != nil {
		t.Error(err)
	}
	emp2, err := GetEmployer("ashkan", db)
	if err != nil {
		t.Error(err)
	}
	if emp2.Password != "fjfjfj" || emp2.Email != "bbbb@gmail.com" {
		t.Errorf("%s %v", "Fail : ", emp2)
	}
}
