package models

// Library that would be used for connecting
type Library string

const (
	LibraryNetmiko  Library = "netmiko"
	LibraryNapalm   Library = "napalm"
	LibraryNCClient Library = "ncclient"
	LibraryRestConf Library = "restconf"
)
