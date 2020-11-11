package models

type Cache struct {
	Enabled bool
	TTL     int
	Poison  bool
}
