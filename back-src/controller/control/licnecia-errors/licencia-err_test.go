package licnecia_errors

import (
	"fmt"
	"testing"
)

func TestSubstring(t *testing.T) {
	testStr := "This Is A Test"
	err := MakeLicenciaError(testStr)
	if IsLicenciaError(err) {
		if GetErrorStrForRespond(err) == testStr {
			fmt.Println("Pass")
		} else {
			str := "\"" + GetErrorStrForRespond(err) + "\""
			t.Error(str)
		}
	} else {
		t.Error(err.Error())
	}
}
