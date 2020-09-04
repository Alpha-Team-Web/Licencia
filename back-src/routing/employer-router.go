package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addEmployerEndpoints() {
	router.addNewEndpointGroup("/employer", "employer", "")
	router.addNewEndpointGroup("/profile", "employer-profile", "employer")
	router.addNewEndpointGroup("/projects", "employer-projects", "employer")

	router.addHandlerToPath("/general", "employer-profile", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditEmployerProfile(context))
	})

	router.addHandlerToPath("/password", "employer-profile", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditEmployerPassword(context))
	})

	router.addHandlerToPath("/add", "employer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.AddEmployerProject(context))
	})

	router.addHandlerToPath("/get", "employer-profile", Get, func(context *gin.Context) {
		respond.Respond(router.handler.GetEmployerProfile(context))
	})

	router.addHandlerToPath("/edit", "employer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditEmployerProject(context))
	})

	router.addHandlerToPath("/assign", "employer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.AssignProjectToFreelancer(context))
	})

	router.addHandlerToPath("/review", "employer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.AddEmployerReview(context))
	})

}
