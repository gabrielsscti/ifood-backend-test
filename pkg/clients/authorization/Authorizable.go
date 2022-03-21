package authorization

import "net/http"

type Authorizable interface {
	GetAuthorization() (Authorization, error)
}

type Authorizer interface {
	SetAuthorization() (Authorization, error)
}

type Authorization interface {
	Authorize(r *http.Request)
}
