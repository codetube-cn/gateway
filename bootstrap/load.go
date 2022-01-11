package bootstrap

import (
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/filters"
	"codetube.cn/gateway/models"
	"codetube.cn/gateway/predicates"
	"encoding/json"
	"log"
	"net/http"
)

type RouteGroup struct {
	Name           string
	UriPrefix      string
	Predicates     Predicates
	Filters        Filters
	PredicateCodes map[string]*Predicate //所有使用的断言 code，用于判断某个断言是否被路由分组使用
	FilterCodes    map[string]*Filter    //所有使用的过滤器 code，用于判断某个过滤器是否被路由分组使用
}

func NewRouteGroup() *RouteGroup {
	return &RouteGroup{}
}

//GetUsedPredicate 获取路由分组使用的指定断言
func (r *RouteGroup) GetUsedPredicate(code string) (*Predicate, bool) {
	p, ok := r.PredicateCodes[code]
	return p, ok
}

//GetUsedFilter 获取路由分组使用的指定过滤器
func (r *RouteGroup) GetUsedFilter(code string) (*Filter, bool) {
	f, ok := r.FilterCodes[code]
	return f, ok
}

//RouteGroupsMapping 路由分组 mapping
type RouteGroupsMapping struct {
	RouteGroups map[uint]*RouteGroup
}

//NewRouteGroupsMapping 创建路由分组 mapping
func NewRouteGroupsMapping() RouteGroupsMapping {
	return RouteGroupsMapping{
		RouteGroups: map[uint]*RouteGroup{},
	}
}

//append 增加一个路由分组 mapping
func (m *RouteGroupsMapping) append(id uint, g *RouteGroup) {
	m.RouteGroups[id] = g
}

//get 获取指定 ID 的路由分组
func (m *RouteGroupsMapping) get(id uint) (*RouteGroup, bool) {
	rg, ok := m.RouteGroups[id]
	return rg, ok
}

type FilterValue string

type Filter struct {
	Code   string
	Filter filters.FilterInterface
}

func NewFilter() *Filter {
	return &Filter{}
}

type Filters []*Filter

type Predicate struct {
	Code      string
	Predicate predicates.PredicateInterface
}

func NewPredicate() *Predicate {
	return &Predicate{}
}

type Predicates []*Predicate

type Route struct {
	ID             uint
	GroupID        uint
	Name           string
	RouteID        string
	Uri            string
	Predicates     Predicates
	Filters        Filters
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

//match 是否匹配断言
func (r *Route) match(request *http.Request) bool {
	for _, p := range r.Predicates {
		if !p.Predicate.Match(request) {
			return false
		}
	}

	return true
}

type GatewayRoutes struct {
	Routes []*Route
}

func NewGatewayRoutes() *GatewayRoutes {
	return &GatewayRoutes{
		Routes: make([]*Route, 0),
	}
}

//getMatchedRoute 获取匹配上的路由
func (gr *GatewayRoutes) getMatchedRoute(request *http.Request) *Route {
	for _, route := range gr.Routes {
		if route.match(request) {
			return route
		}
	}
	return nil
}

//loadGateway 载入网关
func loadGateway() {
	//从文件加载基础配置，网关名称及监听地址
	gatewayConfig = config.InitConfig()

	//获取网关
	gw, err := config.GetGateway(gatewayConfig.Gateway)
	if err != nil {
		log.Fatal(err)
	}

	//获取 routeGroupModel
	//将 routeGroupModel 按 ID 索引
	routeGroupsMapping := NewRouteGroupsMapping()
	routeGroups, _ := config.GetRouteGroups(gw.ID)
	for _, routeGroupModel := range routeGroups {
		routeGroup := NewRouteGroup()
		routeGroup.Name = routeGroupModel.Name
		routeGroup.UriPrefix = routeGroupModel.UriPrefix
		routeGroupFilters := make(Filters, 0)
		routeGroupPredicates := make(Predicates, 0)
		routeGroupPredicateCodes := map[string]*Predicate{}
		routeGroupFilterCodes := map[string]*Filter{}

		//加载路由分组的过滤器
		var gf []models.RouteFilter
		routeGroupModel.Filters.Unmarshal(&gf)
		for _, f := range gf {
			v, _ := json.Marshal(f.Value)
			filter := NewFilter()
			filter.Code = f.Filter
			filter.Filter = filters.SystemFilters[f.Filter](string(v))
			routeGroupFilters = append(routeGroupFilters, filter)
			routeGroupFilterCodes[filter.Code] = filter
		}

		routeGroup.Filters = routeGroupFilters
		routeGroup.FilterCodes = routeGroupFilterCodes

		//加载路由分组的断言
		var gp []models.RoutePredicate
		routeGroupModel.Predicates.Unmarshal(&gp)
		for _, p := range gp {
			v, _ := json.Marshal(p.Value)
			predicate := NewPredicate()
			predicate.Code = p.Predicate
			predicate.Predicate = predicates.SystemPredicates[p.Predicate](string(v))
			routeGroupPredicates = append(routeGroupPredicates, predicate)
			routeGroupPredicateCodes[predicate.Code] = predicate
		}

		routeGroup.Predicates = routeGroupPredicates
		routeGroup.PredicateCodes = routeGroupPredicateCodes
		routeGroupsMapping.append(routeGroupModel.ID, routeGroup)
	}

	//读取路由
	routes, _ := config.GetRoutes(gw.ID)

	for _, r := range routes {
		//路由处理
		//uri 要加上分组的 uri_prefix
		//predicate 和 filter 要覆盖分组的

		//路由
		gr := &Route{
			ID:             r.ID,
			GroupID:        r.GroupId,
			Name:           r.Name,
			RouteID:        r.RouteId,
			Uri:            r.Uri,
			Predicates:     nil,
			Filters:        nil,
			PredicateCodes: map[string]*Predicate{},
			FilterCodes:    map[string]*Filter{},
			SortNumber:     r.SortNumber,
		}

		routeFilters := make(Filters, 0)
		routePredicates := make(Predicates, 0)
		routePredicateCodes := map[string]*Predicate{}
		routeFilterCodes := map[string]*Filter{}

		//解析过滤器
		var rf []models.RouteFilter
		r.Filters.Unmarshal(&rf)
		for _, f := range rf {
			v, _ := json.Marshal(f.Value)
			filter := NewFilter()
			filter.Code = f.Filter
			filter.Filter = filters.SystemFilters[f.Filter](string(v))
			routeFilters = append(routeFilters, filter)
			routeFilterCodes[filter.Code] = filter
		}
		gr.Filters = routeFilters
		gr.FilterCodes = routeFilterCodes

		//解析断言
		var rp []models.RoutePredicate
		r.Predicates.Unmarshal(&rp)
		for _, p := range rp {
			v, _ := json.Marshal(p.Value)
			predicate := NewPredicate()
			predicate.Code = p.Predicate
			predicate.Predicate = predicates.SystemPredicates[p.Predicate](string(v))
			routePredicates = append(routePredicates, predicate)
			routePredicateCodes[predicate.Code] = predicate
		}
		gr.Predicates = routePredicates
		gr.PredicateCodes = routePredicateCodes

		//路由适配分组
		if r.GroupId > 0 {
			if rg, ok := routeGroupsMapping.get(r.GroupId); ok {
				//增加分组 Uri前缀
				if rg.UriPrefix != "" {
					gr.RouteID = rg.UriPrefix + gr.Uri
				}
				//增加分组断言
				rps := make([]*Predicate, 0)    //路由所有断言
				rpcs := map[string]*Predicate{} //路由所有断言 mapping
				for _, p := range rg.Predicates {
					//如果路由分组中某个断言在路由中重新声明了，则用路由中声明的断言覆盖
					if gp, ok := gr.GetUsedPredicate(p.Code); ok {
						rps = append(rps, gp)
						rpcs[p.Code] = gp
					} else {
						rps = append(rps, p)
						rpcs[p.Code] = p
					}
				}
				for _, p := range gr.Predicates {
					//路由中声明的断言，如果在路由分组中未被声明，才会被使用
					if _, ok := rg.GetUsedPredicate(p.Code); !ok {
						rps = append(rps, p)
						rpcs[p.Code] = p
					}
				}
				//增加分组过滤器
				rfs := make([]*Filter, 0) //路由所有过滤器
				rfcs := map[string]*Filter{}
				for _, f := range rg.Filters {
					//如果路由分组中某个断言在路由中重新声明了，则用路由中声明的断言覆盖
					if gf, ok := gr.GetUsedFilter(f.Code); ok {
						rfs = append(rfs, gf)
						rfcs[f.Code] = gf
					} else {
						rfs = append(rfs, f)
						rfcs[f.Code] = f
					}
				}
				for _, f := range gr.Filters {
					//路由中声明的断言，如果在路由分组中未被声明，才会被使用
					if _, ok := rg.GetUsedFilter(f.Code); !ok {
						rfs = append(rfs, f)
						rfcs[f.Code] = f
					}
				}

				gr.Predicates = rps
				gr.PredicateCodes = rpcs
				gr.Filters = rfs
				gr.FilterCodes = rfcs
			}
		}

		gatewayRoutes.Routes = append(gatewayRoutes.Routes, gr)
	}
}
