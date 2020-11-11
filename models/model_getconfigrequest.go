package models

type GetConfigRequest struct {
	Library        Library           `json:"library"`
	ConnectionArgs ConnectionArgs    `json:"connection_args"`
	Command        string            `json:"command"`
	Args           map[string]string `json:"args,omitempty"`
	QueueStrategy  QueueStrategy     `json:"queue_strategy"`
	Cache          Cache             `json:"config"`
}

// UseTextFSM sets the argument "use_textfsm" for pre-formatted results
func (gcr *GetConfigRequest) UseTextFSM() {
	if gcr.Args == nil {
		gcr.Args = make(map[string]string)
	}

	gcr.Args["use_textfsm"] = "true"
}
