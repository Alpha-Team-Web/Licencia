package existence

import "time"

type File struct {
	ProjectID string
	Title     string
	Url       string
}

type Project struct {
	Id                        string            `json:"id" pg:"id,pk"`
	Name                      string            `json:"name"`
	EmployerUsername          string            `json:"employer"`
	FreelancerUsername        string            `json:"freelancer"`
	Duration                  time.Duration     `json:"dur"`
	StartDate                 time.Time         `json:"start"`
	FinishDate                time.Time         `json:"finish"`
	Description               string            `json:"desc" binding:"min=0,max=500"`
	MinBudget                 float64           `json:"min-budget"`
	MaxBudget                 float64           `json:"max-budget"`
	FieldsWithSkills          map[string]string `json:"fields-skills"`
	ProjectStatus             string            `json:"stat"`
	FreelancerToEmployerScore int               `json:"fr-to-emp-score"`
	EmployerToFreeLancerScore int               `json:"emp-to-fr-score"`
}

//Project Stats
const (
	Open    = "Open"
	Closed  = "Closed"
	OnGoing = "OnGoing"
)
