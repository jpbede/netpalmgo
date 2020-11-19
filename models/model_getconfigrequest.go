package models

// GetConfigRequest represents a request for /getconfig
type GetConfigRequest struct {
	Library        Library                `json:"library"`
	ConnectionArgs ConnectionArgs         `json:"connection_args"`
	Command        []string               `json:"command"`
	Args           map[string]interface{} `json:"args,omitempty"`
	QueueStrategy  QueueStrategy          `json:"queue_strategy"`
	PostChecks     []PrePostCheck         `json:"post_checks,omitempty"`
	Cache          Cache                  `json:"config"`
}

// UseTextFSM sets the argument "use_textfsm" for pre-formatted results
func (gcr *GetConfigRequest) UseTextFSM() {
	if gcr.Args == nil {
		gcr.Args = make(map[string]interface{})
	}

	gcr.Args["use_textfsm"] = "true"
}
