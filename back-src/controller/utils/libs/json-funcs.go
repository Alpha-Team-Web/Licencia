package libs

import "encoding/json"

func Marshal(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
