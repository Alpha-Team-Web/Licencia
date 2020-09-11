package handle

import (
	"back-src/controller/control/files"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const ProfileImageUploaderForName = "profileImage"
const ProjectFileFormName = "attachment"
const MaxProfileImageSizeInBytes = (6 * 1000000) //6 MegaBytes

func (handler *Handler) UploadProfileImage(ctx *gin.Context, profileType string) notifications.Notification {
	if newToken, err := CheckTokenIgnoreType(ctx.GetHeader("Token")); err == nil {
		if profileType == existence.ProjectProfile {
			//TODO
			return notifications.Notification{}
		} else {
			if file, header, err := ctx.Request.FormFile(ProfileImageUploaderForName); err == nil {
				if err := files.UploadUserImage(newToken, profileType, file, header, DB); err != nil {
					return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
				} else {
					return notifications.GetSuccessfulNotif(ctx, newToken, nil)
				}
			} else {
				return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
			}
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}

func (handler *Handler) DownloadProfileImage(ctx *gin.Context, profileType string) notifications.Notification {

	if newToken, err := CheckTokenIgnoreType(ctx.GetHeader("Token")); err == nil {
		if profileType == existence.ProjectProfile {
			//TODO
			return notifications.Notification{}
		} else {
			if file, err := files.DownloadUserImage(newToken, profileType, DB); err != nil {
				return notifications.GetDatabaseErrorNotif(ctx, newToken, nil)
			} else {
				fmt.Println(file)
				return notifications.GetSuccessfulNotif(ctx, newToken, file)
			}
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}

func (handler *Handler) DownloadProjectFile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckTokenIgnoreType(ctx.GetHeader("Token")); err == nil {
		fileStruct := struct {
			Id string `json:"id"`
		}{}
		if err := ctx.ShouldBindJSON(&fileStruct); err == nil {
			if file, err := files.DownloadProjectFile(fileStruct.Id, DB); err != nil {
				return notifications.GetDatabaseErrorNotif(ctx, newToken, nil)
			} else {
				return notifications.GetSuccessfulNotif(ctx, newToken, file)
			}
		} else {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}

func (handler *Handler) UploadProjectFile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err == nil {
		form := data.AttachFileForm{}
		if err := ctx.ShouldBind(&form); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		header := form.Attachment
		attachment := existence.ProjectAttachment{}
		attachment.Size = header.Size
		attachment.Name = header.Filename
		if file, err := header.Open(); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		} else {
			data, _ := ioutil.ReadAll(file)
			attachment.Data = data
		}
		attachment.ProjectId = form.ProjectId
		if err := files.AttachFileToProject(newToken, attachment, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}

func (handler *Handler) UpdateProjectFile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err == nil {
		form := data.UpdateFileForm{}
		if err := ctx.ShouldBind(&form); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		header := form.Attachment
		attachment := existence.ProjectAttachment{}
		attachment.Size = header.Size
		attachment.Name = header.Filename
		if file, err := header.Open(); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		} else {
			data, _ := ioutil.ReadAll(file)
			attachment.Data = data
		}
		attachment.FileId = form.FileId
		if err := files.UpdateFileInProject(newToken, attachment, DB); err != nil {
			return notifications.Notification{
				Context:    ctx,
				Token:      newToken,
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}

func (handler *Handler) RemoveProjectFile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err == nil {
		fileStruct := struct {
			Id string `json:"id"`
		}{}
		if err := ctx.ShouldBindJSON(&fileStruct); err == nil {
			if err := files.DetachFileFromProject(newToken, fileStruct.Id, DB); err != nil {
				return notifications.GetDatabaseErrorNotif(ctx, newToken, nil)
			} else {
				return notifications.GetSuccessfulNotif(ctx, newToken, nil)
			}
		} else {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
	} else {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	}
}
