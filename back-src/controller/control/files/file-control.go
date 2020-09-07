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
	profile.Size = header.Size
	if has, _ := db.ProfileTable.HasProfile(username, profileType); has {
		return db.ProfileTable.UpdateProfileImage(profile)
	} else {
		return db.ProfileTable.AddProfileImage(profile)
	}
}

func DownloadUserImage(token string, profileType string, db *database.Database) (existence.File, error) {
	username, err := db.AuthTokenTable.GetUsernameByToken(token)
	if err != nil {
		return existence.File{}, err
	}
	if prof, err := db.ProfileTable.GetProfileImage(profileType, username); err != nil {
		return existence.File{}, err
	} else {
		return prof.File, nil
	}
}

func DownloadProjectFile(fileId string, db *database.Database) (existence.File, error) {
	attachment, err := db.ProjectAttachmentTable.GetProjectAttachmentById(fileId)
	if err != nil {
		return existence.File{}, err
	}
	return attachment.File, nil
}
