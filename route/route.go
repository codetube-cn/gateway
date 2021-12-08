package route

import (
	"codetube.cn/gateway/route/predicates"
)

type Predicates struct {
	Header predicates.HeaderPredicate `yaml:"header"`
	Method predicates.MethodPredicate `yaml:"method"`
	Host   string                     `yaml:"host"`
	Path   predicates.PathPredicate   `yaml:"path"`
}

type Route struct {
	Id             string        `yaml:"id"`
	Url            string        `yaml:"url"`
	Predicates     Predicates    `yaml:"predicates"`
	Filters        []interface{} `yaml:"filters"`
	orderedFilters []interface{} //排序过后的过滤器
}

type Routes []*Route
