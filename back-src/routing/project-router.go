package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addProjectEndpoints() {
	router.addNewEndpointGroup("/projects", "projects", "")
	router.addNewEndpointGroup("/filters", "filters", "projects").addCheckTokenIgnoreType()
	router.addNewEndpointGroup("/fields", "fields", "").addCheckTokenIgnoreType()

	router.addHandlerToPath("/get", "filters", Get, func(ctx *gin.Context) {
		respond.Respond(router.handler.Filter(ctx))
	})

	router.addHandlerToPath("/search-skill", "fields", Get, func(context *gin.Context) {
		respond.Respond(router.handler.SearchSkill(context))
	})

	router.addHandlerToPath("/field-skills", "fields", Get, func(ctx *gin.Context) {
		respond.Respond(router.handler.GetFieldSkills(ctx))
	})

	router.addHandlerToPath("/field-skills", "fields", Post, func(ctx *gin.Context) {
		respond.Respond(router.handler.AddSkillToField(ctx))
	})

	router.addHandlerToPath("", "fields", Get, func(ctx *gin.Context) {
		respond.Respond(router.handler.GetFields(ctx))
	})

}
