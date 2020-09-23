package authentications

import (
	licencia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/view/api/handle"
	"back-src/view/api/respond"
	"back-src/view/notifications"
	"github.com/gin-gonic/gin"
	"time"
)

func GetCheckTokenHandlerFunc(userType string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Token")
		if newToken, err := checkToken(token, userType); err != nil {
			respond.Respond(notifications.GetTokenNotAuthorizedErrorNotif(context, nil))
			context.Header("Token", "N/A")
			context.Abort()
		} else {
			context.Header("Token", newToken)
			context.Next()
		}
	}
}

func GetCheckTokenIgnoreTypeHandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Token")
		if newToken, err := checkTokenIgnoreType(token); err != nil {
			if licencia_errors.IsLicenciaError(err) {
				respond.Respond(notifications.GetTokenNotAuthorizedErrorNotif(context, nil))
			} else {
				respond.Respond(notifications.GetInMemoryDataStructureDownNotif(context, nil))
			}
			context.Abort()
		} else {
			context.Header("Token", newToken)
			context.Next()
		}
	}
}

func checkToken(token, userType string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		if libs.XNor(auth.IsFreelancer, userType == existence.FreelancerType) {
			return reInitToken(auth)
		} else {
			return "", licencia_errors.NewLicenciaError("wrong user type token")
		}
	} else {
		return "", err
	}
}

func checkTokenIgnoreType(token string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		return reInitToken(auth)
	} else {
		return "", err
	}
}

func formalCheckToken(token string) (existence.AuthToken, error) {
	if isThereAuth, err := handle.RedisApi.AuthTokenDb.IsThereAuthWithToken(token); err != nil {
		return existence.AuthToken{}, err
	} else if isThereAuth {
		if auth, err := handle.RedisApi.AuthTokenDb.GetAuthByToken(token); err != nil {
			return existence.AuthToken{}, err
		} else {
			return auth, err
		}
	} else {
		return existence.AuthToken{}, licencia_errors.NewLicenciaError("not authorized token")
	}
}

func reInitToken(auth existence.AuthToken) (string, error) {
	currentTime := time.Now()
	if currentTime.Sub(auth.InitialTime) > handle.AuthExpiryDur {

		if err := handle.RedisApi.AuthTokenDb.ExpireAuth(auth.Token); err != nil {
			return "", err
		} else {
			newToken, err := users.MakeNewAuthToken(auth.Username, auth.IsFreelancer, handle.RedisApi)
			if err != nil {
				return "", err
			}
			handle.KillClockIfExists(auth.Token)
			handle.AddNewClock(newToken)
			return newToken, nil
		}
	} else {
		handle.KillClockIfExists(auth.Token)
		handle.AddNewClock(auth.Token)
		return auth.Token, nil
	}
}
