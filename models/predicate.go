package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

type PredicateExtra struct {
	Options []ExtraOption
	DefaultValue ExtraDefaultValue
}

type Predicate struct {
	gorm.Model
	GatewayId  uint
	Name       string
	Code       string
	InputType  int
	ValueType  string
	Extra      model.JSON `gorm:"type:json"`
	SortNumber int
}
