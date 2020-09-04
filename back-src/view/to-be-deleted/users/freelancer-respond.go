package users

import (
	"back-src/model/existence"
	"back-src/view/notifications"
	"back-src/view/responses"
	"back-src/view/to-be-deleted"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondFreelancerEdit(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, notifications.Response{Message: "Successful"})
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !to_be_deleted.RespondDataValidationErrors(context, err) {
				var status int = http.StatusInternalServerError
				context.JSON(status, notifications.Response{Message: err.Error()})
			}
		}
	}
}

func RespondFreelancerGetProfile(context *gin.Context, token string, frl existence.Freelancer, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, frl)
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			//TODO : add switch cases if there are other types of error
			var status int = http.StatusInternalServerError
			context.JSON(status, notifications.Response{Message: err.Error()})
		}
	}
}

func RespondFreelancerRequestToProject(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.SuccessMessage)
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			if !to_be_deleted.RespondDataValidationErrors(context, err) {
				//TODO : add switch cases if there are other types of error
				var status int
				switch {
				case strings.Contains(err.Error(), "cant request more"):
					status = http.StatusMethodNotAllowed
				case strings.Contains(err.Error(), "invalid project id"):
					status = http.StatusExpectationFailed
				case strings.Contains(err.Error(), "project status not suitable"):
					status = http.StatusBadRequest
				default:
					status = http.StatusInternalServerError
				}
				context.JSON(status, notifications.Response{Message: err.Error()})
			}
		}
	}
}
