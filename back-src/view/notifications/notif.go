package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Notification struct {
	Context    *gin.Context
	Message    string
	StatusCode int
	Data       interface{}
}

func GetShouldBindJsonErrorNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "data invalid",
		StatusCode: http.StatusBadRequest, //400
		Data:       data,
	}
	return notif
}

func GetDatabaseErrorNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "sql down",
		StatusCode: http.StatusServiceUnavailable, //503
		Data:       data,
	}
	return notif
}

func GetInMemoryDataStructureDownNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "redis down",
		StatusCode: http.StatusServiceUnavailable, //503
		Data:       data,
	}
	return notif
}

func GetInvalidQueryErrorNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "Invalid Query",
		StatusCode: http.StatusNotAcceptable, //406
		Data:       data,
	}
	return notif
}

func GetSuccessfulNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "successful",
		StatusCode: http.StatusOK, //200
		Data:       data,
	}
	return notif
}

func GetInternalServerErrorNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "server down",
		StatusCode: http.StatusInternalServerError, //500
		Data:       data,
	}
	return notif
}

func GetTokenNotAuthorizedErrorNotif(ctx *gin.Context, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    "not authorized token",
		StatusCode: http.StatusForbidden, //403
		Data:       data,
	}
	return notif
}

func GetExpectationFailedError(ctx *gin.Context, errStr string, data ...interface{}) Notification {
	notif := Notification{
		Context:    ctx,
		Message:    errStr,
		StatusCode: http.StatusExpectationFailed, //417
		Data:       data,
	}
	return notif
}
