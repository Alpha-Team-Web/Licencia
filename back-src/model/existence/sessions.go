package existence

import "time"

type AuthToken struct {
	//UnderLayingToken string    `json:"under-laying-token" sql:",unique"`
	Token        string    `json:"token" sql:",pk,notnull"`
	Username     string    `json:"username" sql:",notnull"`
	InitialTime  time.Time `json:"initial-time" sql:"init_time,notnull"`
	IsFreelancer bool      `json:"is-freelancer" sql:",notnull"`
	IsUsed       bool      `json:"is-used" sql:",notnull"`
}
