package files

import (
	"back-src/view/notifications"
	"back-src/view/to-be-deleted"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondUploadProfileImage(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, notifications.Response{"Successful"})
	} else {
		if !to_be_deleted.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			status := http.StatusInternalServerError
			context.JSON(status, notifications.Response{Message: err.Error()})
		}
	}
}
