package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditFreelancerProfile(ctx *gin.Context) notifications.Notification {
	frl := existence.Freelancer{}
	frl.Username = "NNNNNN"
	frl.Password = "NNNNNN"
	frl.Email = "N@N.N"
	if err := ctx.ShouldBindJSON(&frl); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.EditFreelancerProfile(getTokenByContext(ctx), frl, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) EditFreelancerPassword(ctx *gin.Context) notifications.Notification {
	frl := data.ChangePassRequest{}
	if err := ctx.ShouldBindJSON(&frl); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.EditFreelancerPassword(getTokenByContext(ctx), frl, DB); err != nil {
		return makeOperationErrorNotification(ctx, err)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) EditFreelancerLinks(ctx *gin.Context) notifications.Notification {
	frl := existence.Freelancer{}
	frl.Username = "NNNNNN"
	frl.Password = "NNNNNN"
	frl.Email = "N@N.N"
	if err := ctx.ShouldBindJSON(&frl); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.EditFreelancerLinks(getTokenByContext(ctx), frl, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) GetFreelancerProfile(ctx *gin.Context) notifications.Notification {
	if frl, file, err := users.GetFreelancer(getTokenByContext(ctx), DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, frl, file)
	}
}

func (handler *Handler) FreelancerRequestToProject(ctx *gin.Context) notifications.Notification {
	request := data.FreelancerRequestForProject{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		if err := users.FreelancerRequestsForProject(getTokenByContext(ctx), request, DB); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	}
}

func (handler *Handler) AddSkillToFreelancer(ctx *gin.Context) notifications.Notification {
	skill := struct {
		Name string `json:"name"`
	}{}
	if err := ctx.ShouldBindJSON(&skill); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		if err := users.AddSkillToFreelancer(getTokenByContext(ctx), skill.Name, DB); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	}
}

func (handler *Handler) RemoveSkillFromFreelancer(ctx *gin.Context) notifications.Notification {
	skill := struct {
		Name string `json:"name"`
	}{}
	if err := ctx.ShouldBindJSON(&skill); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		if err := users.RemoveSkillFromFreelancer(getTokenByContext(ctx), skill.Name, DB); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	}
}

func (handler *Handler) GetSkillsFromFreelancer(context *gin.Context) notifications.Notification {
	if skills, err := users.GetFreelancerSkills(getTokenByContext(context), DB); err != nil {
		return notifications.GetInternalServerErrorNotif(context, skills)
	} else {
		return notifications.GetSuccessfulNotif(context, skills)
	}
}
