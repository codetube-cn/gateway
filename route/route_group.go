package route

// Group 路由分组
type Group struct {
	Name           string
	UriPrefix      string
	Predicates     Predicates
	Filters        Filters
	PredicateCodes map[string]*Predicate //所有使用的断言 code，用于判断某个断言是否被路由分组使用
	FilterCodes    map[string]*Filter    //所有使用的过滤器 code，用于判断某个过滤器是否被路由分组使用
}

// NewGroup 创建路由分组
func NewGroup() *Group {
	return &Group{}
}

//GetUsedPredicate 获取路由分组使用的指定断言
func (r *Group) GetUsedPredicate(code string) (*Predicate, bool) {
	p, ok := r.PredicateCodes[code]
	return p, ok
}

//GetUsedFilter 获取路由分组使用的指定过滤器
func (r *Group) GetUsedFilter(code string) (*Filter, bool) {
	f, ok := r.FilterCodes[code]
	return f, ok
}

//GroupsMapping 路由分组 mapping
type GroupsMapping struct {
	Groups map[uint]*Group
}

//NewGroupsMapping 创建路由分组 mapping
func NewGroupsMapping() *GroupsMapping {
	return &GroupsMapping{
		Groups: map[uint]*Group{},
	}
}

//Append 增加一个路由分组 mapping
func (m *GroupsMapping) Append(id uint, g *Group) {
	m.Groups[id] = g
}

//Get 获取指定 ID 的路由分组
func (m *GroupsMapping) Get(id uint) (*Group, bool) {
	rg, ok := m.Groups[id]
	return rg, ok
}

