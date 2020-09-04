package routing

import (
	"back-src/view/to-be-deleted"
	"github.com/gin-gonic/gin"
)

func (router *router) addIOEndpoints() {
	router.addNewEndpointGroup("/io", "io", "")

	router.addHandlerToPath("/register", "io", Post, func(context *gin.Context) {
		to_be_deleted.RespondIO(router.handler.Register(context))
	})
	router.addHandlerToPath("/login", "io", Post, func(context *gin.Context) {
		notification := router.handler.Login(context)
		if notification.StatusCode == 200 {
			router.handler.AddNewClock(notification.Token)
		}
		to_be_deleted.RespondIO(notification)
	})
}
