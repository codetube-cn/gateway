package models

import "gorm.io/gorm"

type Gateway struct {
	gorm.Model
	Code string
	Name string
	Host string
	Port int
}