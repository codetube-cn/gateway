package gateway

import (
	"codetube.cn/gateway/route"
	"net/http"
)

// Routes 网关所有路由
type Routes struct {
	Routes []*route.Route
}

// NewRoutes 创建网关路由
func NewRoutes() *Routes {
	return &Routes{
		Routes: make([]*route.Route, 0),
	}
}

//GetMatchedRoute 获取匹配上的路由
func (gr *Routes) GetMatchedRoute(request *http.Request) *route.Route {
	for _, route := range gr.Routes {
		if route.Match(request) {
			return route
		}
	}
	return nil
}
