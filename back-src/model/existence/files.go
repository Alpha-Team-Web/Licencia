package existence

type File struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
	Size int64  `json:"size"`
}

const (
	FreelancerProfile = "freelancer"
	EmployerProfile   = "employer"
	ProjectProfile    = "project"
)

type Profile struct {
	File
	Id   string `json:"id" sql:",nopk"`
	Type string `json:"type"`
}

type ProjectAttachment struct {
	File
	ProjectId string `json:"project-id"`
	FileId    string `json:"file-id" sql:",pk"`
}
