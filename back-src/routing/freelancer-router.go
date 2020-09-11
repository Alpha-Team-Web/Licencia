package routing

import (
	"back-src/model/existence"
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addFreelancerEndpoints() {
	router.addNewEndpointGroup("/freelancer", "freelancer", "").addCheckToken(existence.FreelancerType)
	router.addNewEndpointGroup("/profile", "freelancer-profile", "freelancer")
	router.addNewEndpointGroup("/projects", "freelancer-projects", "freelancer")

	router.addHandlerToPath("/get", "freelancer-profile", Get, func(context *gin.Context) {
		respond.Respond(router.handler.GetFreelancerProfile(context))
	})

	router.addHandlerToPath("/general", "freelancer-profile", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditFreelancerProfile(context))
	})

	router.addHandlerToPath("/password", "freelancer-profile", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditFreelancerPassword(context))
	})

	router.addHandlerToPath("/links", "freelancer-profile", Post, func(context *gin.Context) {
		respond.Respond(router.handler.EditFreelancerLinks(context))
	})

	router.addHandlerToPath("/review", "freelancer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.AddFreelancerReview(context))
	})

	router.addHandlerToPath("/request", "freelancer-projects", Post, func(context *gin.Context) {
		respond.Respond(router.handler.FreelancerRequestToProject(context))
	})

}
