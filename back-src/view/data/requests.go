package data

import "mime/multipart"

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

type Filter struct {
	Status            string   `json:"status"`
	MinPrice          float64  `json:"min-price"`
	MaxPrice          float64  `json:"max-price"`
	MustIncludeSkills []string `json:"must-include-skills"`
	IncludeSkills     []string `json:"include-skills"`
	ExcludeSkills     []string `json:"exclude-skills"`
	IsFilterBySkill   bool     `json:"is-filter-by-skill"`
}

type ProjectForm struct {
	Attachments []*multipart.FileHeader `form:"attachments" binding:"-"`
	Profile     *multipart.FileHeader   `form:"profile" binding:"-"`
	Project     string                  `form:"project" binding:"required"`
}
