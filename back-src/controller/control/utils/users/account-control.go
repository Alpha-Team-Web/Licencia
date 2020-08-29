package users

import (
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func RegisterEmployer(emp existence.Employer, DB *database.Database) error {
	if !DB.DoesEmployerExistWithUsername(emp.Username) {
		if !DB.DoesEmployerExistWithEmail(emp.Email) {
			return DB.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Email)
	}
	return errors.New("duplicate username: " + emp.Username)
}

func EditEmployerProfile(emp existence.Employer, DB *database.Database) error {
	if !DB.DoesEmployerExistWithUsername(emp.Username) {
		return errors.New("no user with such username :" + emp.Username)
	}

	if err := DB.UpdateEmployer(emp.Username, emp); err != nil {
		return err
	}

	return nil
}
