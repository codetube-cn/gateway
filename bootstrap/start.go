package bootstrap

import (
	"cn.codetube.gateway/config"
	"cn.codetube.gateway/interfaces"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func Start() {
	config := config.InitConfig()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if route := config.Routes.Match(request); route != nil {
			remote, _ := url.Parse(route.Url)
			exchange := interfaces.BuildServerWebExchange(request)
			responseFilters := route.FilterRequest(exchange)
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ModifyResponse = func(response *http.Response) error {
				responseFilters.Filter(response)
				return nil
			}
			proxy.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(config.Listen.Host+":"+strconv.Itoa(config.Listen.Port), nil)
}
