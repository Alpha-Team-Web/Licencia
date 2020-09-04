package to_be_deleted

import (
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Respond(notification notifications.Notification) {
	ctx := notification.Context
	ctx.Header("Token", notification.Token)
	ctx.JSON(notification.StatusCode, notification.Response{Message: notification.Message, Data: notification.Data})
}

//true for when auth token has happened and the respond is sent
func RespondTokenErrors(context *gin.Context, err error) bool {
	var status int
	switch {
	case strings.Contains(err.Error(), "not authorized token: "):
		status = http.StatusBadRequest
	case strings.Contains(err.Error(), "wrong user type token: "):
		status = http.StatusConflict
	}
	if status != 0 {
		context.Header("Token", "N/A")
		context.JSON(status, notifications.Response{Message: err.Error()})
		return true
	}
	return false
}

func RespondDataValidationErrors(context *gin.Context, err error) bool {
	errorMessage := err.Error()
	if strings.Contains(errorMessage, " tag") && strings.Contains(errorMessage, " validation for ") && strings.Contains(errorMessage, "Key: ") {
		context.JSON(http.StatusBadRequest, notifications.Response{Message: errorMessage})
		return true
	}
	return false
}
