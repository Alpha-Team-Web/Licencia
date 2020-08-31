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
	port    string
	server  *gin.Engine
	handler *handle.Handler
}

func NewRouter(port string) Listener {

	var listener Listener = &router{port, gin.Default(), handle.NewControl()}
	return listener
}

func (router *router) Listen() error {
	router.server.POST("/register", func(context *gin.Context) {
		view.RespondRegister(context, router.handler.Register(context))
	})

	router.server.POST("/login", func(context *gin.Context) {
		token, err := router.handler.Login(context)
		if err == nil {
			router.handler.AddNewClock(token)
		}
		view.RespondLogin(context, token, err)
	})

	router.server.POST("/employer/edit-profile", func(context *gin.Context) {
		users.RespondEmployerEditProfile(context, router.handler.EditEmployerProfile(context))
	})

	router.server.GET("/employer/get-profile", func(context *gin.Context) {
		emp, token, err := router.handler.GetEmployerProfile(context)
		users.RespondEmployerGetProfile(context, token, emp, err)
	})

	router.server.GET("/freelancer/get-profile", func(context *gin.Context) {
		emp, token, err := router.handler.GetFreelancerProfile(context)
		users.RespondFreelancerGetProfile(context, token, emp, err)
	})

	router.server.Run(":" + router.port)
	return nil
}
