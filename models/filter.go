package models

import "gorm.io/gorm"

type Filter struct {
	gorm.Model
	GatewayId  int
	Name       string
	Code       string
	InputType  int
	ValueType  string
	Extra      *map[string]interface{} `gorm:"type:json"`
	SortNumber int
}
