package control

import (
	"back-src/model/database"
	"fmt"
	"time"
)

type Control struct {
}

const notUsedExpiry = 1

var DB *database.Database

func NewControl() *Control {
	DB = database.NewDb()
	err := DB.Initialize()
	if err != nil {
		panic(err)
	}
	return &Control{}
}

func (controller *Control) AddNewClock(token string) {
	clk := clock{notUsedExpiry, func() {
		controller.checkTokenUse(token)
	}}
	clk.startWorking()
}

func (controller *Control) checkTokenUse(token string) {
	if isUsed, err := DB.IsAuthUsed(token); err != nil {
		panic(err)
	} else if isUsed {
		if err := DB.ChangeAuthUsage(token, false); err != nil {
			panic(err)
		}
		controller.AddNewClock(token)
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
		fmt.Println("Clock Started")
		clock.tik()
		fmt.Println("Timer Finished")
		clock.job()
	}()
}

func (clock *clock) tik() {
	for i := 0; i < clock.minutes; i++ {
		time.Sleep(time.Minute)
	}
}
