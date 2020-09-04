package projects

import (
	"back-src/view/notifications"
	"back-src/view/responses"
	"back-src/view/to-be-deleted"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondReview(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.SuccessMessage)
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			//TODO : add switch cases if there are other types of error
			var status int
			switch {
			case strings.Contains(err.Error(), "not involved in project the username: "):
				status = http.StatusMethodNotAllowed
			default:
				context.JSON(status, notifications.Response{Message: err.Error()})
			}
		}
	}
}

func RespondFilterProjects(context *gin.Context, projects []notifications.ListicProject, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, projects)
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			//TODO : add switch cases if there are other types of error
			var status int
			switch {
			case strings.Contains(err.Error(), "not involved in project the username: "):
				status = http.StatusMethodNotAllowed
			default:
				status = http.StatusInternalServerError
			}
			context.JSON(status, notifications.Response{Message: err.Error()})
		}
	}
}
