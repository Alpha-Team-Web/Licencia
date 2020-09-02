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
	if isThereAuth, err := DB.AuthTokenTable.IsThereAuthWithToken(token); err != nil {
		return "", err
	} else if isThereAuth {
		if auth, err := DB.AuthTokenTable.GetAuthByToken(token); err != nil {
			return "", err
		} else {
			if libs.XNor(auth.IsFreelancer, userType == existence.FreelancerType) {
				if newToken, err := reInitToken(auth); err != nil {
					return "", err
				} else {
					return newToken, nil
				}
			} else {
				return "", errors.New("wrong user type token: " + token)
			}
		}
	} else {
		return "", errors.New("not authorized token: " + token)
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
