package track_manager

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers"
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers/tracks"
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers/weather"
	"reflect"
	"testing"
)

var trackManager TrackManager

type TableTest struct {
	name           string
	trackHandler   tracks.Handler
	weatherHandler weather.Handler
	expected       tracks.Tracks
	errMsg         string
}

func testTableTest(data []TableTest, t *testing.T) {
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			trackManager := TrackManager{
				TracksHandler:  d.trackHandler,
				WeatherHandler: d.weatherHandler,
			}
			result, err := trackManager.GetPlaylist(ParameterizableMock{})
			if !reflect.DeepEqual(result, []string(d.expected)) {
				t.Errorf("Expected `%s`, got `%s`", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

func TestGetPlaylist(t *testing.T) {
	trackHandler := ValidTrackHandlerStub{}
	invalidTrackHandler := InvalidTrackHandlerStub{}
	validWeather45 := ValidWeatherStubByTemperature45{}
	validWeather27 := ValidWeatherStubByTemperature27{}
	validWeather13 := ValidWeatherStubByTemperature13{}
	validWeather9 := ValidWeatherStubByTemperature9{}

	data := []TableTest{
		{"party_playlist", &trackHandler, &validWeather45, tracks.Tracks{"party1", "party2", "party3"}, ""},
		{"pop_playlist", &trackHandler, &validWeather27, tracks.Tracks{"pop1", "pop2", "pop3"}, ""},
		{"rock_playlist", &trackHandler, &validWeather13, tracks.Tracks{"rock1", "rock2", "rock3"}, ""},
		{"classical_playlist", &trackHandler, &validWeather9, tracks.Tracks{"classical1", "classical2", "classical3"}, ""},
		{"invalid_playlist", &invalidTrackHandler, &validWeather9, nil, "invalid kind of music type"},
	}

	testTableTest(data, t)
}

func TestGetPlaylistError(t *testing.T) {
	validTrackHandler := ValidTrackHandlerStub{}
	invalidTrackHandlerAuth := AuthErrorTrackHandlerStub{}
	validWeatherHandler := ValidWeatherStubByTemperature45{}
	invalidWeatherHandlerAuth := AuthErrorWeatherStub{}

	data := []TableTest{
		{"valid_trackHandler_invalid_weatherHandler", &validTrackHandler,
			&invalidWeatherHandlerAuth, nil, "invalid authentication"},
		{"invalid_trackHandler_valid_weatherHandler", &invalidTrackHandlerAuth,
			&validWeatherHandler, nil, "invalid authentication"},
		{"both_invalid", &invalidTrackHandlerAuth, &invalidWeatherHandlerAuth,
			nil, "invalid authentication"},
	}

	testTableTest(data, t)
}

type ParameterizableMock struct {
}

func (p ParameterizableMock) GETParameter() string {
	return ""
}

type ValidTrackHandlerStub struct{}

func (h *ValidTrackHandlerStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
	switch musicType {
	case tracks.Pop:
		return tracks.Tracks{"pop1", "pop2", "pop3"}, nil
	case tracks.Rock:
		return tracks.Tracks{"rock1", "rock2", "rock3"}, nil
	case tracks.Classical:
		return tracks.Tracks{"classical1", "classical2", "classical3"}, nil
	case tracks.Party:
		return tracks.Tracks{"party1", "party2", "party3"}, nil
	default:
		return nil, tracks.TrackErr{
			Status:  tracks.ErrorInvalidType,
			Message: "invalid kind of music type",
		}
	}
}

type InvalidTrackHandlerStub struct{}

func (h *InvalidTrackHandlerStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
	return nil, tracks.TrackErr{
		Status:  tracks.ErrorInvalidType,
		Message: "invalid kind of music type"}
}

type AuthErrorTrackHandlerStub struct{}

func (h *AuthErrorTrackHandlerStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
	return nil, handlers.HandlerErr{
		Status:  handlers.ErrorAuthentication,
		Message: "invalid authentication",
	}
}

type ValidWeatherStubByTemperature45 struct{}

func (v *ValidWeatherStubByTemperature45) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 45.3, nil
}

type ValidWeatherStubByTemperature27 struct{}

func (v *ValidWeatherStubByTemperature27) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 27.99, nil
}

type ValidWeatherStubByTemperature13 struct{}

func (v *ValidWeatherStubByTemperature13) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 13.21, nil
}

type ValidWeatherStubByTemperature9 struct{}

func (v *ValidWeatherStubByTemperature9) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 9.19, nil
}

type AuthErrorWeatherStub struct{}

func (v *AuthErrorWeatherStub) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 0, handlers.HandlerErr{
		Status:  handlers.ErrorAuthentication,
		Message: "invalid authentication",
	}
}
