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

func (db *Database) InsertFreelancer(emp existence.Employer) error {
	_, err := db.db.Model(&emp).Insert()
	return err
}
