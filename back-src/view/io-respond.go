package view

import (
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RespondRegister(context *gin.Context, err error) {

	if err == nil {
		context.JSON(http.StatusOK, responses.SuccessMessage)
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

func RespondLogin(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.SuccessMessage)
	} else {
		context.Header("Token", "N/A")
		var status int
		switch {
		case strings.Contains(err.Error(), "invalid query: "):
			status = http.StatusExpectationFailed
		case strings.Contains(err.Error(), "not signed up email: "), strings.Contains(err.Error(), "not signed up username: "):
			status = http.StatusBadRequest
		case strings.Contains(err.Error(), "invalid password: "):
			status = http.StatusMethodNotAllowed
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, responses.Response{Message: err.Error()})
	}
}
