package weather

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers"
)

type OpenWeatherHandler struct{}

func (w *OpenWeatherHandler) FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error) {
	return 0, nil
}
