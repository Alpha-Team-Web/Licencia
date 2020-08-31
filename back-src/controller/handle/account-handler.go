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
		token, error = users.Login(loginReq, getUsernameGetter(loginReq.Id, loginReq.IsFreelancer), getPasswordGetter(loginReq.IsFreelancer), DB)
	default:
		error = errors.New("invalid query: " + accountType)
	}
	return
}

func getUsernameGetter(Id string, isFreelancer bool) func() (username string, error error) {
	if isFreelancer {
		return getUsernameById(Id, DB.DoesFreelancerExistWithEmail, DB.DoesFreelancerExistWithUsername, DB.GetFreelancerUsernameByEmail)
	} else {
		return getUsernameById(Id, DB.DoesEmployerExistWithEmail, DB.DoesEmployerExistWithUsername, DB.GetEmployerUsernameByEmail)
	}
}

type doesExist func(string) bool
type getUsernameByEmail func(string) (string, error)

func getUsernameById(Id string, doesUserExistWithEmail doesExist, doesUserExistWithUsername doesExist, getUsername getUsernameByEmail) func() (string, error) {
	var username string
	var e error
	if libs.IsEmailValid(Id) {
		if doesUserExistWithEmail(Id) {
			if user, err := getUsername(Id); err == nil {
				username = user
			} else {
				e = err
			}
		} else {
			e = errors.New("not signed up email: " + Id)
		}
	} else {
		if doesUserExistWithUsername(Id) {
			username = Id
		} else {
			e = errors.New("not signed up username: " + Id)
		}
	}
	return func() (string, error) {
		return username, e
	}
}

func getPasswordGetter(isFreelancer bool) func(string) (string, error) {
	if isFreelancer {
		return DB.GetFreelancerPasswordByUsername
	} else {
		return DB.GetEmployerPasswordByUsername
	}
}

func CheckToken(token, userType string) (string, error) {
	if isThereAuth, err := DB.IsThereAuthWithToken(token); err != nil {
		return "", err
	} else if isThereAuth {
		if auth, err := DB.GetAuthByToken(token); err != nil {
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
		if err := DB.ChangeAuthUsage(auth.Token, false); err != nil {
			return "", err
		} else {
			newToken, err := users.MakeNewAuthToken(auth.Username, auth.IsFreelancer, DB)
			if err != nil {
				return "", err
			}
			if err := DB.ChangeAuthUsage(newToken, true); err != nil {
				return "", err
			}
			return newToken, nil
		}
	} else {
		if err := DB.ChangeAuthUsage(auth.Token, true); err != nil {
			return "", err
		} else {
			return auth.Token, nil
		}
	}
}
