package models

type ValidationError struct {
	LOC  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}
