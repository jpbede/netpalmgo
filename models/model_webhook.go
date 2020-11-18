package models

type Webhook struct {
	Name       string                 `json:"name"`
	J2Template string                 `json:"j2template"`
	Args       map[string]interface{} `json:"args"`
}
