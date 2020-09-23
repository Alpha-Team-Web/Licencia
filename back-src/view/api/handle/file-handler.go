package handle

import (
	"back-src/controller/control/files"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const ProfileImageUploaderForName = "profileImage"
const ProjectFileFormName = "attachment"
const MaxProfileImageSizeInBytes = (6 * 1000000) //6 MegaBytes

func (handler *Handler) UploadProfileImage(ctx *gin.Context, profileType string) notifications.Notification {
	if profileType == existence.ProjectProfile {
		//TODO
		return notifications.Notification{}
	} else {
		if file, header, err := ctx.Request.FormFile(ProfileImageUploaderForName); err == nil {
			if err := files.UploadUserImage(getUsernameByContextToken(ctx), profileType, file, header, Db); err != nil {
				return notifications.GetInternalServerErrorNotif(ctx, nil)
			} else {
				return notifications.GetSuccessfulNotif(ctx, nil)
			}
		} else {
			return notifications.GetInternalServerErrorNotif(ctx, nil)
		}
	}
}

func (handler *Handler) DeleteProfileImage(context *gin.Context, profileType string) notifications.Notification {
	if profileType == existence.ProjectProfile {
		//TODO
		return notifications.Notification{}
	} else {
		if err := files.DeleteUserImage(getUsernameByContextToken(context), profileType, Db); err != nil {
			return notifications.GetDatabaseErrorNotif(context, nil)
		}
		return notifications.GetSuccessfulNotif(context, nil)
	}
}

func (handler *Handler) DownloadProfileImage(ctx *gin.Context, profileType string) notifications.Notification {
	if profileType == existence.ProjectProfile {
		//TODO
		return notifications.Notification{}
	} else {
		if file, err := files.DownloadUserImage(getUsernameByContextToken(ctx), profileType, Db); err != nil {
			return notifications.GetDatabaseErrorNotif(ctx, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, file)
		}
	}
}

func (handler *Handler) DownloadProjectFile(ctx *gin.Context) notifications.Notification {
	fileStruct := struct {
		Id string `json:"id"`
	}{}
	if err := ctx.ShouldBindJSON(&fileStruct); err == nil {
		if file, err := files.DownloadProjectFile(fileStruct.Id, SqlDb); err != nil {
			return notifications.GetDatabaseErrorNotif(ctx, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, file)
		}
	} else {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
}

func (handler *Handler) UploadProjectFile(ctx *gin.Context) notifications.Notification {
	form := data.AttachFileForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	header := form.Attachment
	attachment := existence.ProjectAttachment{}
	attachment.Size = header.Size
	attachment.Name = header.Filename
	if file, err := header.Open(); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		data, _ := ioutil.ReadAll(file)
		attachment.Data = data
	}
	attachment.ProjectId = form.ProjectId
	if err := files.AttachFileToProject(getUsernameByContextToken(ctx), attachment, SqlDb); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) UpdateProjectFile(ctx *gin.Context) notifications.Notification {
	form := data.UpdateFileForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	header := form.Attachment
	attachment := existence.ProjectAttachment{}
	attachment.Size = header.Size
	attachment.Name = header.Filename
	if file, err := header.Open(); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		data, _ := ioutil.ReadAll(file)
		attachment.Data = data
	}
	attachment.FileId = form.FileId
	if err := files.UpdateFileInProject(getUsernameByContextToken(ctx), attachment, SqlDb); err != nil {
		return notifications.Notification{
			Context:    ctx,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) RemoveProjectFile(ctx *gin.Context) notifications.Notification {
	fileStruct := struct {
		Id string `json:"id"`
	}{}
	if err := ctx.ShouldBindJSON(&fileStruct); err == nil {
		if err := files.DetachFileFromProject(getUsernameByContextToken(ctx), fileStruct.Id, SqlDb); err != nil {
			return notifications.GetDatabaseErrorNotif(ctx, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	} else {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
}
