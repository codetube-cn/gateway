package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

type FilterExtra struct {
	Args struct {
		Args []ExtraArg
	}
	Options struct {
		Options []ExtraOption
	}
	DefaultValue struct {
		DefaultValue ExtraValue
	}
}

type Filter struct {
	gorm.Model
	GatewayId  uint
	Name       string
	Code       string
	InputType  int
	ValueType  string
	Extra      model.JSON `json:"queryParam"`
	SortNumber int
}
