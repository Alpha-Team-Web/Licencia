package users

import (
	"back-src/controller/utils/data"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
	"strconv"
)

func EditEmployerProfile(token string, emp existence.Employer, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return DB.EmployerTable.UpdateEmployerProfile(username, emp)
	} else {
		return err
	}
}

func EditEmployerPassword(token string, emp data.ChangePassRequest, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return DB.EmployerTable.UpdateEmployerPassword(username, emp.OldPass, emp.NewPass)
	} else {
		return err
	}
}

func GetEmployer(token string, DB *database.Database) (existence.Employer, error) {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if emp, err := DB.EmployerTable.GetEmployer(username); err != nil {
			return existence.Employer{}, err
		} else {
			emp.Password = "N/A"
			return emp, nil
		}
	} else {
		return existence.Employer{}, err
	}
}

func GetEmployerProjects(username string, DB *database.Database) ([]existence.Project, error) {
	if !DB.EmployerTable.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return DB.EmployerTable.GetEmployerProjects(username)
}

func AddProjectToEmployer(token string, project existence.Project, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if emp, err := DB.EmployerTable.GetEmployer(username); err == nil {
			project.EmployerUsername = username
			project.Id = username + "-project-" + strconv.Itoa(len(emp.ProjectIds))
			DB.EmployerTable.AddProject(project)
			emp.ProjectIds = append(emp.ProjectIds, project.Id)
			if err := DB.EmployerTable.UpdateEmployerProjects(username, emp); err == nil {
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
