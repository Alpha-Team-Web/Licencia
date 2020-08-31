package users

import (
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
)

func ChooseFreelancerSkills(username string, fieldId string, skills []string, database *database.Database) error {
	if fieldSkills, err := database.GetFieldSkills(fieldId); err == nil {
		if err := database.AddFreelancerSkills(username, fieldId, skills); err != nil {
			return err
		}
		for _, skill := range skills {
			if !libs.ContainsString(fieldSkills, skill) {
				if err := database.AddSkillToField(fieldId, skill); err != nil {
					return err
				}
			}
		}
		return nil
	} else {
		return err
	}
}

func GetFreelancer(token string, DB *database.Database) (existence.Freelancer, error) {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		if frl, err := DB.GetFreelancer(username); err != nil {
			return existence.Freelancer{}, err
		} else {
			frl.Password = "N/A"
			return frl, nil
		}
	} else {
		return existence.Freelancer{}, err
	}
}
