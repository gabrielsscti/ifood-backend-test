package tracks

type SpotifyTrackHandler struct{}

func (s *SpotifyTrackHandler) FetchTracks(musicType MusicType) (Tracks, error) {
	return []string{}, nil
}
