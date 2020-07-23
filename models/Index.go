package models

import "time"

type Alive struct {
	Alive bool `json:"alive"`
	Service string `json:"service"`
	Date time.Weekday `json:"date"`
}
