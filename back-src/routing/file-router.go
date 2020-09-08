package routing

import (
	"back-src/model/existence"
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addFileEndpoints() {
	router.addNewEndpointGroup("/files", "files", "")
	router.addNewEndpointGroup("/profile-pic", "profile-pic", "files")
	router.addNewEndpointGroup("/project", "project-files", "files")
	router.addNewEndpointGroup("/freelancer", "freelancer-profile-pic", "profile-pic")
	router.addNewEndpointGroup("/employer", "employer-profile-pic", "profile-pic")

	router.addHandlerToPath("/upload", "freelancer-profile-pic", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UploadProfileImage(context, existence.FreelancerProfile))
	})

	router.addHandlerToPath("/upload", "employer-profile-pic", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UploadProfileImage(context, existence.EmployerProfile))
	})

	router.addHandlerToPath("/download", "freelancer-profile-pic", Get, func(context *gin.Context) {
		respond.Respond(router.handler.DownloadProfileImage(context, existence.FreelancerProfile))
	})

	router.addHandlerToPath("/download", "employer-profile-pic", Get, func(context *gin.Context) {
		respond.Respond(router.handler.DownloadProfileImage(context, existence.EmployerProfile))
	})

	router.addHandlerToPath("/download", "project-files", Get, func(context *gin.Context) {
		respond.Respond(router.handler.DownloadProjectFile(context))
	})

	router.addHandlerToPath("/upload", "project-files", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UploadProjectFile(context))
	})

	router.addHandlerToPath("/update", "project-files", Post, func(context *gin.Context) {
		respond.Respond(router.handler.UpdateProjectFile(context))
	})

	router.addHandlerToPath("/remove", "project-files", Post, func(context *gin.Context) {
		respond.Respond(router.handler.RemoveProjectFile(context))
	})
}
