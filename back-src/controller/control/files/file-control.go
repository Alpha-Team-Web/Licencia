package files

import (
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
	"errors"
	"io/ioutil"
	"mime/multipart"
)

func UploadUserImage(username string, profileType string, file multipart.File, header *multipart.FileHeader, db *sql.Database) error {
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

func DeleteUserImage(username, profileType string, db *sql.Database) error {
	return db.ProfileTable.DeleteProfileImage(existence.Profile{Id: username, Type: profileType})
}

func DownloadUserImage(username string, profileType string, db *sql.Database) (existence.File, error) {
	if prof, err := db.ProfileTable.GetProfileImage(profileType, username); err != nil {
		return existence.File{}, err
	} else {
		return prof.File, nil
	}
}

func DownloadProjectFile(fileId string, db *sql.Database) (existence.File, error) {
	attachment, err := db.ProjectAttachmentTable.GetProjectAttachmentById(fileId)
	if err != nil {
		return existence.File{}, err
	}
	return attachment.File, nil
}

func AttachFileToProject(username string, attachment existence.ProjectAttachment, db *sql.Database) error {
	if project, err := db.ProjectTable.GetProject(attachment.ProjectId); err != nil {
		return err
	} else {
		if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
			return err
		} else {
			if !libs.ContainsString(emp.ProjectIds, project.Id) {
				return errors.New("access to project denied")
			}
		}
		if len(project.FileIds) == 3 {
			return errors.New("max number of files exceeded")
		}
	}
	if id, err := users.MakeNewFileId(db); err != nil {
		return err
	} else {
		attachment.FileId = id
		if err := db.ProjectAttachmentTable.AddProjectAttachment(attachment); err != nil {
			return err
		}
		if err := db.ProjectAttachmentTable.AddAttachmentIdToProject(attachment.FileId, attachment.ProjectId); err != nil {
			return err
		}
	}
	return nil
}

func DetachFileFromProject(username string, fileId string, db *sql.Database) error {
	var projectId string
	if attachment, err := db.ProjectAttachmentTable.GetProjectAttachmentById(fileId); err != nil {
		return err
	} else {
		if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
			return err
		} else {
			if !libs.ContainsString(emp.ProjectIds, attachment.ProjectId) {
				return errors.New("access to project denied")
			}
			projectId = attachment.ProjectId
		}
	}
	if err := db.ProjectAttachmentTable.RemoveProjectAttachment(fileId); err != nil {
		return err
	}
	if err := db.ProjectAttachmentTable.RemoveAttachmentIdFromProject(fileId, projectId); err != nil {
		return err
	}
	return nil
}

func UpdateFileInProject(username string, attachment existence.ProjectAttachment, db *sql.Database) error {
	if project, err := db.ProjectTable.GetProject(attachment.ProjectId); err != nil {
		return err
	} else {
		if emp, err := db.EmployerTable.GetEmployer(username); err != nil {
			return err
		} else {
			if !libs.ContainsString(emp.ProjectIds, project.Id) {
				return errors.New("access to project denied")
			}
		}
	}
	return db.ProjectAttachmentTable.UpdateProjectAttachment(attachment.FileId, attachment)
}
