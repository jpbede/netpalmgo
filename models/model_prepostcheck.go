package models

type PrePostCheck struct {
	MatchType     Check         `json:"match_type"`
	MatchStr      []interface{} `json:"match_str"`
	GetConfigArgs GetConfigArgs `json:"get_config_args"`
}

type GetConfigArgs struct {
	Command string `json:"command"`
}
