package data

type LoginRequest struct {
	Id           string `json:"id" binding:"min=4,max=100"`
	Password     string `json:"password" binding:"min=6,max=20"`
	IsFreelancer bool   `json:"-"`
}

type ChangePassRequest struct {
	OldPass string `json:"old_pass" binding:"min=6,max=20"`
	NewPass string `json:"new_pass" binding:"min=6,max=20"`
}
