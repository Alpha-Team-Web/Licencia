package handle

import (
	"back-src/controller/control/media"
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func (handler *Handler) Register(ctx *gin.Context) notifications.Notification {

	switch accountType := ctx.Query("account-type"); accountType {

	case existence.EmployerType:
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, NotAssignedToken, nil)
		}
		if err := users.RegisterEmployer(employer, DB); err != nil {
			return notifications.GetDatabaseErrorNotif(ctx, NotAssignedToken, nil)
		}

	case existence.FreelancerType:
		freelancer := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&freelancer); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, NotAssignedToken, nil)
		}
		if err := users.RegisterFreelancer(freelancer, DB); err != nil {
			return notifications.GetDatabaseErrorNotif(ctx, NotAssignedToken, nil)
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
			return notifications.GetInternalServerErrorNotif(ctx, NotAssignedToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, token, nil)
		}
	default:
		return notifications.GetInvalidQueryErrorNotif(ctx, NotAssignedToken, nil)
	}
}

func CheckToken(token, userType string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		if libs.XNor(auth.IsFreelancer, userType == existence.FreelancerType) {
			return reInitToken(auth)
		} else {
			return "", errors.New("wrong user type token: " + token)
		}
	} else {
		return "", err
	}
}

func CheckTokenIgnoreType(token string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		return reInitToken(auth)
	} else {
		return "", err
	}
}

func formalCheckToken(token string) (existence.AuthToken, error) {
	if isThereAuth, err := DB.AuthTokenTable.IsThereAuthWithToken(token); err != nil {
		return existence.AuthToken{}, err
	} else if isThereAuth {
		if auth, err := DB.AuthTokenTable.GetAuthByToken(token); err != nil {
			return existence.AuthToken{}, err
		} else {
			return auth, err
		}
	} else {
		return existence.AuthToken{}, errors.New("not authorized token: " + token)
	}
}

func reInitToken(auth existence.AuthToken) (string, error) {
	currentTime := time.Now()
	if currentTime.Sub(auth.InitialTime) > AuthExpiryDur {
		if err := DB.AuthTokenTable.ChangeAuthUsage(auth.Token, false); err != nil {
			return "", err
		} else {
			newToken, err := users.MakeNewAuthToken(auth.Username, auth.IsFreelancer, DB)
			if err != nil {
				return "", err
			}
			if err := DB.AuthTokenTable.ChangeAuthUsage(newToken, true); err != nil {
				return "", err
			}
			return newToken, nil
		}
	} else {
		if err := DB.AuthTokenTable.ChangeAuthUsage(auth.Token, true); err != nil {
			return "", err
		} else {
			return auth.Token, nil
		}
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
