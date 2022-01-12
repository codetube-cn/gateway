package route

import (
	"codetube.cn/gateway/filter"
)

// FilterValue 路由过滤器值
type FilterValue string

// Filter 路由过滤器
type Filter struct {
	Code   string
	Filter filter.Interface
}

// NewFilter 创建路由过滤器
func NewFilter() *Filter {
	return &Filter{}
}

// Filters 路由过滤器列表
type Filters []*Filter
