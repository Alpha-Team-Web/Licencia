package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Notification struct {
	Context    *gin.Context
	Token      string
	Message    string
	StatusCode int
	Data       interface{}
}

func GetShouldBindJsonErrorNotif(ctx *gin.Context, token string) Notification {
	ctx.JSON()
	notif := Notification{
		Context:    ctx,
		Token:      token,
		Message:    "data invalid",
		StatusCode: http.StatusBadRequest, //400
	}
	return notif
}

func GetDatabaseErrorNotif(ctx *gin.Context, token string) Notification {
	notif := Notification{
		Context:    ctx,
		Token:      token,
		Message:    "database down",
		StatusCode: http.StatusServiceUnavailable, //503
	}
	return notif
}

func GetInvalidQueryErrorNotif(ctx *gin.Context, token string) Notification {
	notif := Notification{
		Context:    ctx,
		Token:      token,
		Message:    "database down",
		StatusCode: http.StatusNotAcceptable, //406
	}
	return notif
}

func GetSuccessfulNotif(ctx *gin.Context, token string) Notification {
	notif := Notification{
		Context:    ctx,
		Token:      token,
		Message:    "successful",
		StatusCode: http.StatusOK, //200
	}
	return notif
}

func GetInternalServerErrorNotif(ctx *gin.Context, token string) Notification {
	notif := Notification{
		Context:    ctx,
		Token:      token,
		Message:    "server down",
		StatusCode: http.StatusInternalServerError, //500
	}
	return notif
}

func GetTokenNotAuthorizedErrorNotif(ctx *gin.Context) Notification {
	notif := Notification{
		Context:    ctx,
		Token:      "N/A",
		Message:    "not authorized token",
		StatusCode: http.StatusForbidden, //403
	}
	return notif
}
