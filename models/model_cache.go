package models

type Cache struct {
	Enabled bool `json:"enabled"`
	TTL     int  `json:"ttl"`
	Poison  bool `json:"poison"`
}
