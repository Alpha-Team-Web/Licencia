package routing

import (
	projects2 "back-src/view/projects"
	"github.com/gin-gonic/gin"
)

func (router *router) addProjectEndpoints() {
	router.addNewEndpointGroup("/projects", "projects", "")
	router.addNewEndpointGroup("/filters", "filters", "projects")

	router.addHandlerToPath("/freelancer", "filters", Get, func(ctx *gin.Context) {
		projects, token, err := router.handler.FilterFreelancer(ctx)
		projects2.RespondFilterProjects(ctx, projects, token, err)
	})
	router.addHandlerToPath("/employer", "filters", Get, func(ctx *gin.Context) {
		projects, token, err := router.handler.FilterEmployer(ctx)
		projects2.RespondFilterProjects(ctx, projects, token, err)
	})

}
