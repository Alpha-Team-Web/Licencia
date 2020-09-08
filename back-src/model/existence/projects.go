package existence

import "time"

/*--------------------------------------Review--------------------------------------*/

type booleanReviewFreelancer struct {
	DeliveredCompletely         bool `json:"delivered-completely"`
	DeliveredInTime             bool `json:"delivered-in-time"`
	UsedThePromisedTechnologies bool `json:"used-the-promised-technologies"`
	InformedProperly            bool `json:"informed-properly"`
}

type booleanReviewEmployer struct {
	PayProperly          bool `json:"pay-properly"`
	GotProjectBackInTime bool `json:"got-project-back-in-time"`
	QuestionedProperly   bool `json:"questioned-properly"`
	HelpedProperly       bool `json:"helped-properly"`
}

type EmployerFreelancerReview struct {
	ProjectID string `json:"project-id" sql:",pk"`

	EmployerToFreeLancerScore   int    `json:"employer-to-freelancer-score" binding:"lte=1,gte=5"`
	EmployerToFreelancerComment string `json:"employer-to-freelancer-comment"  binding:"min=0,max=500"`
	WasFreelancerCommitted      bool   `json:"was-freelancer-committed"`
	booleanReviewFreelancer
}

type FreelancerEmployerReview struct {
	ProjectID string `json:"project-id" sql:",pk"`

	FreelancerToEmployerScore   int    `json:"freelancer-to-employer-score" binding:"lte=1,gte=5"`
	FreelancerToEmployerComment string `json:"freelancer-to-employer-comment"  binding:"min=0,max=500"`
	WasEmployerCommitted        bool   `json:"was-freelancer-committed"`
	booleanReviewEmployer
}

/*--------------------------------------Project--------------------------------------*/

//Project Stats
const (
	Open    = "open"
	Closed  = "closed"
	OnGoing = "ongoing"
)

type Project struct {
	Id string `json:"id" sql:",pk"`

	//main info
	Name             string              `json:"name" binding:"min=0,max=30"`
	Description      string              `json:"description" binding:"min=0,max=500"`
	MinBudget        float64             `json:"min-budget"`
	MaxBudget        float64             `json:"max-budget"`
	FieldsWithSkills map[string][]string `json:"fields-with-skills"`
	ProjectStatus    string              `json:"project-status"`

	//user
	EmployerUsername                  string            `json:"employer"`
	FreelancerUsername                string            `json:"freelancer"`
	FreelancerRequestsWithDescription map[string]string `json:"freelancer-requests-with-description"`

	//time
	Duration   time.Duration `json:"duration"`
	InitDate   time.Time     `json:"init-date"`
	StartDate  time.Time     `json:"start-date"`
	FinishDate time.Time     `json:"finish-date"`

	//files
	FileIds []string `json:"file_ids"`
}

/*--------------------------------------Sample--------------------------------------*/

type OuterSample struct {
	Id               string              `json:"id"`
	Name             string              `json:"name" binding:"min=0,max=30"`
	Description      string              `json:"description"`
	FieldsWithSkills map[string][]string `json:"fields-with-skills"`
	GithubRepo       string              `json:"github-repo"`
	TitlesWithLinks  map[string][]string `json:"titles-links"`
}

type LicenciaSample struct {
	tableName struct{} `sql:"samples"`
	OuterSample
	ProjectId  string `json:"project-id"`
	IsLicencia bool   `json:"is-licencia"`
}
