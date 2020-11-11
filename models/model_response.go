package models

type Response struct {
	Status Status       `json:"status"`
	Data   TaskResponse `json:"data"`
}
