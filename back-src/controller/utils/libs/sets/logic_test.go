package sets

import (
	"back-src/controller/utils/libs"
	"fmt"
	"strconv"
	"testing"
)

func TestTernary(t *testing.T) {
	isThere := true
	fmt.Println(libs.Ternary(isThere, Check2, Check1).(func(int) string)(20))
}

func Check1(x int) string {
	return strconv.Itoa(x) + " Hello"
}
func Check2(x int) string {
	return "Fuck You"
}
