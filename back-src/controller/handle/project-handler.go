package handle

import (
	"back-src/controller/control/projects/filters"
	"back-src/controller/utils/data"
	"back-src/model/existence"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) FilterFreelancer(ctx *gin.Context) ([]responses.ListicProject, string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return nil, "", err
	} else {
		filterReq := data.Filter{}
		if err := ctx.ShouldBindJSON(&filterReq); err != nil {
			return nil, newToken, err
		}
		if projects, err := filters.Filter(filterReq, DB); err != nil {
			return nil, newToken, err
		} else {
			return projects, newToken, nil
		}
	}
}

func (handler *Handler) FilterEmployer(ctx *gin.Context) ([]responses.ListicProject, string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return nil, "", err
	} else {
		filterReq := data.Filter{}
		if err := ctx.ShouldBindJSON(&filterReq); err != nil {
			return nil, newToken, err
		}
		if projects, err := filters.Filter(filterReq, DB); err != nil {
			return nil, newToken, err
		} else {
			return projects, newToken, nil
		}
	}
}
