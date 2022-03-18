package tracks

type Tracks []string

type Handler interface {
	FetchTracks(musicType MusicType) (Tracks, error)
}

type TrackStatus int

const (
	ErrorInvalidType TrackStatus = iota
)

type TrackErr struct {
	Status  TrackStatus
	Message string
}

func (s TrackErr) Error() string {
	return s.Message
}
