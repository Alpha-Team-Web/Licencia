package users

import (
	"back-src/model/existence"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondEmployerEditProfile(context *gin.Context, err error) {
	if err == nil {
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		var status int
		switch {
		case strings.Contains(err.Error(), "no user with such username :"):
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}

func RespondEmployerGetProfile(context *gin.Context, emp existence.Employer, err error) {
	if err == nil {
		context.JSON(http.StatusOK, emp)
	} else {
		var status int
		switch {
		case strings.Contains(err.Error(), "no user with such username :"):
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}

func RespondEmployerGetProjects(context *gin.Context, projects []existence.Project, err error) {
	if err == nil {
		context.JSON(http.StatusOK, struct {
			projects []existence.Project `json:"projects"`
		}{projects: projects})
	} else {
		var status int
		switch {
		case strings.Contains(err.Error(), "no user with such username :"):
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}
