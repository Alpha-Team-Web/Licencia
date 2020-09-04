package existence

type File struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

const (
	FreelancerProfile = "freelancer"
	EmployerProfile   = "employer"
	ProjectProfile    = "project"
)

type Profile struct {
	Id   string `json:"id" sql:",nopk"`
	Type string `json:"type"`
	File
}
