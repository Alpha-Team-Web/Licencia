package users

import (
	"back-src/controller/control/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func RegisterEmployer(emp existence.Employer, Db *database.Database) error {
	if !Db.DoesEmployerExistWithUsername(emp.Username) {
		if !Db.DoesEmployerExistWithEmail(emp.Email) {
			return Db.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Email)
	}
	return errors.New("duplicate username: " + emp.Username)
}

func RegisterFreelancer(frl existence.Freelancer, Db *database.Database) error {
	if !Db.DoesFreelancerExistWithUsername(frl.Username) {
		if !Db.DoesEmployerExistWithEmail(frl.Email) {
			return Db.InsertFreelancer(frl)
		}
		return errors.New("duplicate email: " + frl.Email)
	}
	return errors.New("duplicate username: " + frl.Username)
}

func Login(Id string, password string, isFreelancer bool, Db *database.Database) (string, error) {
	var passwordGetter func(string) (string, error)
	if libs.IsEmailValid(Id) {
		if isFreelancer {

		} else {

		}
	} else {
		if isFreelancer {

		} else {

		}
	}
}
