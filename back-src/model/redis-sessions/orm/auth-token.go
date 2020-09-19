package orm

import (
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"encoding/json"
	"time"
)

func HashAuthToken(auth existence.AuthToken) map[string]interface{} {
	bytes, _ := json.Marshal(auth.InitialTime)
	hash := map[string]interface{}{
		"token":        auth.Token,
		"username":     auth.Username,
		"initialTime":  string(bytes),
		"isFreelancer": libs.Ternary(auth.IsFreelancer, "1", "0").(string),
	}
	return hash
}

func UnHashAuthToken(hash map[string]string) existence.AuthToken {
	var initialTime time.Time
	json.Unmarshal([]byte(hash["initialTime"]), &initialTime)
	return existence.AuthToken{
		Token:        hash["token"],
		Username:     hash["username"],
		InitialTime:  initialTime,
		IsFreelancer: hash["isFreelancer"] == "1",
	}
}
