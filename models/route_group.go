package models

import "gorm.io/gorm"

type RouteGroup struct {
	gorm.Model
	GatewayId  uint
	Name       string
	UriPrefix  string
	Predicates *map[string]interface{} `gorm:"type:json"`
	Filters    *map[string]interface{} `gorm:"type:json"`
	SortNumber int
}
