package models

// Cache represents a config for the task how a result should be cached
type Cache struct {
	Enabled bool `json:"enabled"`
	TTL     int  `json:"ttl"`
	Poison  bool `json:"poison"`
}
