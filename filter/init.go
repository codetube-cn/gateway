package filter

//SystemFilters 系统支持的过滤器，只有这些过滤器才支持使用
var SystemFilters map[string]func(v string) FilterInterface

func init() {
	SystemFilters = make(map[string]func(v string) FilterInterface)
	initSystemFilter("cors", NewCORSFilter)
	initSystemFilter("set_header", NewSetHeaderFilter)
	initSystemFilter("strip_prefix", NewStripPrefixFilter)
}

//initSystemFilter 初始化支持的系统过滤器
func initSystemFilter(f string, fn func() FilterInterface) {
	SystemFilters[f] = func(v string) FilterInterface {
		filter := fn()
		filter.LoadValue(v)
		return filter
	}
}
