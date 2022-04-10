package weather

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"github.com/gabrielsscti/ifood-backend-test/pkg/parameterizable"
	"io/ioutil"
	"net/http"
)

type OpenWeatherClient struct {
	authorization.Authorization
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type weatherReturn struct {
	Main `json:"main"`
}

const openWeatherURL = "https://api.openweathermap.org/data/2.5/weather?"

func NewOpenWeatherClient(auth authorization.Authorization) (o *OpenWeatherClient) {
	o = &OpenWeatherClient{auth}
	return o
}

func (w *OpenWeatherClient) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	client := &http.Client{}
	urlRequest := openWeatherURL + parameterizable.GETParameter()

	req, err := http.NewRequest("GET", urlRequest, nil)
	if err != nil {
		return 0, fmt.Errorf("in FetchTemperature: %w", err)
	}
	w.Authorization.Authorize(req)

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("in FetchTemperature: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("in FetchTemperature: %w", err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var ret weatherReturn
	err = json.Unmarshal(byteValue, &ret)
	if err != nil {
		return 0, fmt.Errorf("in FetchTemperature: %w", err)
	}

	return KelvinToCelsius(ret.Temperature), nil
}
