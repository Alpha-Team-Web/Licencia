package routing

import (
	"back-src/view"
	"github.com/gin-gonic/gin"
)

func (router *router) addIOEndpoints() {
	router.addNewEndpointGroup("/io", "io", "")
	router.addNewEndpointGroup("/projects", "projects", "io")

	router.addHandlerToPath("/register", "io", Post, func(context *gin.Context) {
		view.RespondRegister(context, router.handler.Register(context))
	})
	router.addHandlerToPath("/login", "io", Post, func(context *gin.Context) {
		token, err := router.handler.Login(context)
		if err == nil {
			router.handler.AddNewClock(token)
		}
		view.RespondLogin(context, token, err)
	})
}
