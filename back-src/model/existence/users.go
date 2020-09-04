package existence

const (
	EmployerType   = "employer"
	FreelancerType = "freelancer"
)

type user struct {
	Username    string `json:"username" binding:"min=4,max=20" sql:",pk,notnull"`
	Password    string `json:"password" binding:"min=6,max=20" sql:",notnull"`
	ShownName   string `json:"shown-name" binding:"min=0,max=30" sql:",notnull"`
	Email       string `json:"email" binding:"email,min=5,max=100" sql:"email,unique"`
	Description string `json:"description" binding:"min=0,max=500" sql:"description"`
}

type person struct {
	FirstName   string `json:"first-name" binding:"max=50"`
	LastName    string `json:"last-name" binding:"max=50"`
	PhoneNumber string `json:"phone-number" binding:"max=30"`
	Address     string `json:"address" binding:"max=100"`
}

type Employer struct {
	tableName struct{} `sql:"employers"`
	user
	person
	ProjectIds []string `json:"project-ids"`
}

const (
	FreelancerGold   = "gold"
	FreelancerSilver = "silver"
	FreelancerBronze = "bronze"
)

const (
	GoldRequestSize   = 15
	SilverRequestSize = 8
	BronzeRequestSize = 3
)

type Freelancer struct {
	tableName struct{} `sql:"freelancers"`
	user
	person
	ProjectIds            []string            `json:"project-ids"`
	RequestedProjectIds   []string            `json:"requested-project-ids"`
	AccountType           string              `json:"account-type" sql:",notnull"`
	Website               string              `json:"website"`
	GithubAccount         string              `json:"github-account"`
	GithubRepos           []string            `json:"github-repos"`
	SkillsWithProject     map[string][]string `json:"skills-with-project"`
	ChosenFieldWithSkills map[string][]string `json:"chosen-field-with-skills"`
}
