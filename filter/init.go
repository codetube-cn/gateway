package filter

//SystemFilters 系统支持的过滤器，只有这些过滤器才支持使用
var SystemFilters map[string]func(v string) Interface

func init() {
	SystemFilters = make(map[string]func(v string) Interface)
	initSystemFilter("cors", NewCORSFilter)
	initSystemFilter("set_header", NewSetHeaderFilter)
	initSystemFilter("strip_prefix", NewStripPrefixFilter)
}

//initSystemFilter 初始化支持的系统过滤器
func initSystemFilter(f string, fn func() Interface) {
	SystemFilters[f] = func(v string) Interface {
		filter := fn()
		filter.LoadValue(v)
		return filter
	}
}
