package config

import (
	"codetube.cn/gateway/components"
	"codetube.cn/gateway/models"
	"fmt"
)

//GetGateway 从数据表中获取指定的网关记录
func GetGateway(code string) (gateway *models.Gateway, err error) {
	gw := &models.Gateway{}
	components.GatewayDB.Where("code = ?", code).First(gw)
	fmt.Println(gw)
	if gw.ID < 1 {
		return nil, fmt.Errorf("can not found gateway[%s]", code)
	}
	return gw, err
}

//GetRouteGroups 获取路由分组
func GetRouteGroups(gatewayId uint) (groups []*models.RouteGroup, err error) {
	components.GatewayDB.Model(&models.RouteGroup{}).Where("gateway_id = ?", gatewayId).Scan(&groups)
	return
}

//GetFilters 获取过滤器
func GetFilters(gatewayId uint) (filters []*models.Filter, err error) {
	components.GatewayDB.Model(&models.Filter{}).Where("gateway_id in ?", []uint{gatewayId, 0}).Scan(&filters)
	return
}

//GetPredicates 获取断言
func GetPredicates(gatewayId uint) (predicates []*models.Predicate, err error) {
	components.GatewayDB.Model(&models.Predicate{}).Where("gateway_id in ?", []uint{gatewayId, 0}).Scan(&predicates)
	return
}
