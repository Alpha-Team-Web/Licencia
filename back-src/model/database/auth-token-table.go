package database

import (
	"back-src/model/existence"
	"time"
)

//Make New Auth
//Extend Auth
//Expire Auth

func (db *Database) MakeNewAuth(username, token string, isFreelancer bool) (string, error) {
	auth := existence.AuthToken{"", token, username, time.Now(), isFreelancer}
	if _, err := db.db.Model(&auth).Insert(); err != nil {
		return "", err
	}
	return token, nil
}

func (db *Database) IsThereAuthWithToken(token string) (bool, error) {
	var resultSet []existence.AuthToken
	error := db.db.Model(&resultSet).Where("token = ?", token).Select()
	return len(resultSet) != 0, error
}
