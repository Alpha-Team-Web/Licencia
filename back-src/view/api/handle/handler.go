package handle

import (
	"back-src/controller/control/projects/filters"
	"back-src/controller/control/users"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"back-src/view/api/handle/utils"
	"errors"
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
var tokensWithClocks map[string]*utils.Clock

func NewControl() *Handler {
	var error error
	AuthExpiryDur, error = time.ParseDuration(strconv.Itoa(authExpiryMin) + "m")
	if error != nil {
		panic(error)
	}
	tokensWithClocks = map[string]*utils.Clock{}
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	filters.Inv = filters.NewEngine(DB)
	return &Handler{}
}

func AddNewClock(token string) {
	clk := utils.NewClock(true, notUsedExpiry, func() {
		if err := DB.AuthTokenTable.ExpireAuth(token); err != nil {
			panic(err)
		}
	})
	clk.Start()
	tokensWithClocks[token] = clk
}

func KillClockIfExists(token string) bool {
	clock, ok := tokensWithClocks[token]
	if ok {
		clock.Stop()
	}
	delete(tokensWithClocks, token)
	return ok
}

func CheckToken(token, userType string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		if libs.XNor(auth.IsFreelancer, userType == existence.FreelancerType) {
			return reInitToken(auth)
		} else {
			return "", errors.New("wrong user type token: " + token)
		}
	} else {
		return "", err
	}
}

func CheckTokenIgnoreType(token string) (string, error) {
	if auth, err := formalCheckToken(token); err == nil {
		return reInitToken(auth)
	} else {
		return "", err
	}
}

func formalCheckToken(token string) (existence.AuthToken, error) {
	if isThereAuth, err := DB.AuthTokenTable.IsThereAuthWithToken(token); err != nil {
		return existence.AuthToken{}, err
	} else if isThereAuth {
		if auth, err := DB.AuthTokenTable.GetAuthByToken(token); err != nil {
			return existence.AuthToken{}, err
		} else {
			return auth, err
		}
	} else {
		return existence.AuthToken{}, errors.New("not authorized token: " + token)
	}
}

func reInitToken(auth existence.AuthToken) (string, error) {
	currentTime := time.Now()
	if currentTime.Sub(auth.InitialTime) > AuthExpiryDur {

		if err := DB.AuthTokenTable.ExpireAuth(auth.Token); err != nil {
			return "", err
		} else {
			newToken, err := users.MakeNewAuthToken(auth.Username, auth.IsFreelancer, DB)
			if err != nil {
				return "", err
			}
			KillClockIfExists(auth.Token)
			AddNewClock(newToken)
			return newToken, nil
		}
	} else {
		KillClockIfExists(auth.Token)
		AddNewClock(auth.Token)
		return auth.Token, nil
	}
}
