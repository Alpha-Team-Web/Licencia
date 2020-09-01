package existence

type Field struct {
	Id     string   `json:"id" binding:"min=1,max=6" pg:"id,pk"`
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}
