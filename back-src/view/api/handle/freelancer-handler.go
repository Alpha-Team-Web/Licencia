package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditFreelancerProfile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		frl := existence.Freelancer{}
		frl.Username = "NNNNNN"
		frl.Password = "NNNNNN"
		frl.Email = "N@N.N"
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.EditFreelancerProfile(newToken, frl, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) EditFreelancerPassword(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		frl := data.ChangePassRequest{}
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.EditFreelancerPassword(newToken, frl, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) EditFreelancerLinks(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		frl := existence.Freelancer{}
		frl.Username = "NNNNNN"
		frl.Password = "NNNNNN"
		frl.Email = "N@N.N"
		if err := ctx.ShouldBindJSON(&frl); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.EditFreelancerLinks(newToken, frl, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) GetFreelancerProfile(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		if frl, err := users.GetFreelancer(newToken, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, frl)
		}
	}
}

func (handler *Handler) FreelancerRequestToProject(ctx *gin.Context) notifications.Notification {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.FreelancerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		request := data.FreelancerRequestForProject{}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		} else {
			if err := users.FreelancerRequestsForProject(newToken, request, DB); err != nil {
				if err.Error() == "cant request more" {
					return notifications.GetExpectationFailedError(ctx, newToken, nil)
				} else {
					return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
				}
			} else {
				return notifications.GetSuccessfulNotif(ctx, newToken, nil)
			}
		}
	}
}
