package libs

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var charset = "qwertyuiop[]{}asdfghjkl;:zxcvbnm,.<>1234567890-_=+!@#$%^&*()"
var numbSet = "0123456789"

func getRandom(length int, set string, predator func(string) bool) (random string) {
	for ok := true; ok; ok = predator(random) {
		random = func(length int, set string) string {
			b := make([]byte, length)
			for i := range b {
				b[i] = set[seededRand.Intn(len(set))]
			}
			return string(b)
		}(length, set)
	}
	return
}

func GetRandomString(length int, predator func(string) bool) string {
	return getRandom(length, charset, predator)
}

func GetRandomNumberAsString(length int, predator func(string) bool) string {
	return getRandom(length, numbSet, predator)
}
