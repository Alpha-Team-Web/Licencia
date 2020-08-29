package users

import (
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

func EditEmployerProfile(emp existence.Employer, DB *database.Database) error {
	if !DB.DoesEmployerExistWithUsername(emp.Username) {
		return errors.New("no user with such username :" + emp.Username)
	}

	if err := DB.UpdateEmployer(emp.Username, emp); err != nil {
		return err
	}

	return nil
}

func GetEmployerProfile(username string, DB *database.Database) (existence.Employer, error) {
	if !DB.DoesEmployerExistWithUsername(username) {
		return existence.Employer{}, errors.New("no user with such username :" + username)
	}
	return DB.GetEmployer(username)
}

func GetEmployerProjects(username string, DB *database.Database) ([]existence.Project, error) {
	if !DB.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return DB.GetEmployerProjects(username)
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
