package files

import (
	"back-src/model/database"
	"back-src/model/existence"
	"io/ioutil"
	"mime/multipart"
)

func UploadUserImage(token string, profileType string, file multipart.File, header *multipart.FileHeader, db *database.Database) error {
	username, err := db.AuthTokenTable.GetUsernameByToken(token)
	if err != nil {
		return err
	}
	name := header.Filename
	result, _ := ioutil.ReadAll(file)
	profile := existence.Profile{
		Id:   username,
		Type: profileType,
	}
	profile.Name = name
	profile.Data = result
	return db.ProfileTable.AddProfileImage(profile)
}
