package routing

import (
	"back-src/view/projects"
	"back-src/view/users"
	"github.com/gin-gonic/gin"
)

func (router *router) addEmployerEndpoints() {
	router.addNewEndpointGroup("/employer", "employer", "")
	router.addNewEndpointGroup("/profile", "employer-profile", "employer")
	router.addNewEndpointGroup("/projects", "employer-projects", "employer")

	router.addHandlerToPath("/general", "employer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.addHandlerToPath("/password", "employer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.addHandlerToPath("/add", "employer-projects", Post, func(context *gin.Context) {
		token, err := router.handler.AddEmployerProject(context)
		users.RespondEmployerAddProject(context, token, err)
	})

	router.addHandlerToPath("/get", "employer-profile", Get, func(context *gin.Context) {
		emp, token, err := router.handler.GetEmployerProfile(context)
		users.RespondEmployerGetProfile(context, token, emp, err)
	})

	router.addHandlerToPath("/edit", "employer-projects", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerProject(context)
		users.RespondEmployerEditProject(context, token, err)
	})

	router.addHandlerToPath("/review", "freelancer-projects", Post, func(context *gin.Context) {
		token, err := router.handler.AddEmployerReview(context)
		projects.RespondReview(context, token, err)
	})

}
