package datastore

import "gorm.io/gorm"

type Site struct {
	gorm.Model
	Name string `json:"name"`
	URL  string `json:"url"`
}
