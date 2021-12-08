package route

import (
	"codetube.cn/gateway/interfaces"
	"codetube.cn/gateway/route/filters"
	"reflect"
	"sort"
	"strings"
)

type SimpleFilter string

func (this SimpleFilter) filter() interfaces.GatewayFilter {
	filterValueSplit := strings.Split(string(this), "=")
	filter, ok := filters.FilterMap.Load(filterValueSplit[0])
	if ok {
		applyConfig := ""
		if len(filterValueSplit) > 1 {
			applyConfig = filterValueSplit[1]
		}
		return filter.(interfaces.FilterFactory).Apply(applyConfig)
	}
	return nil
}

func (this SimpleFilter) getClass() interfaces.FilterFactory {
	filterValueSplit := strings.Split(string(this), "=")
	filter, ok := filters.FilterMap.Load(filterValueSplit[0])
	if ok {
		return filter.(interfaces.FilterFactory)
	}
	return nil
}

func (this *Route) FilterRequest(exchange *interfaces.ServerWebExchange) interfaces.ResponseFilters {
	//排序过滤器
	if len(this.orderedFilters) < 1 && len(this.Filters) > 0 {
		this.orderFilter()
	}
	responseFilters := make(interfaces.ResponseFilters, 0)
	for _, filter := range this.orderedFilters {
		responseFilter := filter.(SimpleFilter).filter()(exchange)
		if responseFilter != nil {
			responseFilters = append(responseFilters, responseFilter)
		}
	}
	return responseFilters
}

func (this *Route) orderFilter() {
	this.orderedFilters = make([]interface{}, 0)
	for _, f := range this.Filters {
		v := reflect.ValueOf(f)
		if v.Kind() == reflect.String {
			if obj := SimpleFilter(v.String()).getClass(); obj != nil {
				this.orderedFilters = append(this.orderedFilters, SimpleFilter(v.String()))
			}
		}
	}
	//排序
	sort.SliceStable(this.orderedFilters, func(i, j int) bool {
		return this.orderedFilters[i].(SimpleFilter).getClass().GetOrder() < this.orderedFilters[j].(SimpleFilter).getClass().GetOrder()
	})
}
