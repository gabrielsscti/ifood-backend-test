package weather

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/handlers"
)

type Handler interface {
	FetchTemperature(parameterizable handlers.GETParameterizable) (float64, error)
}
