package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetFreelancerProfile(ctx *gin.Context) (existence.Freelancer, string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := CheckToken(token, existence.FreelancerType); err != nil {
		return existence.Freelancer{}, "", err
	} else {
		frl, err := users.GetFreelancer(newToken, DB)
		return frl, newToken, err
	}
}
