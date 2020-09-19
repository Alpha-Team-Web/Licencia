package users

import (
	licencia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/media"
	"back-src/controller/control/projects/filters"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
	"back-src/view/data"
	"errors"
	"time"
)

const (
	ProjectIdSize = 15
	FileIdSize    = 16
)

func EditEmployerProfile(username string, emp existence.Employer, db *sql.Database) error {
	if err := db.EmployerTable.UpdateEmployerProfile(username, emp); err == nil {
		media.AddUpdateProfileEvent(username, false, db)
		return nil
	} else {
		return err
	}
}

func EditEmployerPassword(username string, emp data.ChangePassRequest, db *sql.Database) error {
	employer, _ := db.EmployerTable.GetEmployer(username)
	if emp.OldPass != employer.Password {
		return licencia_errors.NewLicenciaError("password mismatch")
	}
	return db.EmployerTable.UpdateEmployerPassword(username, emp.OldPass, emp.NewPass)
}

func GetEmployer(username string, db *sql.Database) (existence.Employer, existence.File, error) {
	if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
		return existence.Employer{}, existence.File{}, err
	} else {
		emp.Password = "N/A"
		if profile, err := db.ProfileTable.GetProfileImage(existence.EmployerProfile, username); err == nil {
			return emp, profile.File, nil
		} else {
			return existence.Employer{}, existence.File{}, err
		}
	}
}

func GetEmployerProjects(username string, db *sql.Database) ([]existence.Project, error) {
	if !db.EmployerTable.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return db.EmployerTable.GetEmployerProjects(username)
}

//func AddProjectWithFilesToEmployer(token string, project existence.Project, attachments []existence.ProjectAttachment, db *sql.Database) error {
//	project.FileIds = []string{}
//	if licencia-errors := AddProjectToEmployer(token, project, db); licencia-errors != nil {
//		return licencia-errors
//	}
//	if licencia-errors := checkProjectFiles(project, attachments); licencia-errors != nil {
//		return licencia-errors
//	}
//	return nil
//}

func AddProjectToEmployer(username string, project existence.Project, attachments []existence.ProjectAttachment, db *sql.Database) (e error) {
	e = nil
	if err := checkAddProjectFieldsValidity(project); err == nil {
		if emp, err := db.EmployerTable.GetEmployer(username); err == nil {
			project.EmployerUsername = username
			project.ProjectStatus = existence.Open
			project.FileIds = []string{}
			if project.Id, err = makeNewProjectId(db); err == nil {
				if project.FieldsWithSkills == nil {
					project.FieldsWithSkills = map[string][]string{}
				}
				if err := checkProjectSkills(project.Id, project.FieldsWithSkills, db); err == nil {
					project.InitDate = time.Now()
					project.FreelancerRequestsWithDescription = map[string]string{}
					db.ProjectTable.AddProject(project)
					emp.ProjectIds = append(emp.ProjectIds, project.Id)
					if err := db.EmployerTable.UpdateEmployerProjects(username, emp); err == nil {
						media.AddAddProjectEvent(username, project.Id, db)
						if err := checkProjectFiles(project.Id, attachments, db); err == nil {
							if e == nil {
								e = nil
							}
						} else {
							e = err
						}
					} else {
						e = err
					}
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func checkAddProjectFieldsValidity(project existence.Project) error {
	error := errors.New("project fields not valid")
	if project.MinBudget > project.MaxBudget {
		return error
	}
	if project.FinishDate.Before(project.StartDate) || project.StartDate.Before(time.Now()) {
		return error
	}
	return nil
}

func checkProjectFiles(projectId string, attachments []existence.ProjectAttachment, db *sql.Database) error {
	for i, attachment := range attachments {
		if id, err := MakeNewFileId(db); err == nil {
			attachment.ProjectId = projectId
			attachment.FileId = id
			db.ProjectAttachmentTable.AddProjectAttachment(attachment)
			db.ProjectAttachmentTable.AddAttachmentIdToProject(id, projectId)
		}
		if i > 2 {
			break
		}
	}
	return nil
}

func MakeNewFileId(db *sql.Database) (string, error) {
	var e error
	id := "f" + libs.GetRandomNumberAsString(FileIdSize-1, func(str string) bool {
		if isThere, err := db.ProjectAttachmentTable.IsThereFileWithId("f" + str); err != nil {
			e = err
			return false
		} else {
			return isThere
		}
	})
	return id, e
}

func makeNewProjectId(db *sql.Database) (id string, e error) {
	id = "p" + libs.GetRandomNumberAsString(ProjectIdSize-1, func(str string) bool {
		if isThere, err := db.ProjectTable.IsThereProjectWithId("p" + str); err != nil {
			e = err
			return false
		} else {
			return isThere
		}
	})
	return id, e
}

func checkProjectSkills(projectId string, fieldsWithSkills map[string][]string, db *sql.Database) error {
	for field, skills := range fieldsWithSkills {
		oldSkills, err := db.FieldTable.GetFieldSkills(field)
		if err != nil {
			continue
		}
		for _, skill := range skills {
			if !libs.ContainsString(oldSkills, skill) {
				if err := db.FieldTable.AddSkillToField(field, skill); err != nil {
					return err
				}
			}
			filters.AddSkillToProject(skill, projectId)
		}

	}
	return nil
}

func EditEmployerProject(username string, project existence.Project, db *sql.Database) error {
	if _, err := db.EmployerTable.GetEmployer(username); err == nil {
		if realProject, err := db.ProjectTable.GetProject(project.Id); err == nil {
			if realProject.EmployerUsername == username {
				if realProject.ProjectStatus == existence.Open {
					return db.ProjectTable.EditProject(realProject.Id, project)
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
}

func AssignProjectToFreelancer(username string, freelancer string, projectId string, db *sql.Database) error {
	if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
		return err
	} else if username != realUsername {
		return errors.New("not valid token for this project")
	} else {
		if requests, err := db.ProjectTable.GetProjectRequests(projectId); err != nil {
			return err
		} else if libs.ContainsKey(requests, freelancer) {
			if err := db.FreelancerTable.AddFreelancerProjectId(freelancer, projectId); err != nil {
				return err
			}
			if err := db.ProjectTable.SetProjectStatus(projectId, existence.OnGoing); err != nil {
				return err
			}
			if err := db.ProjectTable.AddFreelancerToProject(freelancer, projectId); err != nil {
				return err
			}
			if err := removeProjectRequests(projectId, db); err != nil {
				return err
			}
			media.AddAssignProjectEvent(username, freelancer, projectId, db)
			return nil
		} else {
			return errors.New("not valid freelancer")
		}
	}
}

func removeProjectRequests(projectId string, db *sql.Database) error {
	if givenMap, err := db.ProjectTable.DeleteProjectDescriptions(projectId); err != nil {
		return err
	} else {
		usernames := libs.GetKeySet(givenMap)
		for _, username := range usernames {
			db.FreelancerTable.DeleteFreelancerRequestedProject(username, projectId)
		}
	}
	return nil
}

func ExtendProject(username string, projectId string, finishDate time.Time, db *sql.Database) error {
	if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
		return err
	} else if username != realUsername {
		return errors.New("not valid token for this project")
	} else {
		if firstFinishDate, err := db.ProjectTable.GetProjectFinishDate(projectId); err != nil {
			return err
		} else if finishDate.Before(firstFinishDate) {
			return errors.New("not valid time")
		} else {
			if status, err := db.ProjectTable.GetProjectStatus(projectId); err != nil {
				return err
			} else if status != existence.Open {
				return errors.New("not valid project status")
			} else {
				media.AddExtendProjectEvent(username, projectId, db)
				return db.EmployerTable.ExtendProject(projectId, finishDate)
			}
		}
	}
}

func CloseProject(username string, projectId string, db *sql.Database) error {
	if realUsername, err := db.ProjectTable.GetEmployerUsernameByProjectId(projectId); err != nil {
		return err
	} else if username != realUsername {
		return errors.New("not valid token for this project")
	} else {
		if finishDate, err := db.ProjectTable.GetProjectFinishDate(projectId); err != nil {
			return err
		} else if finishDate.Before(time.Now()) {
			return errors.New("not valid time")
		} else {
			if status, err := db.ProjectTable.GetProjectStatus(projectId); err != nil {
				return err
			} else if status != existence.OnGoing {
				return errors.New("not valid project status")
			} else {
				freelancer, _ := db.ProjectTable.GetFreelancerUsernameByProjectId(projectId)
				media.AddCloseProjectEvent(username, freelancer, projectId, db)
				panic(errors.New("not implemented function error"))
				//TODO(El Tipo, Close Project)
			}
		}
	}
}
