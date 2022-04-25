package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

// RouteGroup 路由分组模型
type RouteGroup struct {
	gorm.Model
	GatewayId  uint       //网关ID
	Name       string     //名称
	UriPrefix  string     //URI 前缀
	Predicates model.JSON //路由分组断言
	Filters    model.JSON //路由分组过滤器
	Auth       uint       //鉴权
	SortNumber int        //排序序号
}
