package bootstrap

import (
	"codetube.cn/gateway/filters"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func serve() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if route := gatewayRoutes.getMatchedRoute(request); route != nil {
			remote, _ := url.Parse(route.Uri)
			exchange := filters.BuildServerWebExchange(request)
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

	fmt.Println("Gateway Serve: " + gatewayConfig.Listen.Host+":"+strconv.Itoa(gatewayConfig.Listen.Port))
	http.ListenAndServe(gatewayConfig.Listen.Host+":"+strconv.Itoa(gatewayConfig.Listen.Port), nil)
}
