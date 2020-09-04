package routing

import (
	"back-src/model/existence"
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addFileEndpoints() {
	router.addNewEndpointGroup("/files", "files", "")
	router.addNewEndpointGroup("/profile-pic", "profile-pic", "files")
	router.addNewEndpointGroup("/freelancer", "freelancer-profile-pic", "profile-pic")
	router.addNewEndpointGroup("/employer", "employer-profile-pic", "profile-pic")

	router.addHandlerToPath("/upload", "freelancer-profile-pic", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UploadProfileImage(context, existence.FreelancerProfile))
	})

	router.addHandlerToPath("/upload", "employer-profile-pic", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UploadProfileImage(context, existence.EmployerProfile))
	})
}
