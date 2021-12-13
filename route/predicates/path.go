package predicates

import (
	"net/http"
	"path/filepath"
)

type PathPredicate string

func (p PathPredicate) Match(request *http.Request) bool {
	matched, err := filepath.Match(string(p), request.URL.Path)
	if err != nil || !matched {
		return false
	}
	return true
}
