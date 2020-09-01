package data

type LoginRequest struct {
	Id           string `json:"id" binding:"min=4,max=100"`
	Password     string `json:"password" binding:"min=6,max=20"`
	IsFreelancer bool   `json:"-"`
}

type ChangePassRequest struct {
	OldPass string `json:"old-pass" binding:"min=6,max=20"`
	NewPass string `json:"new-pass" binding:"min=6,max=20"`
}

type FreelancerRequestForProject struct {
	Id          string `json:"id" binding:"min=15,max=15"`
	Description string `json:"description" binding:"min:0,max=500"`
}
