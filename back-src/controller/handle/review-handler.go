package handle

import (
	"back-src/controller/control/projects"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) AddFreelancerReview(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := CheckToken(token, existence.FreelancerType); err != nil {
		return "", err
	} else {
		frlReview := existence.FreelancerEmployerReview{}
		if err := ctx.ShouldBindJSON(frlReview); err != nil {
			return newToken, err
		}
		return newToken, projects.AddFreelancerReview(newToken, frlReview, DB)
	}
}
