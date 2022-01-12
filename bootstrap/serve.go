package bootstrap

import (
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/filter"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

//serve 启动监听服务
func serve() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if route := gatewayRoutes.GetMatchedRoute(request); route != nil {
			remote, _ := url.Parse(route.Uri)
			exchange := filter.BuildServerWebExchange(request)
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ModifyResponse = func(response *http.Response) error {
				for _, f := range route.Filters {
					f.Filter.Apply(exchange, response)
				}
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

	fmt.Println("Gateway Serve: " + config.GatewayConfig.Listen.Host+":"+strconv.Itoa(config.GatewayConfig.Listen.Port))
	http.ListenAndServe(config.GatewayConfig.Listen.Host+":"+strconv.Itoa(config.GatewayConfig.Listen.Port), nil)
}
