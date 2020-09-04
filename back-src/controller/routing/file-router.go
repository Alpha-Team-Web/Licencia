package routing

import (
	"back-src/model/existence"
	"back-src/view/files"
	"github.com/gin-gonic/gin"
)

func (router *router) addFileEndpoints() {
	router.addNewEndpointGroup("/files", "files", "")
	router.addNewEndpointGroup("/profile-pic", "profile-pic", "files")
	router.addNewEndpointGroup("/freelancer", "freelancer-profile-pic", "profile-pic")
	router.addNewEndpointGroup("/employer", "employer-profile-pic", "profile-pic")

	router.addHandlerToPath("/upload", "freelancer-profile-pic", Post, func(context *gin.Context) {
		token, err := router.handler.UploadProfileImage(context, existence.FreelancerProfile)
		files.RespondUploadProfileImage(context, token, err)
	})

	router.addHandlerToPath("/upload", "employer-profile-pic", Post, func(context *gin.Context) {
		token, err := router.handler.UploadProfileImage(context, existence.EmployerProfile)
		files.RespondUploadProfileImage(context, token, err)
	})
}
