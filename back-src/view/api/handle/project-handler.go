package handle

import (
	"back-src/controller/control/projects/filters"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) FilterFreelancer(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		filterReq := data.Filter{}
		if err := ctx.ShouldBindJSON(&filterReq); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if projects, err := filters.Filter(filterReq, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, projects)
		}
	}
}

func (handler *Handler) FilterEmployer(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		filterReq := data.Filter{}
		if err := ctx.ShouldBindJSON(&filterReq); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if projects, err := filters.Filter(filterReq, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, projects)
		}
	}
}
