package main

import (
	"back-src/routing"
)

func main() {
	lst := routing.NewRouter("8008")
	if err := lst.Listen(); err != nil {
		panic(err)
	}
}
