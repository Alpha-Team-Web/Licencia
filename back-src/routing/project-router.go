package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addProjectEndpoints() {
	router.addNewEndpointGroup("/projects", "projects", "")
	router.addNewEndpointGroup("/filters", "filters", "projects").addCheckTokenIgnoreType()

	router.addHandlerToPath("/get", "filters", Get, func(ctx *gin.Context) {
		respond.Respond(router.handler.Filter(ctx))
	})

}
