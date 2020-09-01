package handle

import (
	"back-src/controller/control/users"
	"back-src/controller/utils/data"
	"back-src/model/existence"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditEmployerProfile(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return "", err
	} else {
		emp := existence.Employer{}
		if err := ctx.ShouldBindJSON(&emp); err != nil {
			return newToken, err
		}
		return newToken, users.EditEmployerProfile(newToken, emp, DB)
	}
}

func (handler *Handler) EditEmployerPassword(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return "", err
	} else {
		emp := data.ChangePassRequest{}
		if err := ctx.ShouldBindJSON(&emp); err != nil {
			return newToken, err
		}
		return newToken, users.EditEmployerPassword(newToken, emp, DB)
	}
}

func (handler *Handler) GetEmployerProfile(ctx *gin.Context) (existence.Employer, string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
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

func (handler *Handler) AddEmployerProject(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return newToken, err
	} else {
		project := existence.Project{}
		if err := ctx.ShouldBindJSON(&project); err != nil {
			return newToken, err
		}
		err := users.AddProjectToEmployer(newToken, project, DB)
		return newToken, err
	}
}

func (handler *Handler) EditEmployerProject(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return newToken, err
	} else {
		project := existence.Project{}
		if err := ctx.ShouldBindJSON(&project); err != nil {
			return newToken, err
		}
		err := users.EditEmployerProject(newToken, project, DB)
		return newToken, err
	}
}

func (handler *Handler) AssignProjectToFreelancer(ctx *gin.Context) (string, error) {
	if newToken, err := CheckToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return newToken, err
	} else {
		assign := struct {
			id         string `json:"string"`
			freelancer string `json:"freelancer"`
		}{}
		if err := ctx.ShouldBindJSON(&assign); err != nil {
			return newToken, err
		}
		err := users.AssignProjectToFreelancer(newToken, assign.freelancer, assign.id, DB)
		return newToken, err
	}
}
