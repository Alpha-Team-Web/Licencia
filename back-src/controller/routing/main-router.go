package routing

import (
	"back-src/controller/handle"
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
	router.addIOEndpoints()
	router.addEmployerEndpoints()
	router.addFreelancerEndpoints()
	router.addProjectEndpoints()

	router.server.Run(":" + router.port)
	return nil
}
