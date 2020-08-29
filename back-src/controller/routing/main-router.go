package routing

import (
	"back-src/controller/control"
	"back-src/view"

	"github.com/gin-gonic/gin"
)

type Listener interface {
	Listen() error
}

type router struct {
	port       string
	server     *gin.Engine
	controller *control.Control
}

func NewRouter(port string) Listener {

	var listener Listener = &router{port, gin.Default(), control.NewControl()}
	return listener
}

func (router *router) Listen() error {
	router.server.POST("/register", func(context *gin.Context) {
		view.RespondRegister(context, router.controller.Register(context))
	})

	router.server.POST("/employer/edit-profile", func(context *gin.Context) {
		view.RespondEmployerEditProfile(context, router.controller.EditEmployerProfile(context))
	})

	router.server.Run(":" + router.port)
	return nil
}
