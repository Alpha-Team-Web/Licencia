package handle

import (
	"back-src/controller/control/projects/fields"
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
	if projects, err := filters.Filter(getAuthByContext(ctx), filterReq, Db); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, projects)
	}
}

func (handler *Handler) SearchSkill(context *gin.Context) notifications.Notification {
	return notifications.GetSuccessfulNotif(context, fields.SearchSkillStartsWith(context.Query("skill-name")))
}

func (handler *Handler) GetFieldSkills(ctx *gin.Context) notifications.Notification {
	return notifications.GetSuccessfulNotif(ctx, fields.GetSkillsByField(ctx.Query("field-id")))
}

func (handler *Handler) AddSkillToField(ctx *gin.Context) notifications.Notification {
	req := struct {
		FieldId   string `json:"field-id"`
		SkillName string `json:"skill-name"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := fields.AddSkillToField(req.FieldId, req.SkillName, SqlDb); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	}
	return notifications.GetSuccessfulNotif(ctx, nil)
}

func (handler *Handler) GetFields(ctx *gin.Context) notifications.Notification {
	return notifications.GetSuccessfulNotif(ctx, fields.GetFieldsWithoutSkills())
}
