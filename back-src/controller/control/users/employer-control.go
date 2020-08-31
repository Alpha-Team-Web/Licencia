package users

import (
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func EditEmployerProfile(emp existence.Employer, DB *database.Database) error {
	if !DB.DoesEmployerExistWithUsername(emp.Username) {
		return errors.New("no user with such username :" + emp.Username)
	}

	if err := DB.UpdateEmployerProfile(emp.Username, emp); err != nil {
		return err
	}

	return nil
}

func GetEmployer(token string, DB *database.Database) (existence.Employer, error) {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		return DB.GetEmployer(username)
	} else {
		return existence.Employer{}, err
	}
}

func GetEmployerProjects(username string, DB *database.Database) ([]existence.Project, error) {
	if !DB.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return DB.GetEmployerProjects(username)
}
