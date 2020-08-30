package control

import (
	"back-src/controller/control/utils/data"
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
		token, error = users.Login(loginReq.Id, loginReq.Password, accountType == existence.FreelancerType, DB)
	default:
		error = errors.New("invalid query: " + accountType)
	}
	return
}
