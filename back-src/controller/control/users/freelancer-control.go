package users

import (
	licencia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/media"
	"back-src/controller/control/projects/fields"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
)

/*func ChooseFreelancerSkills(username string, fieldId string, skills []string, db *database.Database) error {
	if fieldSkills, err := db.FieldTable.GetFieldSkills(fieldId); err == nil {
		if err := db.FreelancerTable.AddFreelancerSkills(username, fieldId, skills); err != nil {
			return err
		}
		for _, skill := range skills {
			if !libs.ContainsString(fieldSkills, skill) {
				if err := db.FieldTable.AddSkillToField(fieldId, skill); err != nil {
					return err
				}
			}
		}
		return nil
	} else {
		return err
	}
}*/

func EditFreelancerProfile(token string, frl existence.Freelancer, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if err := db.FreelancerTable.UpdateFreelancerProfile(username, frl); err == nil {
			media.AddUpdateProfileEvent(username, true, db)
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func EditFreelancerPassword(token string, frl data.ChangePassRequest, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		freelancer, _ := db.FreelancerTable.GetFreelancer(username)
		if frl.OldPass != freelancer.Password {
			return licencia_errors.NewLicenciaError("password mismatch")
		}
		return db.FreelancerTable.UpdateFreelancerPassword(username, frl.OldPass, frl.NewPass)
	} else {
		return err
	}
}

func EditFreelancerLinks(token string, frl existence.Freelancer, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if err := db.FreelancerTable.UpdateFreelancerLinks(username, frl); err == nil {
			media.AddUpdateProfileEvent(username, true, db)
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func GetFreelancer(token string, db *database.Database) (existence.Freelancer, existence.File, error) {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if frl, err := db.FreelancerTable.GetFreelancer(username); err != nil {
			return existence.Freelancer{}, existence.File{}, err
		} else {
			frl.Password = "N/A"
			if profile, err := db.ProfileTable.GetProfileImage(existence.FreelancerType, username); err == nil {
				return frl, profile.File, nil
			} else {
				return existence.Freelancer{}, existence.File{}, err
			}
		}
	} else {
		return existence.Freelancer{}, existence.File{}, err
	}
}

func FreelancerRequestsForProject(token string, request data.FreelancerRequestForProject, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if err := checkAbilityToRequestNewProject(username, db); err == nil {
			if err := checkProjectStatus(request.Id, existence.Open, db); err == nil {
				if err := db.FreelancerTable.AddRequestedProjectToFreelancer(username, request.Id); err == nil {
					media.AddRequestEvent(username, request.Id, db)
					return db.ProjectTable.AddRequestToProject(request.Id, username, request.Description)
				} else {
					return err
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

func checkProjectStatus(projectId, status string, db *database.Database) error {
	if isThere, err := db.ProjectTable.IsThereProjectWithId(projectId); err == nil {
		if isThere {
			if projectStatus, err := db.ProjectTable.GetProjectStatus(projectId); err == nil {
				if projectStatus == status {
					return nil
				} else {
					return licencia_errors.NewLicenciaError("project status not suitable")
				}
			} else {
				return err
			}
		} else {
			return licencia_errors.NewLicenciaError("invalid project id")
		}
	} else {
		return err
	}
}

func checkAbilityToRequestNewProject(username string, db *database.Database) error {
	e := licencia_errors.NewLicenciaError("cant request more")
	if accountType, err := db.FreelancerTable.GetFreelancerTypeByUsername(username); err == nil {
		if requestedProjectIds, err := db.FreelancerTable.GetFreelancerRequestedProjectIds(username); err == nil {
			if projectIds, err := db.FreelancerTable.GetFreelancerProjectIds(username); err == nil {
				len := len(requestedProjectIds)
				for _, id := range projectIds {
					stat, _ := db.ProjectTable.GetProjectStatus(id)
					if stat == existence.OnGoing {
						len++
					}
				}
				switch accountType {
				case existence.FreelancerBronze:
					if len >= existence.BronzeRequestSize {
						return e
					}
				case existence.FreelancerSilver:
					if len >= existence.SilverRequestSize {
						return e
					}
				case existence.FreelancerGold:
					if len >= existence.GoldRequestSize {
						return e
					}
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
	return nil
}

func AddSkillToFreelancer(token string, skillName string, db *database.Database) error {
	if _, ok := fields.Engine.SkillWithField[skillName]; !ok {
		return licencia_errors.NewLicenciaError("no skill with such name exists.")
	} else {
		if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
			if skills, err := db.FreelancerTable.GetFreelancerSkills(username); err != nil {
				return err
			} else {
				if !libs.ContainsString(skills, skillName) {
					return db.FreelancerTable.AddFreelancerSkill(username, skillName)
				}
			}
		} else {
			return err
		}
	}
	return nil
}

func RemoveSkillFromFreelancer(token string, skillName string, db *database.Database) error {
	if _, ok := fields.Engine.SkillWithField[skillName]; !ok {
		return licencia_errors.NewLicenciaError("no skill with such name exists.")
	} else {
		if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
			return db.FreelancerTable.RemoveFreelancerSkill(username, skillName)
		} else {
			return err
		}
	}
	return nil
}

func GetFreelancerSkills(token string, db *database.Database) ([]string, error) {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		skills, err := db.FreelancerTable.GetFreelancerSkills(username)
		return skills, err
	} else {
		return []string{}, err
	}
}
