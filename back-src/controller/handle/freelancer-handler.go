package handle

import (
	"back-src/controller/control/users"
	"back-src/controller/utils/data"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditFreelancerProfile(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return "", err
	} else {
		frl := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerProfile(newToken, frl, DB)
	}
}

func (handler *Handler) EditFreelancerPassword(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return "", err
	} else {
		frl := data.ChangePassRequest{}
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerPassword(newToken, frl, DB)
	}
}

func (handler *Handler) EditFreelancerLinks(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return "", err
	} else {
		frl := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerLinks(newToken, frl, DB)
	}
}

func (handler *Handler) GetFreelancerProfile(ctx *gin.Context) (existence.Freelancer, string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return existence.Freelancer{}, "", err
	} else {
		frl, err := users.GetFreelancer(newToken, DB)
		return frl, newToken, err
	}
}
