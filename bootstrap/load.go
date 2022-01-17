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
	"time"
)

func loadGateway() {
	//从数据库网关
	gwm, err := gateway.GetGateway(config.GatewayConfig.Gateway)
	if err != nil {
		log.Fatal(err)
	}

	//网关监听信息
	if gwm.Host != "" {
		config.GatewayConfig.Listen.Host = gwm.Host
	} else if config.GatewayConfig.Listen.Host == "" {
		config.GatewayConfig.Listen.Host = "localhost"
	}
	if gwm.Port > 0 {
		config.GatewayConfig.Listen.Port = gwm.Port
	} else if config.GatewayConfig.Listen.Port < 1000 {
		config.GatewayConfig.Listen.Port = 8088
	}
	gw = gwm
}

// 获取网关所有路由（路由分组 mapping 及各路由）
func getGatewayRoutes() (*route.GroupsMapping, *gateway.Routes, error) {
	//载入路由分组 mapping
	groupsMapping, err := getRouteGroupMapping(gw.ID)
	if err != nil {
		return nil, nil, err
	}

	//读取路由
	routes, _ := gateway.GetRoutes(gw.ID)
	gwRoutes := gateway.NewRoutes()
	//转换路由并压入网关路由列表
	for _, r := range routes {
		gr, err := transRoute(r, groupsMapping)
		if err != nil {
			return nil, nil, err
		}
		gwRoutes.Routes = append(gwRoutes.Routes, gr)
	}

	return groupsMapping, gwRoutes, nil
}

//获取路由分组 mapping
func getRouteGroupMapping(gatewayId uint) (*route.GroupsMapping, error) {
	//获取 routeGroupModel
	//将 routeGroupModel 按 ID 索引
	routeGroups, _ := gateway.GetRouteGroups(gatewayId)
	groupsMapping := route.NewGroupsMapping()
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
		err := routeGroupModel.Filters.Unmarshal(&gf)
		if err != nil {
			return groupsMapping, err
		}
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
		err = routeGroupModel.Predicates.Unmarshal(&gp)
		if err != nil {
			return groupsMapping, err
		}
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
		groupsMapping.Append(routeGroupModel.ID, routeGroup)
	}

	return groupsMapping, nil
}

//将数据库中路由转换成路由对象
func transRoute(r *models.Route, gm *route.GroupsMapping) (*route.Route, error) {
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
	var mrfs []models.RouteFilter
	err := r.Filters.Unmarshal(&mrfs)
	if err != nil {
		return nil, err
	}
	for _, f := range mrfs {
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
	var mrps []models.RoutePredicate
	err = r.Predicates.Unmarshal(&mrps)
	if err != nil {
		return nil, err
	}
	for _, p := range mrps {
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
		if rg, ok := gm.Get(r.GroupId); ok {
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
				if _, ok = rg.GetUsedFilter(f.Code); !ok {
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

	return gr, nil
}

// 定时重载网关配置
func watchGateway() *time.Ticker {
	interval := 5
	if config.GatewayConfig.WatchIntervalSeconds > 0 {
		interval = config.GatewayConfig.WatchIntervalSeconds
	}
	log.Println("重新加载网关配置定时器时间间隔（秒）：", interval)
	t := time.NewTicker(time.Duration(interval) * time.Second)
	go func() {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				//@todo 这里最好加个缓存判断，毕竟网关配置不会经常动，每次查数据库过于浪费
				//在启动时生成一个 uuid 来标识当前进程，值为最近刷新时间（过期时间 2 倍于配置的间隔时间，如果没有，则设置为当前时间，且当前不处理）
				//在网关相关变更时，要求有一个网关 code 的缓存，值为最后变更时间（如无，则当作无变更，不处理）
				//当网关最后变更时间 > 最近刷新时间时，才重新载入配置
				go func() {
					if err := recover(); err != nil {
						log.Printf("recover: %v", err)
					}
					rGm, rRs, rErr := getGatewayRoutes()
					if rErr != nil {
						log.Printf("error: %v\n", rErr)
					} else {
						routeGroupsMapping = rGm
						gatewayRoutes = rRs
					}
				}()
			}
		}
	}()
	return t
}
