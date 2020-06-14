package predicates

import (
	"net/http"
	"strings"
)

type MethodPredicate string

func (this MethodPredicate) Match(request *http.Request) bool {
	s := string(this)
	methods := strings.Split(s, ",")
	if len(methods) == 0 {
		return true
	}
	for _, method := range methods {
		if strings.ToLower(method) == strings.ToLower(request.Method) {
			return true
		}
	}

	return false
}
