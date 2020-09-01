package users

import (
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs"
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
			project.ProjectStatus = existence.Open

			//add new skills to all skills
			for field, skills := range project.FieldsWithSkills {
				oldSkills, err := DB.FieldTable.GetFieldSkills(field)
				//skips if field not found
				if err != nil {
					continue
				}
				for _, skill := range skills {
					if !libs.ContainsString(oldSkills, skill) {
						DB.FieldTable.AddSkillToField(field, skill)
					}
				}
			}

			project.Id = username + "-project-" + strconv.Itoa(len(emp.ProjectIds))
			DB.ProjectTable.AddProject(project)
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

func EditEmployerProject(token string, project existence.Project, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if _, err := DB.EmployerTable.GetEmployer(username); err == nil {
			if realProject, err := DB.ProjectTable.GetProject(project.Id); err == nil {
				if realProject.EmployerUsername == username {
					if realProject.ProjectStatus == existence.Open {
						return DB.ProjectTable.EditProject(realProject.Id, project)
					} else {
						return errors.New("project not open")
					}
				} else {
					return errors.New("project access denied")
				}
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

func AssignProjectToFreelancer(token string, freelancer string, projectId string, DB *database.Database) error {
	if _, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if requests, err := DB.ProjectTable.GetProjectRequests(projectId); err == nil {
			for s := range requests {
				DB.FreelancerTable.DeleteFreelancerRequestedProject(s, projectId)
			}
			if err := DB.FreelancerTable.AddFreelancerProjectId(freelancer, projectId); err != nil {
				return err
			}
			if err := DB.ProjectTable.SetProjectStatus(projectId, existence.OnGoing); err != nil {
				return err
			}
			if err := DB.ProjectTable.AddFreelancerToProject(freelancer, projectId); err != nil {
				return err
			}
			if err := DB.ProjectTable.DeleteProjectDescriptions(projectId); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
