package users

import (
	"back-src/model/existence"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondFreelancerEdit(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !RespondTokenErrors(context, err) {
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
}

func RespondFreelancerGetProfile(context *gin.Context, token string, frl existence.Freelancer, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, frl)
	} else {
		context.Header("Token", "N/A")
		var status int
		switch {
		case strings.Contains(err.Error(), "not authorized token: "):
			status = http.StatusBadRequest
		case strings.Contains(err.Error(), "wrong user type token: "):
			status = http.StatusConflict
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}
