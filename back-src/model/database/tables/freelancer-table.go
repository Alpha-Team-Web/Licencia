package tables

import (
	"back-src/model/existence"
	"errors"
	"github.com/go-pg/pg"
)

type FreelancerTable struct {
	*pg.DB
}

func (table *FreelancerTable) DoesFreelancerExistWithUsername(username string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) DoesFreelancerExistWithEmail(email string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) InsertFreelancer(frl existence.Freelancer) error {
	_, err := table.Model(&frl).Insert()
	return err
}

func (table *FreelancerTable) AddFreelancerSkills(username string, fieldId string, skills []string) error {
	var frl existence.Freelancer
	if err := table.Model(&frl). /*.Column("chosen_field_with_skills")*/ Where("username = ?", username).Select(); err != nil {
		return err
	}
	if frl.ChosenFieldWithSkills == nil {
		frl.ChosenFieldWithSkills = map[string][]string{}
	}
	frl.ChosenFieldWithSkills[fieldId] = skills
	_, err := table.Model(&frl).Column("chosen_field_with_skills").Where("username = ?", username).Update()
	return err
}

func (table *FreelancerTable) UpdateFreelancerProfile(username string, frl existence.Freelancer) error {
	if _, err := table.Model(&frl).Column("shown_name", "description", "first_name", "last_name", "phone_number", "address").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) UpdateFreelancerPassword(username string, oldPass string, newPass string) error {
	frl, _ := table.GetFreelancer(username)
	if frl.Password != oldPass {
		return errors.New("password mismatch")
	}

	frl.Password = newPass
	if _, err := table.Model(&frl).Column("password").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) UpdateFreelancerLinks(username string, frl existence.Freelancer) error {
	if _, err := table.Model(&frl).Column("website", "github", "github_repos").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) GetFreelancerByUsername(username string) (existence.Freelancer, error) {
	var frl existence.Freelancer
	if err := table.Model(&frl).Where("username = ?", username).Select(); err != nil {
		return frl, err
	}
	return frl, nil
}

func (table *FreelancerTable) GetFreelancerPasswordByUsername(username string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.Model(&freelancer).Where("username = ?", username).Column("password").Select()
	return freelancer.Password, err
}

func (table *FreelancerTable) GetFreelancerUsernameByEmail(email string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.Model(&freelancer).Where("email = ?", email).Column("username").Select()
	return freelancer.Username, err
}

func (table *FreelancerTable) GetFreelancer(username string) (existence.Freelancer, error) {
	frl := new(existence.Freelancer)
	err := table.Model(frl).Where("username = ?", username).Select()
	return *frl, err
}

func (table *FreelancerTable) GetFreelancerTypeByUsername(username string) (string, error) {
	frl := existence.Freelancer{}
	err := table.Model(&frl).Column("account_type").Where("username = ?", username).Select()
	if err != nil {
		return "", err
	}
	return frl.AccountType, nil
}

func (table *FreelancerTable) GetFreelancerRequestedProjectIds(username string) ([]string, error) {
	frl := existence.Freelancer{}
	err := table.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Select()
	if err != nil {
		return []string{}, err
	}
	return frl.RequestedProjectIds, nil
}

func (table *FreelancerTable) AddRequestedProjectToFreelancer(username, projectId string) error {
	frl := existence.Freelancer{}
	if projectIds, err := table.GetFreelancerRequestedProjectIds(username); err == nil {
		frl.ProjectIds = projectIds
		if _, err := table.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Update(); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}
