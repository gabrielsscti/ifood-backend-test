package weather

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/parameterizable"
)

type WeatherClient interface {
	FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error)
}
