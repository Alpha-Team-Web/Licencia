package handle

import (
	licnecia_errors "back-src/controller/control/licencia-errors"
	"back-src/controller/control/projects/filters"
	"back-src/model/database"
	"back-src/view/api/handle/utils"
	"back-src/view/notifications"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Handler struct {
}

const NotAssignedToken = "N/A"

const notUsedExpiry = 10
const authExpiryMin = 30

var AuthExpiryDur time.Duration

var DB *database.Database
var TokensWithClocks map[string]*utils.Clock

func NewControl() *Handler {
	var error error
	AuthExpiryDur, error = time.ParseDuration(strconv.Itoa(authExpiryMin) + "m")
	if error != nil {
		panic(error)
	}
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	filters.Inv = filters.NewEngine(DB)
	initTokensWithClocks()
	return &Handler{}
}

func initTokensWithClocks() {
	TokensWithClocks = map[string]*utils.Clock{}
	if auths, err := DB.AuthTokenTable.GetAllTokens(); err == nil {
		for _, auth := range auths {
			AddNewClock(auth.Token)
		}
	} else {
		fmt.Println("ERROR:", "Server Could Not Init Previous Auth Tokens")
	}
}

func AddNewClock(token string) {
	clk := utils.NewClock(true, notUsedExpiry, func() {
		if err := DB.AuthTokenTable.ExpireAuth(token); err != nil {
			panic(err)
		}
	})
	clk.Start()
	TokensWithClocks[token] = clk
}

func KillClockIfExists(token string) bool {
	clock, ok := TokensWithClocks[token]
	if ok {
		clock.Stop()
	}
	delete(TokensWithClocks, token)
	return ok
}

func makeOperationErrorNotification(ctx *gin.Context, err error) notifications.Notification {
	if licnecia_errors.IsLicenciaError(err) {
		return notifications.GetExpectationFailedError(ctx, NotAssignedToken, licnecia_errors.GetErrorStrForRespond(err), nil)
	} else {
		return notifications.GetInternalServerErrorNotif(ctx, NotAssignedToken, nil)
	}
}

func getTokenByContext(ctx *gin.Context) string {
	return ctx.Writer.Header().Get("Token")
}
