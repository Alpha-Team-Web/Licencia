package tables

import (
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"errors"
	"github.com/go-pg/pg"
)

type FreelancerTable struct {
	conn *pg.DB
}

func NewFreelancerTable(db *pg.DB) *FreelancerTable {
	return &FreelancerTable{db}
}

func (table *FreelancerTable) DoesFreelancerExistWithUsername(username string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.conn.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) DoesFreelancerExistWithEmail(email string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = table.conn.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (table *FreelancerTable) InsertFreelancer(frl existence.Freelancer) error {
	_, err := table.conn.Model(&frl).Insert()
	return err
}

func (table *FreelancerTable) AddFreelancerSkills(username string, fieldId string, skills []string) error {
	var frl existence.Freelancer
	if err := table.conn.Model(&frl). /*.Column("chosen_field_with_skills")*/ Where("username = ?", username).Select(); err != nil {
		return err
	}
	if frl.ChosenFieldWithSkills == nil {
		frl.ChosenFieldWithSkills = map[string][]string{}
	}
	frl.ChosenFieldWithSkills[fieldId] = skills
	_, err := table.conn.Model(&frl).Column("chosen_field_with_skills").Where("username = ?", username).Update()
	return err
}

func (table *FreelancerTable) UpdateFreelancerProfile(username string, frl existence.Freelancer) error {
	if _, err := table.conn.Model(&frl).Column("shown_name", "description", "first_name", "last_name", "phone_number", "address").Where("username = ?", username).Update(); err != nil {
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
	if _, err := table.conn.Model(&frl).Column("password").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) UpdateFreelancerLinks(username string, frl existence.Freelancer) error {
	if _, err := table.conn.Model(&frl).Column("website", "github_account", "github_repos").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) GetFreelancerByUsername(username string) (existence.Freelancer, error) {
	var frl existence.Freelancer
	if err := table.conn.Model(&frl).Where("username = ?", username).Select(); err != nil {
		return frl, err
	}
	return frl, nil
}

func (table *FreelancerTable) GetFreelancerPasswordByUsername(username string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.conn.Model(&freelancer).Where("username = ?", username).Column("password").Select()
	return freelancer.Password, err
}

func (table *FreelancerTable) GetFreelancerUsernameByEmail(email string) (string, error) {
	freelancer := existence.Freelancer{}
	err := table.conn.Model(&freelancer).Where("email = ?", email).Column("username").Select()
	return freelancer.Username, err
}

func (table *FreelancerTable) GetFreelancer(username string) (existence.Freelancer, error) {
	frl := new(existence.Freelancer)
	err := table.conn.Model(frl).Where("username = ?", username).Select()
	return *frl, err
}

func (table *FreelancerTable) GetFreelancerTypeByUsername(username string) (string, error) {
	frl := existence.Freelancer{}
	err := table.conn.Model(&frl).Column("account_type").Where("username = ?", username).Select()
	if err != nil {
		return "", err
	}
	return frl.AccountType, nil
}

func (table *FreelancerTable) GetFreelancerRequestedProjectIds(username string) ([]string, error) {
	frl := existence.Freelancer{}
	err := table.conn.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Select()
	if err != nil {
		return []string{}, err
	}
	return frl.RequestedProjectIds, nil
}

func (table *FreelancerTable) AddRequestedProjectToFreelancer(username, projectId string) error {
	frl := existence.Freelancer{}
	if projectIds, err := table.GetFreelancerRequestedProjectIds(username); err == nil {
		frl.ProjectIds = projectIds
		if _, err := table.conn.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Update(); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func (table *FreelancerTable) DeleteFreelancerRequestedProject(username string, projectId string) error {
	frl, err := table.GetFreelancer(username)
	if err != nil {
		return err
	}
	var index int
	for i, id := range frl.RequestedProjectIds {
		if id == projectId {
			index = i
			break
		}
	}
	frl.RequestedProjectIds = libs.RemoveStringElement(frl.RequestedProjectIds, index)
	if _, err = table.conn.Model(&frl).Column("requested_project_ids").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) AddFreelancerProjectId(username string, projectId string) error {
	frl, err := table.GetFreelancer(username)
	if err != nil {
		return err
	}
	found := false
	for _, id := range frl.ProjectIds {
		if id == projectId {
			found = true
		}
	}
	if found {
		return nil
	}
	frl.ProjectIds = append(frl.ProjectIds, projectId)
	if _, err = table.conn.Model(&frl).Column("project_ids").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *FreelancerTable) GetFreelancerProjectIds(username string) ([]string, error) {
	frl := existence.Freelancer{}
	err := table.conn.Model(&frl).Column("project_ids").Where("username = ?", username).Select()
	if err != nil {
		return []string{}, err
	}
	return frl.ProjectIds, nil
}
