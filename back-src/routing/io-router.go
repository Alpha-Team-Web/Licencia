package routing

import (
	"back-src/view/api/respond"
	"github.com/gin-gonic/gin"
)

func (router *router) addIOEndpoints() {
	router.addNewEndpointGroup("/io", "io", "")

	router.addHandlerToPath("/register", "io", Post, func(context *gin.Context) {
		respond.Respond(router.handler.Register(context))
	})
	router.addHandlerToPath("/login", "io", Post, func(context *gin.Context) {
		notification := router.handler.Login(context)
		if notification.StatusCode == 200 {
			router.handler.AddNewClock(notification.Token)
		}
		respond.Respond(notification)
	})
}
