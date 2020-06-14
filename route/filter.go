package route

import (
	"cn.codetube.gateway/interfaces"
	"cn.codetube.gateway/route/filters"
	"reflect"
	"strings"
)

type SimpleFilter string

func (this SimpleFilter) filter() interfaces.GatewayFilter {
	filterValueSplit := strings.Split(string(this), "=")
	if len(filterValueSplit) != 2 {
		return nil
	}
	filter, ok := filters.FilterMap.Load(filterValueSplit[0])
	if ok {
		return filter.(interfaces.FilterFactory).Apply()
	}
	return nil
}

func (this *Route) FilterBefore(exchange *interfaces.ServerWebExchange) {
	for _, filter := range this.Filters {
		v := reflect.ValueOf(filter)
		if v.Kind() == reflect.String {
			gatewayFilter := SimpleFilter(v.String()).filter()
			if gatewayFilter != nil {
				gatewayFilter(exchange)
			}
		}
	}
}
