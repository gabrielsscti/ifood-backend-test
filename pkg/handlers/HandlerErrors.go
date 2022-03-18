package handlers

type FetchStatus int

const (
	ErrorAuthentication FetchStatus = iota
	ErrorFetching
	ErrorNotFound
)

type HandlerErr struct {
	Status  FetchStatus
	Message string
}

func (s HandlerErr) Error() string {
	return s.Message
}
