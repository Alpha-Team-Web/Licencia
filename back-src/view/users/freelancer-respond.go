package users

import (
	"back-src/model/existence"
	"back-src/view"
	"back-src/view/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondFreelancerEdit(context *gin.Context, token string, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, responses.Response{Message: "Successful"})
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			var status int = http.StatusInternalServerError
			context.JSON(status, responses.Response{Message: err.Error()})
		}
	}
}

func RespondFreelancerGetProfile(context *gin.Context, token string, frl existence.Freelancer, err error) {
	if err == nil {
		context.Header("Token", token)
		context.JSON(http.StatusOK, frl)
	} else {
		if !view.RespondTokenErrors(context, err) {
			context.Header("Token", token)
			//TODO : add switch cases if there are other types of error
			var status int = http.StatusInternalServerError
			context.JSON(status, responses.Response{Message: err.Error()})
		}
	}
}
