package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	ProjectFilesFormName = "attachments"
	ProjectImageFormName = "profileImage"
	ProjectDataFormName  = "project"
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

/*func (handler *Handler) AddEmployerProject(ctx *gin.Context) notifications.Notification {
	//TODO : fix attachments
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		project := existence.Project{}
		if err := ctx.ShouldBindJSON(&project); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		if err := users.AddProjectToEmployer(newToken, project, []existence.ProjectAttachment{}, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, newToken, nil)
		}
	}
}*/
//TODO : handle projectImage, handle proper responses
func (handler *Handler) AddEmployerProject(ctx *gin.Context) notifications.Notification {
	if newToken, err := checkToken(ctx.GetHeader("Token"), existence.EmployerType); err != nil {
		return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
	} else {
		form := data.ProjectForm{}
		if err := ctx.ShouldBind(&form); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		mForm, err := ctx.MultipartForm()
		if err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		attachmentHeaders := mForm.File[ProjectFilesFormName]
		var project existence.Project
		if err := json.Unmarshal([]byte(form.Project), &project); err != nil {
			return notifications.GetShouldBindJsonErrorNotif(ctx, newToken, nil)
		}
		attachments := []existence.ProjectAttachment{}
		for _, header := range attachmentHeaders {
			if file, err := header.Open(); err == nil {
				data, _ := ioutil.ReadAll(file)
				attachment := existence.ProjectAttachment{}
				attachment.Data = data
				attachment.Name = header.Filename
				attachment.Size = header.Size
				attachments = append(attachments, attachment)
			}
		}
		if err := users.AddProjectToEmployer(newToken, project, attachments, DB); err != nil {
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
