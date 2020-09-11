package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addMediaEndpoints() {
	router.addNewEndpointGroup("/media", "media", "")
	router.addNewEndpointGroup("/following", "media-following", "media").addCheckTokenIgnoreType()

	router.addHandlerToPath("/modify", "media-following", Post, func(ctx *gin.Context) {
		respond.Respond(router.handler.ModifyFollow(ctx, true))
	})
	router.addHandlerToPath("/modify", "media-following", Delete, func(ctx *gin.Context) {
		respond.Respond(router.handler.ModifyFollow(ctx, false))
	})

}
