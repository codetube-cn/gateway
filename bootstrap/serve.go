package bootstrap

import (
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/filter"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

//http服务
func httpServer() *http.Server {
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
				log.Println(err)
			}
			proxy.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})

	return &http.Server{
		Addr:    config.GatewayConfig.Listen.Host + ":" + strconv.Itoa(config.GatewayConfig.Listen.Port),
		Handler: http.DefaultServeMux,
	}
}
