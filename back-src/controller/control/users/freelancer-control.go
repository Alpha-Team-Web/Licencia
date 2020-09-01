package users

import (
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func ChooseFreelancerSkills(username string, fieldId string, skills []string, DB *database.Database) error {
	if fieldSkills, err := DB.FieldTable.GetFieldSkills(fieldId); err == nil {
		if err := DB.FreelancerTable.AddFreelancerSkills(username, fieldId, skills); err != nil {
			return err
		}
		for _, skill := range skills {
			if !libs.ContainsString(fieldSkills, skill) {
				if err := DB.FieldTable.AddSkillToField(fieldId, skill); err != nil {
					return err
				}
			}
		}
		return nil
	} else {
		return err
	}
}

func EditFreelancerProfile(token string, frl existence.Freelancer, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return DB.FreelancerTable.UpdateFreelancerProfile(username, frl)
	} else {
		return err
	}
}

func EditFreelancerPassword(token string, frl data.ChangePassRequest, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return DB.FreelancerTable.UpdateFreelancerPassword(username, frl.OldPass, frl.NewPass)
	} else {
		return err
	}
}

func EditFreelancerLinks(token string, frl existence.Freelancer, DB *database.Database) error {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		return DB.FreelancerTable.UpdateFreelancerLinks(username, frl)
	} else {
		return err
	}
}

func GetFreelancer(token string, DB *database.Database) (existence.Freelancer, error) {
	if username, err := DB.AuthTokenTable.GetUsernameByToken(token); err == nil {
		if frl, err := DB.FreelancerTable.GetFreelancer(username); err != nil {
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
			len := len(requestedProjectIds)
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
	return nil
}
