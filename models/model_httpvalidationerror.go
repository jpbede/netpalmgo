package models

// ValidationError represents a error returned by API if request validation fails
type ValidationError struct {
	LOC  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// HTTPValidationError represents a error returned by API if some http validations fails eg. invalid token
type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}
