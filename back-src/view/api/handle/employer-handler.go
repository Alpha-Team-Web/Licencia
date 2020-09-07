package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) EditEmployerProfile(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		emp := existence.Employer{}
		emp.Username = "NNNNNN"
		emp.Password = "NNNNNN"
		emp.Email = "N@N.N"
		if err := ctx.ShouldBindJSON(&emp); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		} else {
			if err := users.EditEmployerProfile(newToken, emp, DB); err != nil {
				return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
			} else {
				return notifications.GetSuccessfulNotif(ctx, newToken, nil)
			}
		}
	}
}

func (handler *Handler) EditEmployerPassword(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		emp := data.ChangePassRequest{}
		if err := ctx.ShouldBindJSON(&emp); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		} else {
			if err := users.EditEmployerPassword(newToken, emp, DB); err != nil {
				return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
			} else {
				return notifications.GetSuccessfulNotif(ctx, newToken, nil)
			}
		}
	}
}

func (handler *Handler) GetEmployerProfile(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		if emp, err := users.GetEmployer(newToken, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, emp)
		}
	}
}

//TODO : FIX THIS FUCKING FUNCTION
func (handler *Handler) GetEmployerProjects(ctx *gin.Context) ([]existence.Project, error) {
	user := struct {
		username string `json:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return users.GetEmployerProjects(user.username, DB)
}

func (handler *Handler) AddEmployerProject(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		project := existence.Project{}
		if err := ctx.ShouldBindJSON(&project); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.AddProjectToEmployer(newToken, project, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) EditEmployerProject(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		project := existence.Project{}
		if err := ctx.ShouldBindJSON(&project); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.EditEmployerProject(newToken, project, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}

func (handler *Handler) AssignProjectToFreelancer(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		assign := struct {
			id         string `json:"string"`
			freelancer string `json:"freelancer"`
		}{}
		if err := ctx.ShouldBindJSON(&assign); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.AssignProjectToFreelancer(newToken, assign.freelancer, assign.id, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}
