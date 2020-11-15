package models

// Status of a request
type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)
