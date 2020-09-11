package handle

import (
	"back-src/controller/control/projects"
	"back-src/model/existence"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) AddFreelancerReview(ctx *gin.Context) notifications.Notification {
	frlReview := existence.FreelancerEmployerReview{}
	if err := ctx.ShouldBindJSON(frlReview); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := projects.AddFreelancerReview(getTokenByContext(ctx), frlReview, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) AddEmployerReview(ctx *gin.Context) notifications.Notification {
	empReview := existence.EmployerFreelancerReview{}
	if err := ctx.ShouldBindJSON(empReview); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := projects.AddEmployerReview(getTokenByContext(ctx), empReview, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}
