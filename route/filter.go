package route

import (
	"codetube.cn/gateway/interfaces"
	"codetube.cn/gateway/route/filters"
	"reflect"
	"sort"
	"strings"
)

type SimpleFilter string

func (f SimpleFilter) filter() interfaces.GatewayFilter {
	filterValueSplit := strings.Split(string(f), "=")
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

func (f SimpleFilter) getClass() interfaces.FilterFactory {
	filterValueSplit := strings.Split(string(f), "=")
	filter, ok := filters.FilterMap.Load(filterValueSplit[0])
	if ok {
		return filter.(interfaces.FilterFactory)
	}
	return nil
}

func (r *Route) FilterRequest(exchange *interfaces.ServerWebExchange) interfaces.ResponseFilters {
	//排序过滤器
	if len(r.orderedFilters) < 1 && len(r.Filters) > 0 {
		r.orderFilter()
	}
	responseFilters := make(interfaces.ResponseFilters, 0)
	for _, filter := range r.orderedFilters {
		responseFilter := filter.(SimpleFilter).filter()(exchange)
		if responseFilter != nil {
			responseFilters = append(responseFilters, responseFilter)
		}
	}
	return responseFilters
}

func (r *Route) orderFilter() {
	r.orderedFilters = make([]interface{}, 0)
	for _, f := range r.Filters {
		v := reflect.ValueOf(f)
		if v.Kind() == reflect.String {
			if obj := SimpleFilter(v.String()).getClass(); obj != nil {
				r.orderedFilters = append(r.orderedFilters, SimpleFilter(v.String()))
			}
		}
	}
	//排序
	sort.SliceStable(r.orderedFilters, func(i, j int) bool {
		return r.orderedFilters[i].(SimpleFilter).getClass().GetOrder() < r.orderedFilters[j].(SimpleFilter).getClass().GetOrder()
	})
}
