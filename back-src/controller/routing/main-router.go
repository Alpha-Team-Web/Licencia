package routing

import (
	"back-src/controller/handle"
	"back-src/view"
	"back-src/view/projects"
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

	router.server.POST("/employer/edit/profile", func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.server.POST("/employer/edit/password", func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.server.POST("/freelancer/edit/profile", func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.POST("/freelancer/edit/password", func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.POST("/freelancer/edit/links", func(context *gin.Context) {
		token, err := router.handler.EditFreelancerLinks(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.GET("/employer/get-profile", func(context *gin.Context) {
		emp, token, err := router.handler.GetEmployerProfile(context)
		users.RespondEmployerGetProfile(context, token, emp, err)
	})

	router.server.GET("/freelancer/get-profile", func(context *gin.Context) {
		emp, token, err := router.handler.GetFreelancerProfile(context)
		users.RespondFreelancerGetProfile(context, token, emp, err)
	})

	router.server.POST("/freelancer/project/review", func(context *gin.Context) {
		token, err := router.handler.AddFreelancerReview(context)
		projects.RespondFreelancerReview(context, token, err)
	})

	router.server.Run(":" + router.port)
	return nil
}
