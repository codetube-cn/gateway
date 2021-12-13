package models

import "gorm.io/gorm"

type FilterExtra struct {
	Args         []ExtraArg
	Options      []ExtraOption
	DefaultValue ExtraDefaultValue
}

type Filter struct {
	gorm.Model
	GatewayId  uint
	Name       string
	Code       string
	InputType  int
	ValueType  string
	Extra      *map[string]interface{} `gorm:"type:json"`
	SortNumber int
}
