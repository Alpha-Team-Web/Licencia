package handle

import (
	licnecia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/media"
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Register(ctx *gin.Context) notifications.Notification {

	switch accountType := ctx.Query("account-type"); accountType {

	case existence.EmployerType:
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, NotAssignedToken, nil)
		}
		if err := users.RegisterEmployer(employer, DB); err != nil {
			if licnecia_errors.IsLicenciaError(err) {
				return notifications.GetExpectationFailedError(ctx, NotAssignedToken, licnecia_errors.GetErrorStrForRespond(err), nil)
			} else {
				return notifications.GetDatabaseErrorNotif(ctx, NotAssignedToken, nil)
			}
		}

	case existence.FreelancerType:
		freelancer := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&freelancer); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, NotAssignedToken, nil)
		}
		if err := users.RegisterFreelancer(freelancer, DB); err != nil {
			if licnecia_errors.IsLicenciaError(err) {
				return notifications.GetExpectationFailedError(ctx, NotAssignedToken, licnecia_errors.GetErrorStrForRespond(err), nil)
			} else {
				return notifications.GetDatabaseErrorNotif(ctx, NotAssignedToken, nil)
			}
		}

	default:
		return notifications.GetInvalidQueryErrorNotif(ctx, NotAssignedToken, nil)
	}
	return notifications.GetSuccessfulNotif(ctx, NotAssignedToken, nil)
}

func (handler *Handler) Login(ctx *gin.Context) notifications.Notification {
	loginReq := data.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, NotAssignedToken, nil)
	}
	switch accountType := ctx.Query("account-type"); accountType {
	case existence.EmployerType, existence.FreelancerType:
		loginReq.IsFreelancer = accountType == existence.FreelancerType
		if token, err := users.Login(loginReq, DB); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			AddNewClock(token)
			return notifications.GetSuccessfulNotif(ctx, token, nil)
		}
	default:
		return notifications.GetInvalidQueryErrorNotif(ctx, NotAssignedToken, nil)
	}
}

func (handler *Handler) ModifyFollow(ctx *gin.Context, isFollow bool) notifications.Notification {
	if newToken, err := CheckTokenIgnoreType(ctx.GetHeader("Token")); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		follow := existence.Follow{}
		if err := ctx.ShouldBindJSON(&follow); err == nil {
			job := libs.Ternary(isFollow, media.Follow, media.UnFollow).(func(string, existence.Follow, *database.Database) error)
			if err := job(newToken, follow, DB); err == nil {
				return notifications.GetSuccessfulNotif(ctx, newToken, nil)
			} else {
				return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
			}
		} else {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
	}
}
