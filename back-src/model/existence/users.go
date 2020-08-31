package existence

const (
	EmployerType   = "employer"
	FreelancerType = "freelancer"
)

type user struct {
	Username    string `json:"username" binding:"min=4,max=20" sql:"username,pk,notnull"`
	Password    string `json:"password" binding:"min=6,max=20" sql:"password,notnull"`
	ShownName   string `json:"shown-name" binding:"min=4,max=30" sql:",notnull"`
	Email       string `json:"email" binding:"email,min=5,max=100" sql:"email,unique"`
	Description string `json:"description" sql:"description"`
}

type person struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"addr"`
}

type image struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Employer struct {
	tableName struct{} `sql:"employers"`
	user
	person
	image
	ProjectIds []string `json:"project-ids"`
}

const (
	FreelancerGold   = "Gold"
	FreelancerSilver = "Silver"
	FreelancerBronze = "Bronze"
)

type Freelancer struct {
	tableName struct{} `sql:"freelancers"`
	user
	person
	image
	ProjectIds            []string            `json:"project-ids"`
	RequestedProjectIds   []string            `json:"req-project-ids"`
	AccountType           string              `json:"account-type"`
	Website               string              `json:"website"`
	GithubAccount         string              `json:"github"`
	GithubRepos           []string            `json:"github-repos"`
	SkillsWithProject     map[string][]string `json:"skills-with-project"`
	ChosenFieldWithSkills map[string][]string `json:"chosen-field-with-skills"`
}
