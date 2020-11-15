package models

// Response is the overall response of the API
type Response struct {
	Status Status       `json:"status"`
	Data   TaskResponse `json:"data"`
}
