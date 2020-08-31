package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditEmployerProfile(ctx *gin.Context) error {
	emp := existence.Employer{}
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		return err
	}
	return users.EditEmployerProfile(emp, DB)
}

func (handler *Handler) GetEmployerProfile(ctx *gin.Context) (existence.Employer, string, error) {
	token := ctx.GetHeader("Token")
	if newToken, err := CheckToken(token, existence.EmployerType); err != nil {
		return existence.Employer{}, "", err
	} else {
		emp, err := users.GetEmployer(newToken, DB)
		return emp, newToken, err
	}
}

func (handler *Handler) GetEmployerProjects(ctx *gin.Context) ([]existence.Project, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return users.GetEmployerProjects(user.username, DB)
}
