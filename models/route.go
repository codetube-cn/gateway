package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	GatewayId  int
	GroupId    int
	RouteId    string
	Name       string
	Uri        string
	Predicates *map[string]interface{} `gorm:"type:json"`
	Filters    *map[string]interface{} `gorm:"type:json"`
	SortNumber int
}
