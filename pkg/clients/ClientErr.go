package clients

type FetchStatus int

const (
	ErrorAuthentication FetchStatus = iota
	ErrorFetching
	ErrorNotFound
)

type ClientErr struct {
	Status  FetchStatus
	Message string
}

func (s ClientErr) Error() string {
	return s.Message
}
