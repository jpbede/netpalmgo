package models

// SetConfigRequest represents a request for /setconfig
type SetConfigRequest struct {
	Library        Library        `json:"library"`
	ConnectionArgs ConnectionArgs `json:"connection_args"`
	Config         []string       `json:"config,omitempty"`
	J2Config       *J2Config      `json:"j2config,omitempty"`
	Args           *SetConfigArgs `json:"args,omitempty"`
	QueueStrategy  QueueStrategy  `json:"queue_strategy"`
	PreChecks      []PrePostCheck `json:"pre_checks,omitempty"`
	PostChecks     []PrePostCheck `json:"post_checks,omitempty"`
	EnableMode     bool           `json:"enable_mode,omitempty"`
}

// J2Config jinja2 template and arguments
type J2Config struct {
	Template string                 `json:"template"`
	Args     map[string]interface{} `json:"args"`
}

// SetConfigArgs arguments for the SetConfigRequest
type SetConfigArgs struct {
	Payload          interface{} `json:"payload,omitempty"`
	DefaultOperation string      `json:"default_operation,omitempty"`
	Target           string      `json:"target,omitempty"`
	Config           string      `json:"config,omitempty"`
	URI              string      `json:"uri,omitempty"`
	Action           string      `json:"action,omitempty"`
}
