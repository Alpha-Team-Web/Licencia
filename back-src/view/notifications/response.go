package notifications

import "time"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListicProject struct {
	Id                  string    `json:"id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	StartDate           time.Time `json:"start-date"`
	NumberOfSuggestions int       `json:"number_of_suggestions"`
	Skills              []string  `json:"skills"`
	EmployerUsername    string    `json:"employer_username"`
	EmployerShownName   string    `json:"employer_shown_name"`
}
