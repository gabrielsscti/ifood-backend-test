package weather

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"github.com/gabrielsscti/ifood-backend-test/pkg/server/models"
	"os"
	"testing"
)

var openWeatherClient WeatherClient

func TestMain(t *testing.M) {
	clients.TryLoadEnvironmentFile()
	openWeatherClient = NewOpenWeatherClient(authorization.ApiKey{ApiKey: os.Getenv("OPEN_WEATHER_API_KEY")})
	os.Exit(t.Run())
}

func TestFetchTemperatureByCity(t *testing.T) {
	_, err := openWeatherClient.FetchTemperature(&models.CityLocation{"São Luís"})

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFetchTemperatureByCoordinate(t *testing.T) {
	_, err := openWeatherClient.FetchTemperature(&models.CoordinateLocation{-29.948600, 51.100500})

	if err != nil {
		t.Errorf(err.Error())
	}
}
