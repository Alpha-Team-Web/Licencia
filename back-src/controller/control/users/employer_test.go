package users

import (
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
}*/

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
	if err := db.UpdateEmployerProfile(emp.Username, emp); err != nil {
		t.Error(err)
	}
	emp2, err := db.GetEmployer("ashkan")
	if err != nil {
		t.Error(err)
	}
	if emp2.Email != "bbbb@gmail.com" {
		t.Errorf("%s %v", "Fail : ", emp2)
	}
	if emp2.Password == "fjfjfj" {
		t.Errorf("%s %v", "Fail : ", emp2)
	}
	if err := db.UpdateEmployerPassword(emp.Username, "dasdsa", "sadasdas"); err == nil {
		t.Error("Old pass must be the same. Fail.")
	}
	if err := db.UpdateEmployerPassword(emp.Username, "a12345", "sadasdas"); err != nil {
		t.Error(err)
	}
	emp3, err := db.GetEmployer("ashkan")
	if emp3.Password != "sadasdas" {
		t.Errorf("%s %v", "Fail : ", emp3)
	}
}

/*{
	"name":"helloworld2"
	"desc":"This is my first project."
	"fr-to-emp-comment":"N/A"
	"emp-to-fr-comment":"N/A"
}*/
/*{
	"name":"helloworld3"
	"desc":"This is my first project."
	"fr-to-emp-comment":"N/A"
	"emp-to-fr-comment":"N/A"
}*/
func TestAddProjectToEmployer(t *testing.T) {
	db := database.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	//after sending request
	emp, _ := db.GetEmployer("ashkan")
	if len(emp.ProjectIds) != 2 {
		t.Errorf("%s %v", "Number of projects is not 2.", emp.ProjectIds)
	}
}
