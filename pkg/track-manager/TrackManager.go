package track_manager

import (
	"fmt"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/tracks"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/weather"
	"github.com/gabrielsscti/ifood-backend-test/pkg/parameterizable"
)

type TrackManager struct {
	TracksClient  tracks.TrackClient
	WeatherClient weather.WeatherClient
}

type TrackResponse struct {
	MusicType string   `json:"MusicType"`
	Tracks    []string `json:"Tracks"`
}

func CreateTrackManager(trackClient tracks.TrackClient, weatherClient weather.WeatherClient) *TrackManager {
	return &TrackManager{trackClient, weatherClient}
}

func (t *TrackManager) GetPlaylist(location parameterizable.GETParameterizable) (*TrackResponse, error) {
	temperature, err := t.WeatherClient.FetchTemperature(location)
	if err != nil {
		return nil, fmt.Errorf("in GetPlaylist: %w", err)
	}

	musicType := temperatureToMusicType(temperature)
	reqTracks, err := t.TracksClient.FetchTracks(musicType)

	if err != nil {
		return nil, fmt.Errorf("in GetPlaylist: %w", err)
	}
	return &TrackResponse{MusicType: musicType.String(), Tracks: reqTracks}, err
}

func temperatureToMusicType(temperature float64) tracks.MusicType {
	const (
		minParty = 30
		minPop   = 15
		minRock  = 10
	)

	if temperature > minParty {
		return tracks.Party
	} else if temperature > minPop {
		return tracks.Pop
	} else if temperature > minRock {
		return tracks.Rock
	} else {
		return tracks.Classical
	}
}
