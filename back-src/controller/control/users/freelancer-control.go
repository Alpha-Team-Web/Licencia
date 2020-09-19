package users

import (
	licencia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/media"
	"back-src/controller/control/projects/fields"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
	"back-src/view/data"
)

func EditFreelancerProfile(username string, frl existence.Freelancer, db *sql.Database) error {
	if err := db.FreelancerTable.UpdateFreelancerProfile(username, frl); err == nil {
		media.AddUpdateProfileEvent(username, true, db)
		return nil
	} else {
		return err
	}
}

func EditFreelancerPassword(username string, frl data.ChangePassRequest, db *sql.Database) error {
	freelancer, _ := db.FreelancerTable.GetFreelancer(username)
	if frl.OldPass != freelancer.Password {
		return licencia_errors.NewLicenciaError("password mismatch")
	}
	return db.FreelancerTable.UpdateFreelancerPassword(username, frl.OldPass, frl.NewPass)
}

func EditFreelancerLinks(username string, frl existence.Freelancer, db *sql.Database) error {
	if err := db.FreelancerTable.UpdateFreelancerLinks(username, frl); err == nil {
		media.AddUpdateProfileEvent(username, true, db)
		return nil
	} else {
		return err
	}
}

func GetFreelancer(username string, db *sql.Database) (existence.Freelancer, existence.File, error) {
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
}

func FreelancerRequestsForProject(username string, request data.FreelancerRequestForProject, db *sql.Database) error {
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
}

func checkProjectStatus(projectId, status string, db *sql.Database) error {
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

func checkAbilityToRequestNewProject(username string, db *sql.Database) error {
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

func AddSkillToFreelancer(username string, skillName string, db *sql.Database) error {
	if _, ok := fields.Engine.SkillWithField[skillName]; !ok {
		return licencia_errors.NewLicenciaError("no skill with such name exists.")
	} else {
		if skills, err := db.FreelancerTable.GetFreelancerSkills(username); err != nil {
			return err
		} else {
			if !libs.ContainsString(skills, skillName) {
				return db.FreelancerTable.AddFreelancerSkill(username, skillName)
			}
		}
	}
	return nil
}

func RemoveSkillFromFreelancer(username string, skillName string, db *sql.Database) error {
	if _, ok := fields.Engine.SkillWithField[skillName]; !ok {
		return licencia_errors.NewLicenciaError("no skill with such name exists.")
	} else {
		return db.FreelancerTable.RemoveFreelancerSkill(username, skillName)
	}
}

func GetFreelancerSkills(username string, db *sql.Database) ([]string, error) {
	skills, err := db.FreelancerTable.GetFreelancerSkills(username)
	return skills, err
}
