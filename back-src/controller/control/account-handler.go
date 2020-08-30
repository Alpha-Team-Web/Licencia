package control

import (
	"back-src/controller/control/utils/data"
	"back-src/controller/control/utils/libs"
	"back-src/controller/control/utils/users"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
	"github.com/gin-gonic/gin"
)

type Control struct {
}

var DB *database.Database

func NewControl() *Control {
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Control{}
}

/*
json{
username
email
password
}
QUERIES : "account-type"
*/
func (controller *Control) Register(ctx *gin.Context) error {

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

func (controller *Control) Login(ctx *gin.Context) (token string, error error) {
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
