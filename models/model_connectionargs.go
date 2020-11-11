package models

type ConnectionArgs struct {
	DeviceType string `json:"device_type"`
	Host       string `json:"host"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
