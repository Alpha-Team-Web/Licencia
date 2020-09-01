package responses

type Response struct {
	Message string `json:"message"`
}

var SuccessMessage = Response{"successful"}
