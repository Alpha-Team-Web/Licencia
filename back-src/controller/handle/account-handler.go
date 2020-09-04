package handle

import (
	"back-src/controller/control/users"
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func (handler *Handler) Register(ctx *gin.Context) error {

	switch accountType := ctx.Query("account-type"); accountType {

	case existence.EmployerType:
		employer := existence.Employer{}
		if err := ctx.ShouldBindJSON(&employer); err != nil {
			return err
		}
		return users.RegisterEmployer(employer, DB)

	case existence.FreelancerType:
		freelancer := existence.Freelancer{}
		if err := ctx.ShouldBindJSON(&freelancer); err != nil {
			return err
		}
		return users.RegisterFreelancer(freelancer, DB)

	default:
		return errors.New("invalid query: " + accountType)
	}

}

func (handler *Handler) Login(ctx *gin.Context) (token string, error error) {
	loginReq := data.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		error = err
		return
	}
	switch accountType := ctx.Query("account-type"); accountType {
	case existence.EmployerType, existence.FreelancerType:
		loginReq.IsFreelancer = accountType == existence.FreelancerType
		token, error = users.Login(loginReq, DB)
	default:
		error = errors.New("invalid query: " + accountType)
	}
	return
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
