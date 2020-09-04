package handle

import (
	"back-src/controller/control/users"
	"back-src/controller/utils/data"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditFreelancerProfile(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := checkToken(token, existence.FreelancerType); err != nil {
		return "", err
	} else {
		frl := existence.Freelancer{}
		frl.Username = "NNNNNN"
		frl.Password = "NNNNNN"
		frl.Email = "N@N.N"
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerProfile(token, frl, DB)
	}
}

func (handler *Handler) EditFreelancerPassword(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := checkToken(token, existence.FreelancerType); err != nil {
		return "", err
	} else {
		frl := data.ChangePassRequest{}
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerPassword(token, frl, DB)
	}
}

func (handler *Handler) EditFreelancerLinks(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := checkToken(token, existence.FreelancerType); err != nil {
		return "", err
	} else {
		frl := existence.Freelancer{}
		frl.Username = "NNNNNN"
		frl.Password = "NNNNNN"
		frl.Email = "N@N.N"
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return newToken, err
		}
		return newToken, users.EditFreelancerLinks(token, frl, DB)
	}
}

func (handler *Handler) GetFreelancerProfile(ctx *gin.Context) (existence.Freelancer, string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := checkToken(token, existence.FreelancerType); err != nil {
		return existence.Freelancer{}, "", err
	} else {
		frl, err := users.GetFreelancer(newToken, DB)
		return frl, newToken, err
	}
}

func (handler *Handler) FreelancerRequestToProject(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := checkToken(token, existence.FreelancerType); err != nil {
		return "", err
	} else {
		request := data.FreelancerRequestForProject{}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			return newToken, err
		} else {
			err := users.FreelancerRequestsForProject(newToken, request, DB)
			return newToken, err
		}
	}
}
