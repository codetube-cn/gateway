package bootstrap

import (
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/filter"
	"codetube.cn/gateway/gateway"
	"codetube.cn/gateway/models"
	"codetube.cn/gateway/predicate"
	"codetube.cn/gateway/route"
	"encoding/json"
	"log"
)



//loadGateway 载入网关
func loadGateway() {
	//获取网关
	gw, err := gateway.GetGateway(config.GatewayConfig.Gateway)
	if err != nil {
		log.Fatal(err)
	}

	//获取 routeGroupModel
	//将 routeGroupModel 按 ID 索引
	routeGroupsMapping := route.NewGroupsMapping()
	routeGroups, _ := gateway.GetRouteGroups(gw.ID)
	for _, routeGroupModel := range routeGroups {
		routeGroup := route.NewGroup()
		routeGroup.Name = routeGroupModel.Name
		routeGroup.UriPrefix = routeGroupModel.UriPrefix
		routeGroupFilters := make(route.Filters, 0)
		routeGroupPredicates := make(route.Predicates, 0)
		routeGroupPredicateCodes := map[string]*route.Predicate{}
		routeGroupFilterCodes := map[string]*route.Filter{}

		//加载路由分组的过滤器
		var gf []models.RouteFilter
		routeGroupModel.Filters.Unmarshal(&gf)
		for _, f := range gf {
			v, _ := json.Marshal(f.Value)
			rf := route.NewFilter()
			rf.Code = f.Filter
			rf.Filter = filter.SystemFilters[f.Filter](string(v))
			routeGroupFilters = append(routeGroupFilters, rf)
			routeGroupFilterCodes[rf.Code] = rf
		}

		routeGroup.Filters = routeGroupFilters
		routeGroup.FilterCodes = routeGroupFilterCodes

		//加载路由分组的断言
		var gp []models.RoutePredicate
		routeGroupModel.Predicates.Unmarshal(&gp)
		for _, p := range gp {
			v, _ := json.Marshal(p.Value)
			rp := route.NewPredicate()
			rp.Code = p.Predicate
			rp.Predicate = predicate.SystemPredicates[p.Predicate](string(v))
			routeGroupPredicates = append(routeGroupPredicates, rp)
			routeGroupPredicateCodes[rp.Code] = rp
		}

		routeGroup.Predicates = routeGroupPredicates
		routeGroup.PredicateCodes = routeGroupPredicateCodes
		routeGroupsMapping.Append(routeGroupModel.ID, routeGroup)
	}

	//读取路由
	routes, _ := gateway.GetRoutes(gw.ID)

	for _, r := range routes {
		//路由处理
		//uri 要加上分组的 uri_prefix
		//predicate 和 filter 要覆盖分组的

		//路由
		gr := &route.Route{
			ID:             r.ID,
			GroupID:        r.GroupId,
			Name:           r.Name,
			RouteID:        r.RouteId,
			Uri:            r.Uri,
			Predicates:     nil,
			Filters:        nil,
			PredicateCodes: map[string]*route.Predicate{},
			FilterCodes:    map[string]*route.Filter{},
			SortNumber:     r.SortNumber,
		}

		routeFilters := make(route.Filters, 0)
		routePredicates := make(route.Predicates, 0)
		routePredicateCodes := map[string]*route.Predicate{}
		routeFilterCodes := map[string]*route.Filter{}

		//解析过滤器
		var rf []models.RouteFilter
		r.Filters.Unmarshal(&rf)
		for _, f := range rf {
			v, _ := json.Marshal(f.Value)
			rf := route.NewFilter()
			rf.Code = f.Filter
			rf.Filter = filter.SystemFilters[f.Filter](string(v))
			routeFilters = append(routeFilters, rf)
			routeFilterCodes[rf.Code] = rf
		}
		gr.Filters = routeFilters
		gr.FilterCodes = routeFilterCodes

		//解析断言
		var rp []models.RoutePredicate
		r.Predicates.Unmarshal(&rp)
		for _, p := range rp {
			v, _ := json.Marshal(p.Value)
			rp := route.NewPredicate()
			rp.Code = p.Predicate
			rp.Predicate = predicate.SystemPredicates[p.Predicate](string(v))
			routePredicates = append(routePredicates, rp)
			routePredicateCodes[rp.Code] = rp
		}
		gr.Predicates = routePredicates
		gr.PredicateCodes = routePredicateCodes

		//路由适配分组
		if r.GroupId > 0 {
			if rg, ok := routeGroupsMapping.Get(r.GroupId); ok {
				//增加分组 Uri前缀
				if rg.UriPrefix != "" {
					gr.RouteID = rg.UriPrefix + gr.Uri
				}
				//增加分组断言
				rps := make([]*route.Predicate, 0)    //路由所有断言
				rpcs := map[string]*route.Predicate{} //路由所有断言 mapping
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
				rfs := make([]*route.Filter, 0) //路由所有过滤器
				rfcs := map[string]*route.Filter{}
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
