package users

import (
	"back-src/controller/control/media"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"errors"
)

func ChooseFreelancerSkills(username string, fieldId string, skills []string, db *database.Database) error {
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
}

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

func GetFreelancer(token string, db *database.Database) (existence.Freelancer, error) {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if frl, err := db.FreelancerTable.GetFreelancer(username); err != nil {
			return existence.Freelancer{}, err
		} else {
			frl.Password = "N/A"
			return frl, nil
		}
	} else {
		return existence.Freelancer{}, err
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
					return errors.New("project status not suitable")
				}
			} else {
				return err
			}
		} else {
			return errors.New("invalid project id")
		}
	} else {
		return err
	}
}

func checkAbilityToRequestNewProject(username string, db *database.Database) error {
	e := errors.New("cant request more")
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
