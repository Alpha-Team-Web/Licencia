package handle

import (
	"back-src/controller/control/projects"
	"back-src/model/existence"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) AddFreelancerReview(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		frlReview := existence.FreelancerEmployerReview{}
		if err := ctx.ShouldBindJSON(frlReview); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := projects.AddFreelancerReview(newToken, frlReview, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) AddEmployerReview(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		empReview := existence.EmployerFreelancerReview{}
		if err := ctx.ShouldBindJSON(empReview); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := projects.AddEmployerReview(newToken, empReview, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}
