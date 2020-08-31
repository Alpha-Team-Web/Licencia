package users

import (
	"back-src/controller/utils/data"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func EditEmployerProfile(token string, emp existence.Employer, DB *database.Database) error {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		return DB.UpdateEmployerProfile(username, emp)
	} else {
		return err
	}
}

func EditEmployerPassword(token string, emp data.ChangePassRequest, DB *database.Database) error {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		return DB.UpdateEmployerPassword(username, emp.OldPass, emp.NewPass)
	} else {
		return err
	}
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

func AddProjectToEmployer(token string, project existence.Project, DB *database.Database) error {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		project.EmployerUsername = username
		DB.AddProject(project)
		if emp, err := DB.GetEmployer(username); err == nil {
			emp.ProjectIds = append(emp.ProjectIds, project.Id)
			if err := DB.UpdateEmployerProjects(username, emp); err == nil {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
