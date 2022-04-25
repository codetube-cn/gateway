package bootstrap

import (
	"codetube.cn/gateway/config"
	"codetube.cn/gateway/filter"
	"codetube.cn/gateway/gateway"
	"codetube.cn/gateway/vars"
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
			//鉴权
			if route.Auth == vars.JwtAuthMust || route.Auth == vars.JwtAuthShould {
				authorization := ""
				if v, ok := request.Header["Authorization"]; ok {
					if len(v) > 0 {
						authorization = v[0]
					}
				}
				userId := ""
				_, claims, err := gateway.ParseJwt(authorization)
				if err != nil && route.Auth == vars.JwtAuthMust {
					writer.WriteHeader(http.StatusUnauthorized)
					return
				}
				if err == nil && claims.ID != "" {
					userId = claims.ID
				}
				request.Header.Set("CodeTube-User-ID", userId)
			} else {
				request.Header.Del("CodeTube-User-ID")
			}

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
