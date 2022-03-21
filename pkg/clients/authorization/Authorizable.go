package authorization

import "net/http"

type Authorizer interface {
	SetAuthorization() (Authorization, error)
}

type Authorization interface {
	Authorize(r *http.Request)
}
