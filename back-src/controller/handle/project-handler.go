package handle

import (
	"back-src/controller/utils/data"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) FilterFreelancer(ctx *gin.Context) ([]existence.ListicProject, string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return nil, "", err
	} else {
		filterReq := data.Filter{}
		if err := ctx.ShouldBindJSON(&filterReq); err != nil {
			return nil, newToken, err
		}
		return nil, newToken, err
	}
}
