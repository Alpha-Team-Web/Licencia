package handle

import (
	"back-src/model/database"
	"time"
)

type Handler struct {
}

const notUsedExpiry = 15

var AuthExpiryDur time.Duration

var DB *database.Database

func NewControl() *Handler {
	var error error
	AuthExpiryDur, error = time.ParseDuration("30m")
	if error != nil {
		panic(error)
	}
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Handler{}
}

func (handler *Handler) AddNewClock(token string) {
	clk := clock{notUsedExpiry, func() {
		handler.checkTokenUse(token)
	}}
	clk.startWorking()
}

func (handler *Handler) checkTokenUse(token string) {
	if isUsed, err := DB.IsAuthUsed(token); err != nil {
		panic(err)
	} else if isUsed {
		if err := DB.ChangeAuthUsage(token, false); err != nil {
			panic(err)
		}
		handler.AddNewClock(token)
	} else {
		if err := DB.ExpireAuth(token); err != nil {
			panic(err)
		}
	}
}

type clock struct {
	minutes int
	job     func()
}

func (clock *clock) startWorking() {
	go func() {
		clock.tik()
		clock.job()
	}()
}

func (clock *clock) tik() {
	for i := 0; i < clock.minutes; i++ {
		time.Sleep(time.Minute)
	}
}
