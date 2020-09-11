package handle

import (
	"back-src/controller/control/projects/filters"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Filter(ctx *gin.Context) notifications.Notification {
	filterReq := data.Filter{}
	if err := ctx.ShouldBindJSON(&filterReq); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if projects, err := filters.Filter(filterReq, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, projects)
	}
}
