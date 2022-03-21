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

func (t *TrackManager) GetPlaylist(location parameterizable.GETParameterizable) ([]string, error) {
	temperature, err := t.WeatherClient.FetchTemperature(location)
	if err != nil {
		return nil, fmt.Errorf("in GetPlaylist: %w", err)
	}

	musicType := temperatureToMusicType(temperature)
	reqTracks, err := t.TracksClient.FetchTracks(musicType)

	if err != nil {
		return nil, fmt.Errorf("in GetPlaylist: %w", err)
	}
	return reqTracks, err
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
