package track_manager

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers"
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers/tracks"
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers/weather"
)

type TrackManager struct {
	TracksHandler  tracks.Handler
	WeatherHandler weather.Handler
}

func (t *TrackManager) GetPlaylist(location handlers.GETParameterizable) ([]string, error) {
	temperature, err := t.WeatherHandler.FetchTemperature(location)
	if err != nil {
		return nil, err
	}

	musicType := temperatureToMusicType(temperature)
	reqTracks, err := t.TracksHandler.FetchTracks(musicType)

	if err != nil {
		return nil, err
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
