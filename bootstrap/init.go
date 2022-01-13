package bootstrap

import (
	"codetube.cn/gateway/gateway"
	"codetube.cn/gateway/route"
)

// 网关路由
var gatewayRoutes = gateway.NewRoutes()

// 路由分组 mapping
var routeGroupsMapping = route.NewGroupsMapping()
