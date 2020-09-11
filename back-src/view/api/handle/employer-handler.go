package handle

import (
	"back-src/controller/control/users"
	"back-src/model/existence"
	"back-src/view/data"
	"back-src/view/notifications"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

const (
	ProjectFilesFormName = "attachments"
	ProjectImageFormName = "profile"
	ProjectDataFormName  = "project"
)

func (handler *Handler) EditEmployerProfile(ctx *gin.Context) notifications.Notification {
	emp := existence.Employer{}
	emp.Username = "NNNNNN"
	emp.Password = "NNNNNN"
	emp.Email = "N@N.N"
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		if err := users.EditEmployerProfile(getTokenByContext(ctx), emp, DB); err != nil {
			return notifications.GetInternalServerErrorNotif(ctx, nil)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	}
}

func (handler *Handler) EditEmployerPassword(ctx *gin.Context) notifications.Notification {
	emp := data.ChangePassRequest{}
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	} else {
		if err := users.EditEmployerPassword(getTokenByContext(ctx), emp, DB); err != nil {
			return makeOperationErrorNotification(ctx, err)
		} else {
			return notifications.GetSuccessfulNotif(ctx, nil)
		}
	}
}

func (handler *Handler) GetEmployerProfile(ctx *gin.Context) notifications.Notification {
	if emp, file, err := users.GetEmployer(getTokenByContext(ctx), DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, emp, file)
	}
}

//TODO : handle projectImage, handle proper responses
func (handler *Handler) AddEmployerProject(ctx *gin.Context) notifications.Notification {
	form := data.ProjectForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	mForm, err := ctx.MultipartForm()
	if err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	attachmentHeaders := mForm.File[ProjectFilesFormName]
	var project existence.Project
	if err := json.Unmarshal([]byte(form.Project), &project); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
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
	if err := users.AddProjectToEmployer(getTokenByContext(ctx), project, attachments, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
		//TODO
		/*			if err.Error() == "project fields not valid" {
						return notifications.GetExpectationFailedError(ctx, newToken, nil)
					} else {
						return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
					}*/
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) EditEmployerProject(ctx *gin.Context) notifications.Notification {
	project := existence.Project{}
	if err := ctx.ShouldBindJSON(&project); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.EditEmployerProject(getTokenByContext(ctx), project, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) AssignProjectToFreelancer(ctx *gin.Context) notifications.Notification {
	assign := struct {
		Id         string `json:"id"`
		Freelancer string `json:"freelancer"`
	}{}
	if err := ctx.ShouldBindJSON(&assign); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.AssignProjectToFreelancer(getTokenByContext(ctx), assign.Freelancer, assign.Id, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
		//TODO
		/*			switch err.Error() {
					case "not valid freelancer":
						return notifications.GetExpectationFailedError(ctx, newToken, nil)
					case "not valid token for this project":
						return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
					default:
						return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)

					}*/
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) ExtendProject(ctx *gin.Context) notifications.Notification {
	extend := struct {
		Id         string    `json:"id"`
		FinishDate time.Time `json:"finish-date"`
	}{}
	if err := ctx.ShouldBindJSON(&extend); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.ExtendProject(getTokenByContext(ctx), extend.Id, extend.FinishDate, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
		//TODO
		/*			switch err.Error() {
					case "not valid token for this project":
						return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
					case "not valid time", "not valid project status":
						return notifications.GetExpectationFailedError(ctx, newToken, nil)
					default:
						return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
					}*/
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}

func (handler *Handler) CloseProject(ctx *gin.Context) notifications.Notification {
	close := struct {
		Id string `json:"id"`
	}{}
	if err := ctx.ShouldBindJSON(&close); err != nil {
		return notifications.GetShouldBindJsonErrorNotif(ctx, nil)
	}
	if err := users.CloseProject(getTokenByContext(ctx), close.Id, DB); err != nil {
		return notifications.GetInternalServerErrorNotif(ctx, nil)
		//TODO
		/*			switch err.Error() {
					case "not valid token for this project":
						return notifications.GetTokenNotAuthorizedErrorNotif(ctx, nil)
					case "not valid time", "not valid project status":
						return notifications.GetExpectationFailedError(ctx, newToken, nil)
					default:
						return notifications.GetInternalServerErrorNotif(ctx, newToken, nil)
					}*/
	} else {
		return notifications.GetSuccessfulNotif(ctx, nil)
	}
}
