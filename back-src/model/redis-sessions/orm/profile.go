package orm

import (
	"back-src/model/existence"
	"strconv"
)

func HashProfileImage(profile existence.Profile) map[string]interface{} {
	return map[string]interface{}{
		"name": profile.Name,
		"id":   profile.Id,
		"size": profile.Size,
		"type": profile.Type,
		"data": string(profile.Data),
	}
}

func UnHashProfileImage(hashed map[string]string) existence.Profile {
	profile := existence.Profile{
		Id:   hashed["id"],
		Type: hashed["type"],
	}
	profile.Name = hashed["name"]
	profile.Id = hashed["id"]
	profile.Size, _ = strconv.ParseInt(hashed["size"], 10, 64)
	profile.Data = []byte(hashed["data"])
	return profile
}
