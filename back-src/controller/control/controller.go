package control

import (
	"back-src/model/database"
	"time"
)

type Control struct {
	notUsedTokenExpiry int
	tokensWithClock    map[string]clock
}

var DB *database.Database

func NewControl() *Control {
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Control{1, map[string]clock{}}
}

func (controller *Control) AddNewClock(token string) {
	controller.tokensWithClock[token] = clock{controller.notUsedTokenExpiry, func() {
		controller.checkTokenUse(token)
		return
	}}
}

func (controller *Control) checkTokenUse(token string) {
	if isUsed, err := DB.IsAuthUsed(token); err != nil {
		panic(err)
	} else if isUsed {
		if err := DB.ChangeAuthUsage(token, false); err != nil {
			panic(err)
		}
	} else {

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
