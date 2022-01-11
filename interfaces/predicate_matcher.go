package interfaces

import "net/http"

type PredicateMatcher interface {
	Match(request *http.Request) bool
}

