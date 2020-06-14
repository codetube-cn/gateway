package filters

import (
	"cn.codetube.gateway/interfaces"
	"strings"
)

func init() {
	RegisterFilter("StripPrefix", NewStripPrefixFilter())
}

type StripPrefixFilter struct{}

func (this *StripPrefixFilter) Apply() interfaces.GatewayFilter {
	return func(exchange *interfaces.ServerWebExchange) {
		path := exchange.Request.URL.Path
		pathList := strings.Split(path, "/")
		exchange.Request.URL.Path = strings.Join(pathList[2:], "/")
	}
}

func NewStripPrefixFilter() *StripPrefixFilter {
	return &StripPrefixFilter{}
}
