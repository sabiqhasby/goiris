package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
