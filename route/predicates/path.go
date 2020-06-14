package predicates

import (
	"net/http"
	"path/filepath"
)

type PathPredicate string

func (this PathPredicate) Match(request *http.Request) bool {
	matched, err := filepath.Match(string(this), request.URL.Path)
	if err != nil || !matched {
		return false
	}
	return true
}
