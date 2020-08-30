package database

import (
	"back-src/model/existence"
	"time"
)

//Make New Auth
//Extend Auth
//Expire Auth

func (db *Database) MakeNewAuth(username, token string, isFreelancer bool) (string, error) {
	auth := existence.AuthToken{"", token, username, time.Now(), isFreelancer, false}
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

func (db *Database) IsAuthUsed(token string) (bool, error) {
	var auth = existence.AuthToken{}
	if err := db.db.Model(&auth).Where("token = ?", token).Column("is_used").Select(); err != nil {
		return false, err
	}
	return auth.IsUsed, nil
}

func (db *Database) ChangeAuthUsage(token string, isUsed bool) error {
	var auth = existence.AuthToken{Token: token, IsUsed: isUsed}
	if _, err := db.db.Model(&auth).Column("is_used").Where("token = ?", token).Update(); err != nil {
		return err
	}
	return nil
}

func (db *Database) ExpireAuth(token string) error {
	_, err := db.db.Model(&existence.AuthToken{}).Where("token = ?", token).Delete()
	return err
}
