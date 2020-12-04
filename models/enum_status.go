package models

// Status of a request
type Status string

const (
	// StatusSuccess request successfully
	StatusSuccess Status = "success"
	// StatusError request failed
	StatusError Status = "error"
)
