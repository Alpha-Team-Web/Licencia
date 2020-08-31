package routing

import (
	"github.com/gin-gonic/gin"
)

type endpointGroup struct {
	name        string
	group       *gin.RouterGroup
	fatherGroup *endpointGroup
}

func (router *router) addNewEndpointGroup(endpointAddr string, name, fatherName string) *endpointGroup {
	fatherGroup := router.getEndpointGroupByName(fatherName)
	if fatherGroup == nil {
		endpointGroup := &endpointGroup{name: name, group: router.server.Group(endpointAddr), fatherGroup: nil}
		return endpointGroup
	} else {
		endpointGroup := &endpointGroup{name: name, group: fatherGroup.group.Group(endpointAddr), fatherGroup: fatherGroup}
		return endpointGroup
	}
}

func (router *router) getEndpointGroupByName(name string) *endpointGroup {
	for _, group := range router.endpointGroups {
		if group.name == name {
			return group
		}
	}
	return nil
}

const (
	Post = "POST"
	Get  = "GET"
)

func (router *router) addHandlerToPath(addr, endpointGroupName, method string, handle func(ctx *gin.Context)) {
	if addr == "" {
		//TODO
	} else {
		switch method {
		case "POST":
			router.getEndpointGroupByName(endpointGroupName).group.POST(addr, handle)
		case "GET":
			router.getEndpointGroupByName(endpointGroupName).group.GET(addr, handle)
		}
	}
}
