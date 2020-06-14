package route

import (
	"cn.codetube.gateway/interfaces"
	"net/http"
	"reflect"
	"strings"
)

func (this Routes) Match(request *http.Request) *Route {
	for _, route := range this {
		if this.isMatch(route, request) {
			return route
		}
	}
	return nil
}

func (this Routes) isMatch(route *Route, request *http.Request) bool {
	v := reflect.ValueOf(route.Predicates)
	for i := 0; i < v.NumField(); i++ {
		if matcher, ok := v.Field(i).Interface().(interfaces.PredicateMatcher); ok && strings.Trim(v.Field(i).String(), "") != "" {
			if !matcher.Match(request) {
				return false
			}
		}
	}
	return true
}
