package models

// Library that would be used for connecting
type Library string

const (
	// LibraryNetmiko use netmiko for connection
	LibraryNetmiko Library = "netmiko"
	// LibraryNapalm use napalm for connection
	LibraryNapalm Library = "napalm"
	// LibraryNCClient use ncclient for connection
	LibraryNCClient Library = "ncclient"
	// LibraryRestConf use restconf for connection
	LibraryRestConf Library = "restconf"
)
