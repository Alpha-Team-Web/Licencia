package projects

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type withMap struct {
	Mp map[string][]string
}

type timeSt struct {
	T  time.Duration
	T2 time.Time
}

func TestShowProject(t *testing.T) {
	thing := withMap{map[string][]string{}}
	thing.Mp["foo"] = []string{"a", "b", "c"}
	bytes, err := json.Marshal(thing)
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	fmt.Print(string(bytes))
	timeTmp := time.Now()
	bytes, err = json.Marshal(timeTmp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	fmt.Print(string(bytes))
	dur, err := time.ParseDuration("30m")
	if err != nil {
		t.Error(err)
	}

	bytes2, err := json.Marshal(&dur)
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	fmt.Println(string(bytes2))
	structTime := timeSt{}
	structTime.T = dur
	structTime.T2 = timeTmp
	bytes3, err := json.Marshal(&structTime)

	if err != nil {
		t.Error(err)
	}
	fmt.Println("Here :")
	fmt.Println(string(bytes3))
}
