package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

// Route 路由模型
type Route struct {
	gorm.Model
	GatewayId  uint       //网关ID
	GroupId    uint       //路由分组ID
	RouteId    string     //路由标识ID
	Name       string     //名称
	Uri        string     //URI
	Predicates model.JSON //路由使用的断言
	Filters    model.JSON //路由使用的过滤器
	SortNumber uint       //排序序号
}

// RoutePredicate 路由断言
type RoutePredicate struct {
	Predicate string      //断言标识
	Value     interface{} //断言值
}

// RouteFilter 路由过滤器
type RouteFilter struct {
	Filter string      //过滤器标识
	Value  interface{} //过滤器值
}
