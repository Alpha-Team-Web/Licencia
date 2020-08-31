package users

import (
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
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

func EditFreelancerProfile(frl existence.Freelancer, DB *database.Database) error {
	if !DB.DoesFreelancerExistWithUsername(frl.Username) {
		return errors.New("no user with such username :" + frl.Username)
	}

	if err := DB.UpdateFreelancerProfile(frl.Username, frl); err != nil {
		return err
	}

	return nil
}

func EditFreelancerPassword(frl data.ChangePassRequest, DB *database.Database) error {
	if !DB.DoesFreelancerExistWithUsername(frl.Username) {
		return errors.New("no user with such username:" + frl.Username)
	}

	if err := DB.UpdateFreelancerPassword(frl.Username, frl.OldPass, frl.NewPass); err != nil {
		return err
	}
	return nil
}

func EditFreelancerLinks(frl existence.Freelancer, DB *database.Database) error {
	if !DB.DoesFreelancerExistWithUsername(frl.Username) {
		return errors.New("no user with such username :" + frl.Username)
	}

	if err := DB.UpdateFreelancerLinks(frl.Username, frl); err != nil {
		return err
	}

	return nil
}

func GetFreelancer(token string, DB *database.Database) (existence.Freelancer, error) {
	if username, err := DB.GetUsernameByToken(token); err == nil {
		return DB.GetFreelancer(username)
	} else {
		return existence.Freelancer{}, err
	}
}
