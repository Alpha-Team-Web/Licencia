package existence

type Field struct {
	ID     string   `json:"id" binding:"min=1,max=6" pg:"id,pk"`
	Name   string   `json:"name" pg:"name"`
	Skills []string `json:"skills" pg:"skills"`
}
