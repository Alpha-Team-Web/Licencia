package users

import (
	"back-src/controller/control/projects/filters"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"errors"
)

const (
	ProjectIdSize = 15
	FileIdSize    = 16
)

func EditEmployerProfile(token string, emp existence.Employer, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return db.EmployerTable.UpdateEmployerProfile(username, emp)
	} else {
		return err
	}
}

func EditEmployerPassword(token string, emp data.ChangePassRequest, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return db.EmployerTable.UpdateEmployerPassword(username, emp.OldPass, emp.NewPass)
	} else {
		return err
	}
}

func GetEmployer(token string, db *database.Database) (existence.Employer, error) {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
			return existence.Employer{}, err
		} else {
			emp.Password = "N/A"
			return emp, nil
		}
	} else {
		return existence.Employer{}, err
	}
}

func GetEmployerProjects(username string, db *database.Database) ([]existence.Project, error) {
	if !db.EmployerTable.DoesEmployerExistWithUsername(username) {
		return nil, errors.New("no user with such username :" + username)
	}
	return db.EmployerTable.GetEmployerProjects(username)
}

//func AddProjectWithFilesToEmployer(token string, project existence.Project, attachments []existence.ProjectAttachment, db *database.Database) error {
//	project.FileIds = []string{}
//	if err := AddProjectToEmployer(token, project, db); err != nil {
//		return err
//	}
//	if err := checkProjectFiles(project, attachments); err != nil {
//		return err
//	}
//	return nil
//}

func AddProjectToEmployer(token string, project existence.Project, attachments []existence.ProjectAttachment, db *database.Database) (e error) {
	e = nil
	if err := checkAddProjectFieldsValidity(project); err == nil {
		if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
			if emp, err := db.EmployerTable.GetEmployer(username); err == nil {
				project.EmployerUsername = username
				project.ProjectStatus = existence.Open
				project.FileIds = []string{}
				if project.Id, err = makeNewProjectId(db); err == nil {
					if err := checkProjectSkills(project.Id, project.FieldsWithSkills, db); err == nil {
						db.ProjectTable.AddProject(project)
						emp.ProjectIds = append(emp.ProjectIds, project.Id)
						if err := db.EmployerTable.UpdateEmployerProjects(username, emp); err == nil {
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
	} else {
		e = err
	}
	return
}

func checkAddProjectFieldsValidity(project existence.Project) error {
	error := errors.New("project fields not valid")
	if project.MinBudget < project.MaxBudget {
		return error
	}
	return nil
}

func checkProjectFiles(projectId string, attachments []existence.ProjectAttachment, db *database.Database) error {
	for i, attachment := range attachments {
		if id, err := makeNewFileId(db); err == nil {
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

func makeNewFileId(db *database.Database) (string, error) {
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

func makeNewProjectId(db *database.Database) (id string, e error) {
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

func checkProjectSkills(projectId string, fieldsWithSkills map[string][]string, db *database.Database) error {
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

func EditEmployerProject(token string, project existence.Project, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
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
	} else {
		return err
	}
}

func AssignProjectToFreelancer(token string, freelancer string, projectId string, db *database.Database) error {
	if _, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if requests, err := db.ProjectTable.GetProjectRequests(projectId); err == nil {
			for s := range requests {
				db.FreelancerTable.DeleteFreelancerRequestedProject(s, projectId)
			}
			if err := db.FreelancerTable.AddFreelancerProjectId(freelancer, projectId); err != nil {
				return err
			}
			if err := db.ProjectTable.SetProjectStatus(projectId, existence.OnGoing); err != nil {
				return err
			}
			if err := db.ProjectTable.AddFreelancerToProject(freelancer, projectId); err != nil {
				return err
			}
			if err := db.ProjectTable.DeleteProjectDescriptions(projectId); err != nil {
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
