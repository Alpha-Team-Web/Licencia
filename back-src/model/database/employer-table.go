package database

import (
	"back-src/model/existence"
)

func (db Database) DoesEmployerExistWithUsername(username string) bool {
	resultSet := &[]existence.Employer{}
	_ = db.db.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (db *Database) DoesEmployerExistWithEmail(email string) bool {
	resultSet := &[]existence.Employer{}
	_ = db.db.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (db *Database) InsertEmployer(emp existence.Employer) error {
	_, err := db.db.Model(&emp).Insert()
	return err
}
