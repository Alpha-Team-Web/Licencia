package existence

import "time"

type File struct {
	ProjectID string
	Title     string
	Url       string
}

//Project Stats
const (
	Open    = "Open"
	Closed  = "Closed"
	OnGoing = "OnGoing"
)

type Project struct {
	Id string `json:"id" sql:",pk"`

	//main info
	Name             string              `json:"name" binding:"min=0,max=30"`
	Description      string              `json:"desc" binding:"min=0,max=500"`
	MinBudget        float64             `json:"min-budget"`
	MaxBudget        float64             `json:"max-budget"`
	FieldsWithSkills map[string][]string `json:"fields-skills"`
	ProjectStatus    string              `json:"stat"`

	//user
	EmployerUsername                  string            `json:"employer"`
	FreelancerUsername                string            `json:"freelancer"`
	FreelancerRequestsWithDescription map[string]string `json:"frl-req-with-desc"`

	//time
	Duration   time.Duration `json:"dur"`
	StartDate  time.Time     `json:"start"`
	FinishDate time.Time     `json:"finish"`

	//review
	FreelancerToEmployerScore   int    `json:"fr-to-emp-score"`
	FreelancerToEmployerComment string `json:"fr-to-emp-comment"  binding:"min=0,max=500"`
	EmployerToFreeLancerScore   int    `json:"emp-to-fr-score"`
	EmployerToFreelancerComment string `json:"emp-to-fr-comment"  binding:"min=0,max=500"`
}

type OuterSample struct {
	Id               string              `json:"id"`
	Name             string              `json:"name" binding:"min=0,max=30"`
	Description      string              `json:"desc"`
	FieldsWithSkills map[string][]string `json:"fields-skills"`
	GithubRepo       string              `json:"github"`
	TitlesWithLinks  map[string][]string `json:"titles-links"`
}

type LicenciaSample struct {
	tableName struct{} `sql:"samples"`
	OuterSample
	ProjectId  string `json:"project-id"`
	IsLicencia bool   `json:"is-licencia"`
}
