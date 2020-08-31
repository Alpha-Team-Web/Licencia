package database

import (
	"back-src/model/existence"
)

func (db Database) DoesFreelancerExistWithUsername(username string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = db.db.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (db *Database) DoesFreelancerExistWithEmail(email string) bool {
	resultSet := &[]existence.Freelancer{}
	_ = db.db.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (db *Database) InsertFreelancer(frl existence.Freelancer) error {
	_, err := db.db.Model(&frl).Insert()
	return err
}

func (db *Database) AddFreelancerSkills(username string, fieldId string, skills []string) error {
	var frl existence.Freelancer
	if err := db.db.Model(&frl). /*.Column("chosen_field_with_skills")*/ Where("username = ?", username).Select(); err != nil {
		return err
	}
	if frl.ChosenFieldWithSkills == nil {
		frl.ChosenFieldWithSkills = map[string][]string{}
	}
	frl.ChosenFieldWithSkills[fieldId] = skills
	_, err := db.db.Model(&frl).Column("chosen_field_with_skills").Where("username = ?", username).Update()
	return err
}

func (db *Database) GetFreelancerByUsername(username string) (existence.Freelancer, error) {
	var frl existence.Freelancer
	if err := db.db.Model(&frl).Where("username = ?", username).Select(); err != nil {
		return frl, err
	}
	return frl, nil
}

func (db *Database) GetFreelancerPasswordByUsername(username string) (string, error) {
	freelancer := existence.Freelancer{}
	err := db.db.Model(&freelancer).Where("username = ?", username).Column("password").Select()
	return freelancer.Password, err
}

func (db *Database) GetFreelancerUsernameByEmail(email string) (string, error) {
	freelancer := existence.Freelancer{}
	err := db.db.Model(&freelancer).Where("email = ?", email).Column("username").Select()
	return freelancer.Username, err
}

func (db *Database) GetFreelancer(username string) (existence.Freelancer, error) {
	frl := new(existence.Freelancer)
	err := db.db.Model(frl).Where("username = ?", username).Select()
	return *frl, err
}
