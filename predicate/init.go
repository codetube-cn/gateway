package predicate

//SystemPredicates 系统支持的断言，只有这些断言才支持使用
var SystemPredicates map[string]func(v string) Interface

func init() {
	SystemPredicates = make(map[string]func(v string) Interface)
	initSystemPredicate("path", NewPathPredicate)
	initSystemPredicate("header", NewHeaderPredicate)
	initSystemPredicate("method", NewMethodPredicate)
}

//initSystemPredicate 初始化支持的系统过滤器
func initSystemPredicate(p string, fn func() Interface) {
	SystemPredicates[p] = func(v string) Interface {
		predicate := fn()
		predicate.LoadValue(v)
		return predicate
	}
}