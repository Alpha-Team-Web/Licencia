package view

import (
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondRegister(context *gin.Context, err error) {

	if err == nil {
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		var status int
		switch {
		case strings.Contains(err.Error(), "invalid query: "):
			status = http.StatusExpectationFailed
		case strings.Contains(err.Error(), "duplicate username: "), strings.Contains(err.Error(), "duplicate email: "):
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}

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
