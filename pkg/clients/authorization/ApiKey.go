package authorization

import "net/http"

type ApiKey struct {
	ApiKey string
}

func (a ApiKey) Authorize(r *http.Request) {
	v := r.URL.Query()
	v.Add("appid", a.ApiKey)
	r.URL.RawQuery = v.Encode()
}
