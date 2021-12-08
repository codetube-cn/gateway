package route

import (
	"codetube.cn/gateway/route/predicates"
)

type Predicates struct {
	Header predicates.HeaderPredicate
	Method predicates.MethodPredicate
	Host   string
	Path   predicates.PathPredicate
}

type Route struct {
	Id             string
	Url            string
	Predicates     Predicates
	Filters        []interface{}
	orderedFilters []interface{} //排序过后的过滤器
}

type Routes []*Route
