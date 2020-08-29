package existence

type user struct {
	Username    string `json:"username" binding:"min=4,max=20" pg:"username,pk,notnull"`
	Password    string `json:"password" binding:"min=6,max=20" pg:"password,notnull"`
	Email       string `json:"email" binding:"email" pg:"email,unique"`
	Description string `json:"description" pg:"description"`
}

type person struct {
	FirstName   string `json:"firstname" pg:"firstname"`
	LastName    string `json:"lastname" pg:"lastname"`
	PhoneNumber string `json:"phonenumber" pg:"phonenumber"`
	Address     string `json:"addr" pg:"addr"`
}

type image struct {
	Title string `json:"title" pg:"title"`
	Url   string `json:"url" pg:"url"`
}

type Employer struct {
	tableName struct{} `pg:"employers"`
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
	tableName struct{} `pg:"freelancers"`
	user
	person
	image
	ProjectIds            []string            `json:"project-ids"`
	AccountType           string              `json:"account-type"`
	Website               string              `json:"website"`
	GithubAccount         string              `json:"github"`
	GithubRepos           []string            `json:"github-repos"`
	SkillsWithProject     map[string][]string `json:"skills-with-project"`
	ChosenFieldWithSkills map[string][]string `json:"chosen-field-with-skills"`
}
