package files

import (
	"back-src/view"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondUploadProfileImage(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{"Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			status := http.StatusInternalServerError
			context.JSON(status, responses.Response{Message: err.Error()})
		}
	}
}
