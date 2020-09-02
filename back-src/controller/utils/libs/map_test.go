package libs

import "testing"

func TestMapCount(t *testing.T) {
	myMap := map[string]bool{
		"ashkan": true,
		"sultan": false,
	}
	if len(myMap) != 2 {
		t.Error()
	}
}
