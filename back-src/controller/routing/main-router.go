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
	port           string
	server         *gin.Engine
	handler        *handle.Handler
	endpointGroups []*endpointGroup
}

func NewRouter(port string) Listener {
	var listener Listener = &router{port, gin.Default(), handle.NewControl(), []*endpointGroup{}}
	return listener
}

func (router *router) Listen() error {

	//jesus
	router.addNewEndpointGroup("/io", "io", "")
	//jesus
	router.addNewEndpointGroup("/employer", "employer", "")
	router.addNewEndpointGroup("/profile", "employer-profile", "employer")
	router.addNewEndpointGroup("/projects", "employer-projects", "employer")
	//jesus
	router.addNewEndpointGroup("/freelancer", "freelancer", "")
	router.addNewEndpointGroup("/profile", "freelancer-profile", "freelancer")
	router.addNewEndpointGroup("/projects", "freelancer-projects", "freelancer")

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

	router.addHandlerToPath("/general", "employer-profile", Post, func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondEmployerEdit(context, token, err)
	})

	//io
	//employer
	router.server.POST("/employer/profile/general", func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.server.POST("/employer/profile/password", func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondEmployerEdit(context, token, err)
	})

	router.server.POST("/employer/projects/add", func(context *gin.Context) {
		token, err := router.handler.AddEmployerProject(context)
		users.RespondEmployerAddProject(context, token, err)
	})

	router.server.GET("/employer/profile/get", func(context *gin.Context) {
		emp, token, err := router.handler.GetEmployerProfile(context)
		users.RespondEmployerGetProfile(context, token, emp, err)
	})

	router.server.GET("/freelancer/profile/get", func(context *gin.Context) {
		emp, token, err := router.handler.GetFreelancerProfile(context)
		users.RespondFreelancerGetProfile(context, token, emp, err)
	})

	router.server.POST("/freelancer/profile/general", func(context *gin.Context) {
		token, err := router.handler.EditEmployerProfile(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.POST("/freelancer/profile/password", func(context *gin.Context) {
		token, err := router.handler.EditEmployerPassword(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.POST("/freelancer/profile/links", func(context *gin.Context) {
		token, err := router.handler.EditFreelancerLinks(context)
		users.RespondFreelancerEdit(context, token, err)
	})

	router.server.POST("/freelancer/projects/review", func(context *gin.Context) {
		token, err := router.handler.AddFreelancerReview(context)
		projects.RespondFreelancerReview(context, token, err)
	})

	router.server.Run(":" + router.port)
	return nil
}
