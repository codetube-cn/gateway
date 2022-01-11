package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	GatewayId  uint
	GroupId    uint
	RouteId    string
	Name       string
	Uri        string
	Predicates model.JSON
	Filters    model.JSON
	SortNumber uint
}

type RoutePredicate struct {
	Predicate string
	Value interface{}
}

type RouteFilter struct {
	Filter string
	Value interface{}
}