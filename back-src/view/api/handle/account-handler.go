package handle

import (
	licnecia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/media"
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
	"strings"
)

func (handler *Handler) Register(ctx *gin.Context) notifications.Notification {
	switch accountType := ctx.Query("account-type"); accountType {

	case existence.EmployerType:
		return handler.registerEmployer(ctx)

	case existence.FreelancerType:
		return handler.registerFreelancer(ctx)

	default:
		return notifications.GetInvalidQueryErrorNotif(ctx, nil)
	}
}

func (*Handler) registerEmployer(ctx *gin.Context) notifications.Notification {
	employer := existence.Employer{}
	if err := ctx.ShouldBindJSON(&employer); err != nil {
		if strings.Contains(err.Error(), "the 'email' tag") {
			return notifications.GetExpectationFailedError(ctx, "invalid email", nil)
		} else {
			return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
		}
	}
	if err := users.RegisterEmployer(employer, DB); err != nil {
		if licnecia_errors.IsLicenciaError(err) {
			return notifications.GetExpectationFailedError(ctx, licnecia_errors.GetErrorStrForRespond(err), nil)
		} else {
			return notifications.GetDatabaseErrorNotif(ctx, nil)
		}
	}
	return notifications.GetSuccessfulNotif(ctx, nil)
}

func (*Handler) registerFreelancer(ctx *gin.Context) notifications.Notification {
	freelancer := existence.Freelancer{}
	if err := ctx.ShouldBindJSON(&freelancer); err != nil {
		if strings.Contains(err.Error(), "the 'email' tag") {
			return notifications.GetExpectationFailedError(ctx, "invalid email", nil)
		} else {
			return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
		}
	}
	if err := users.RegisterFreelancer(freelancer, DB); err != nil {
		if licnecia_errors.IsLicenciaError(err) {
			return notifications.GetExpectationFailedError(ctx, licnecia_errors.GetErrorStrForRespond(err), nil)
		} else {
			return notifications.GetDatabaseErrorNotif(ctx, nil)
		}
	}
	return notifications.GetSuccessfulNotif(ctx, nil)
}

func (handler *Handler) Login(ctx *gin.Context) notifications.Notification {
	loginReq := data.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	switch accountType := ctx.Query("account-type"); accountType {
	case existence.EmployerType, existence.FreelancerType:
		loginReq.IsFreelancer = accountType == existence.FreelancerType
		if token, err := users.Login(loginReq, DB, RedisApi); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			AddNewClock(token)
			ctx.Header("Token", token)
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	default:
		return notifications.GetInvalidQueryErrorNotif(ctx, nil)
	}
}

func (handler *Handler) ModifyFollow(ctx *gin.Context, isFollow bool) notifications.Notification {
	follow := existence.Follow{}
	if err := ctx.ShouldBindJSON(&follow); err == nil {
		job := libs.Ternary(isFollow, media.Follow, media.UnFollow).(func(existence.AuthToken, existence.Follow, *sql.Database) error)
		if err := job(getAuthByContext(ctx), follow, DB); err == nil {
			return notifications.GetSuccessfulNotif(ctx, nil)
		} else {
			return notifications.GetInternalServerErrorNotif(ctx, nil)
		}
	} else {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
}
