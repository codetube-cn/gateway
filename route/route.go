package route

import "net/http"

// Route 路由
type Route struct {
	ID             uint
	GroupID        uint
	Name           string
	RouteID        string
	Uri            string
	Predicates     Predicates
	Filters        Filters
	Auth           uint
	PredicateCodes map[string]*Predicate //所有使用的断言 code，用于判断某个断言是否被路由使用
	FilterCodes    map[string]*Filter    //所有使用的过滤器 code，用于判断某个过滤器是否被路由使用
	SortNumber     uint
}

//GetUsedPredicate 获取路由使用的指定断言
func (r *Route) GetUsedPredicate(code string) (*Predicate, bool) {
	p, ok := r.PredicateCodes[code]
	return p, ok
}

//GetUsedFilter 获取路由使用的指定过滤器
func (r *Route) GetUsedFilter(code string) (*Filter, bool) {
	f, ok := r.FilterCodes[code]
	return f, ok
}

//Match 是否匹配断言
func (r *Route) Match(request *http.Request) bool {
	for _, p := range r.Predicates {
		if !p.Predicate.Match(request) {
			return false
		}
	}

	return true
}
