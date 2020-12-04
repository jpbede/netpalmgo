package models

// Library that would be used for connecting
type Library string

const (
	// LibraryNetmiko use netmiko for connection
	LibraryNetmiko Library = "netmiko"
	// LibraryNetmiko use napalm for connection
	LibraryNapalm Library = "napalm"
	// LibraryNetmiko use ncclient for connection
	LibraryNCClient Library = "ncclient"
	// LibraryNetmiko use restconf for connection
	LibraryRestConf Library = "restconf"
)
