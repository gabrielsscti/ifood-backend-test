package track_manager

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/tracks"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/weather"
	"github.com/gabrielsscti/ifood-backend-test/pkg/parameterizable"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

type TableTest struct {
	name          string
	trackClient   tracks.TrackClient
	weatherClient weather.WeatherClient
	expected      tracks.Tracks
	errMsg        string
}

func testTableTest(data []TableTest, t *testing.T) {
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			trackManager := TrackManager{
				TracksClient:  d.trackClient,
				WeatherClient: d.weatherClient,
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

func TestGetPlaylistIntegrationSpotify(t *testing.T) {
	const spotifyTokenURL = "https://accounts.spotify.com/api/token"
	clients.TryLoadEnvironmentFile()
	trackClient := tracks.NewSpotifyTrackClient(authorization.NewClientCredentials(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"), spotifyTokenURL))
	weatherClient := weather.NewOpenWeatherClient(authorization.ApiKey{ApiKey: os.Getenv("OPEN_WEATHER_API_KEY")})
	trackManager := TrackManager{
		TracksClient:  trackClient,
		WeatherClient: weatherClient,
	}
	_tracks, err := trackManager.GetPlaylist(&weather.CityLocation{CityName: "São Luís"})
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.NotNil(t, _tracks)
	assert.Greater(t, len(_tracks), 0)
}

func TestGetPlaylist(t *testing.T) {
	trackClient := ValidTrackClientStub{}
	invalidTrackClient := InvalidTrackClientStub{}
	validWeather45 := ValidWeatherStubByTemperature45{}
	validWeather27 := ValidWeatherStubByTemperature27{}
	validWeather13 := ValidWeatherStubByTemperature13{}
	validWeather9 := ValidWeatherStubByTemperature9{}

	data := []TableTest{
		{"party_playlist", &trackClient, &validWeather45, tracks.Tracks{"party1", "party2", "party3"}, ""},
		{"pop_playlist", &trackClient, &validWeather27, tracks.Tracks{"pop1", "pop2", "pop3"}, ""},
		{"rock_playlist", &trackClient, &validWeather13, tracks.Tracks{"rock1", "rock2", "rock3"}, ""},
		{"classical_playlist", &trackClient, &validWeather9, tracks.Tracks{"classical1", "classical2", "classical3"}, ""},
		{"invalid_playlist", &invalidTrackClient, &validWeather9, nil, "in GetPlaylist: in GetPlaylist: invalid kind of music type"},
	}

	testTableTest(data, t)
}

func TestGetPlaylistError(t *testing.T) {
	validTrackClient := ValidTrackClientStub{}
	invalidTrackClientAuth := AuthErrorTrackClientStub{}
	validWeatherClient := ValidWeatherStubByTemperature45{}
	invalidWeatherClientAuth := AuthErrorWeatherStub{}

	data := []TableTest{
		{"valid_trackClient_invalid_weatherClient", &validTrackClient,
			&invalidWeatherClientAuth, nil, "in GetPlaylist: in GetPlaylist: invalid authentication"},
		{"invalid_trackClient_valid_weatherClient", &invalidTrackClientAuth,
			&validWeatherClient, nil, "in GetPlaylist: in GetPlaylist: invalid authentication"},
		{"both_invalid", &invalidTrackClientAuth, &invalidWeatherClientAuth,
			nil, "in GetPlaylist: in GetPlaylist: invalid authentication"},
	}

	testTableTest(data, t)
}

type ParameterizableMock struct {
}

func (p ParameterizableMock) GETParameter() string {
	return ""
}

type ValidTrackClientStub struct{}

func (h *ValidTrackClientStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
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
			Message: "in GetPlaylist: invalid kind of music type",
		}
	}
}

type InvalidTrackClientStub struct{}

func (h *InvalidTrackClientStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
	return nil, tracks.TrackErr{
		Status:  tracks.ErrorInvalidType,
		Message: "in GetPlaylist: invalid kind of music type"}
}

type AuthErrorTrackClientStub struct{}

func (h *AuthErrorTrackClientStub) FetchTracks(musicType tracks.MusicType) (tracks.Tracks, error) {
	return nil, clients.ClientErr{
		Status:  clients.ErrorAuthentication,
		Message: "in GetPlaylist: invalid authentication",
	}
}

type ValidWeatherStubByTemperature45 struct{}

func (v *ValidWeatherStubByTemperature45) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	return 45.3, nil
}

type ValidWeatherStubByTemperature27 struct{}

func (v *ValidWeatherStubByTemperature27) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	return 27.99, nil
}

type ValidWeatherStubByTemperature13 struct{}

func (v *ValidWeatherStubByTemperature13) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	return 13.21, nil
}

type ValidWeatherStubByTemperature9 struct{}

func (v *ValidWeatherStubByTemperature9) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	return 9.19, nil
}

type AuthErrorWeatherStub struct{}

func (v *AuthErrorWeatherStub) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	return 0, clients.ClientErr{
		Status:  clients.ErrorAuthentication,
		Message: "in GetPlaylist: invalid authentication",
	}
}
