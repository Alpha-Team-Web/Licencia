package handle

import (
	"back-src/controller/control/files"
	"back-src/model/existence"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

const ProfileImageUploaderForName = "profileImage"

func (handler *Handler) UploadProfileImage(ctx *gin.Context, profileType string) notifications.Notification {

	if newToken, err := checkTokenIgnoreType(ctx.GetHeader("Token")); err == nil {
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
