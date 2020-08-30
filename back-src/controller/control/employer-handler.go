package control

import (
	"back-src/controller/control/utils/users"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (controller *Control) EditEmployerProfile(ctx *gin.Context) error {
	emp := existence.Employer{}
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		return err
	}
	return users.EditEmployerProfile(emp, DB)
}

func (controller *Control) GetEmployerProfile(ctx *gin.Context) (existence.Employer, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return existence.Employer{}, err
	}
	return users.GetEmployer(user.username, DB)
}

func (controller *Control) GetEmployerProjects(ctx *gin.Context) ([]existence.Project, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return users.GetEmployerProjects(user.username, DB)
}
