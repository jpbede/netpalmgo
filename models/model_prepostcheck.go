package models

// PrePostCheck to validate config before or after change
type PrePostCheck struct {
	MatchType     Check         `json:"match_type"`
	MatchStr      []string      `json:"match_str"`
	GetConfigArgs GetConfigArgs `json:"get_config_args"`
}

// GetConfigArgs command config that is run by PrePostCheck
type GetConfigArgs struct {
	Command string `json:"command"`
}
