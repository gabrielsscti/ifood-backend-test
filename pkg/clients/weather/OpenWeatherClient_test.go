package weather

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients"
	"os"
	"testing"
)

var openWeatherClient WeatherClient

func TestMain(t *testing.M) {
	clients.TryLoadEnvironmentFile()
	openWeatherClient = NewOpenWeatherClient()
	os.Exit(t.Run())
}

func TestFetchTemperatureByCity(t *testing.T) {
	_, err := openWeatherClient.FetchTemperature(&CityLocation{"São Luís"})

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFetchTemperatureByCoordinate(t *testing.T) {
	_, err := openWeatherClient.FetchTemperature(&CoordinateLocation{-29.948600, 51.100500})

	if err != nil {
		t.Errorf(err.Error())
	}
}
