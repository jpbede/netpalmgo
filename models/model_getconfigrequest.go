package models

type GetConfigRequest struct {
	Library        Library           `json:"library"`
	ConnectionArgs ConnectionArgs    `json:"connection_args"`
	Command        string            `json:"command"`
	Args           map[string]string `json:"args"`
	QueueStrategy  QueueStrategy     `json:"queue_strategy"`
	Cache          Cache             `json:"config"`
}
