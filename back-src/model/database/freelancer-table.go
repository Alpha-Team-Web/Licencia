package database

import "back-src/model/existence"

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
	var fieldsWithSkills = map[string][]string{}
	if err := db.db.Model().Table("freelancers").Where("username = ?", username).Select(&fieldsWithSkills); err != nil {
		return err
	}
	fieldsWithSkills[fieldId] = skills
	frl := existence.Freelancer{ChosenFieldWithSkills: fieldsWithSkills}
	_, err := db.db.Model(frl).Column("chosen_field_with_skills").Where("username = ?", username).Update()
	return err
}
