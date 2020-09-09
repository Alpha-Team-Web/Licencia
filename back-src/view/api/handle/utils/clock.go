package utils

import (
	"time"
)

type Clock struct {
	On      bool
	minutes int
	job     func()
}

func NewClock(on bool, minutes int, job func()) *Clock {
	return &Clock{On: on, minutes: minutes, job: job}
}

func (clock *Clock) Start() {
	go func() {
		clock.tik()
		if clock.On {
			clock.job()
		}
	}()
}

func (clock *Clock) Stop() {
	clock.job = func() {
		//:)
	}
	clock.On = false
}

func (clock *Clock) tik() {
	for i := 0; (i < (clock.minutes * 60)) && clock.On; i++ {
		time.Sleep(time.Second)
	}
}
