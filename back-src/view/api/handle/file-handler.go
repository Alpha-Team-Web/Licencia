package handle

import (
	"back-src/controller/control/files"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

const ProfileImageUploaderForName = "profileImage"

func (handler *Handler) UploadProfileImage(ctx *gin.Context, profileType string) (string, error) {

	if newToken, err := checkTokenIgnoreType(ctx.GetHeader("Token")); err == nil {
		if profileType == existence.ProjectProfile {
			//TODO
			return newToken, nil
		} else {
			if file, header, err := ctx.Request.FormFile(ProfileImageUploaderForName); err == nil {
				return newToken, files.UploadUserImage(newToken, profileType, file, header, DB)
			} else {
				panic(err)
				return newToken, err
			}
		}
	} else {
		return "", err
	}
}
