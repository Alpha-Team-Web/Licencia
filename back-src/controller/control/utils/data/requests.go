package data

type LoginRequest struct {
	Id           string `json:"id" binding:"min=4,max=100"`
	Password     string `json:"password" binding:"min=6,max=20"`
	IsFreelancer bool   `json:"-"`
}
