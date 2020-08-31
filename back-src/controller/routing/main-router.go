package routing

import (
	"back-src/controller/handle"
	"back-src/view"
	"back-src/view/users"
	"github.com/gin-gonic/gin"
)

type Listener interface {
	Listen() error
}

type router struct {
	port       string
	server     *gin.Engine
	controller *handle.Handler
}

func NewRouter(port string) Listener {

	var listener Listener = &router{port, gin.Default(), handle.NewControl()}
	return listener
}

func (router *router) Listen() error {
	router.server.POST("/register", func(context *gin.Context) {
		view.RespondRegister(context, router.controller.Register(context))
	})

	router.server.POST("/login", func(context *gin.Context) {
		token, err := router.controller.Login(context)
		if err == nil {
			router.controller.AddNewClock(token)
		}
		view.RespondLogin(context, token, err)
	})

	router.server.POST("/employer/edit-profile", func(context *gin.Context) {
		users.RespondEmployerEditProfile(context, router.controller.EditEmployerProfile(context))
	})

	router.server.GET("/employer/get-profile", func(context *gin.Context) {
		emp, token, err := router.controller.GetEmployerProfile(context)
		users.RespondEmployerGetProfile(context, token, emp, err)
	})

	router.server.Run(":" + router.port)
	return nil
}
