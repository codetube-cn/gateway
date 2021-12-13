package predicates

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type HeaderPredicate string

func (p HeaderPredicate) Match(request *http.Request) bool {
	s := string(p)
	headers := strings.Split(s, ",")
	headerCount := len(headers)
	if headerCount < 2 || headerCount%2 != 0 {
		return true
	}
	for i := 0; i < headerCount; i += 2 {
		key := headers[i]
		pattern := headers[i+1]
		//未取到指定头信息
		if value, ok := request.Header[key]; !ok {
			return false
		} else {
			reg, err := regexp.Compile(pattern)
			if err != nil {
				log.Println(err)
				return false
			}
			if !reg.MatchString(value[0]) {
				return false
			}
		}
	}
	return true
}
