package routing

import (
	"back-src/view/projects"
	"back-src/view/users"
	"github.com/gin-gonic/gin"
)

func (router *router) addFreelancerEndpoints() {
	router.addNewEndpointGroup("/freelancer", "freelancer", "")
	router.addNewEndpointGroup("/profile", "freelancer-profile", "freelancer")
	router.addNewEndpointGroup("/projects", "freelancer-projects", "freelancer")

	router.addHandlerToPath("/get", "freelancer-profile", Get, func(context *gin.Context) {
		emp, token, err := router.handler.GetFreelancerProfile(context)
		users.RespondFreelancerGetProfile(context, token, emp, err)
	})

	router.addHandlerToPath("/general", "freelancer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.addHandlerToPath("/password", "freelancer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.addHandlerToPath("/links", "freelancer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditFreelancerLinks(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.addHandlerToPath("/review", "freelancer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.AddFreelancerReview(context)
		projects.RespondFreelancerReview(context, token, err)
	})

}
