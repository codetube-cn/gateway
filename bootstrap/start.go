package bootstrap

import (
	"codetube.cn/gateway/components"
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/interfaces"
	"codetube.cn/gateway/models"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func Start() {
	gatewayConfig := config.InitConfig()
	err := components.DB.MysqlInit()
	if err != nil {
		log.Fatal(err)
	}

	//@todo 调试 start
	gw, err := config.GetGateway(gatewayConfig.Gateway)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gw)
	gs,_ := config.GetRouteGroups(gw.ID)
	fmt.Println("gs:", gs)

	fs,_ := config.GetFilters(gw.ID)
	fmt.Println("fs:", fs)

	ps,_ := config.GetPredicates(gw.ID)
	fmt.Println("ps:", ps)

	for _,p := range ps {
		pe := &models.PredicateExtra{}
		fmt.Println(p.Extra.Unmarshal(pe), pe)
		for k,v := range pe.Options {
			fmt.Println(k,v.Value,v.Label)
		}
	}
	//@todo 调试 end

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if route := gatewayConfig.Routes.Match(request); route != nil {
			remote, _ := url.Parse(route.Url)
			//fmt.Println(remote)
			exchange := interfaces.BuildServerWebExchange(request)
			responseFilters := route.FilterRequest(exchange)
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ModifyResponse = func(response *http.Response) error {
				responseFilters.Filter(response)
				return nil
			}
			proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
				fmt.Println(err)
			}
			proxy.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(gatewayConfig.Listen.Host+":"+strconv.Itoa(gatewayConfig.Listen.Port), nil)
}
