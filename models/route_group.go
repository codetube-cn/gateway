package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

type RouteGroup struct {
	gorm.Model
	GatewayId  uint
	Name       string
	UriPrefix  string
	Predicates model.JSON
	Filters    model.JSON
	SortNumber int
}
