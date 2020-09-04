package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addProjectEndpoints() {
	router.addNewEndpointGroup("/projects", "projects", "")
	router.addNewEndpointGroup("/filters", "filters", "projects")

	router.addHandlerToPath("/freelancer", "filters", Get, func(context *gin.Context) {
		respond.Respond(router.handler.FilterFreelancer(context))
	})
	router.addHandlerToPath("/employer", "filters", Get, func(context *gin.Context) {
		respond.Respond(router.handler.FilterEmployer(context))
	})

}
