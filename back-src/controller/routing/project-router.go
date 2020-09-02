package routing

import "github.com/gin-gonic/gin"

func (router *router) addProjectEndpoints() {
	router.addNewEndpointGroup("/projects", "projects", "")
	router.addNewEndpointGroup("/filter", "filter", "project")

	router.addHandlerToPath("/freelancer", "filter", Get, func(ctx *gin.Context) {
		//handle

		//respond

	})

}
