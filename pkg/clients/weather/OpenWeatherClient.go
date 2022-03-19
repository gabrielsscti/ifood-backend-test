package weather

import (
	"encoding/json"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients"
	"github.com/gabrielsscti/ifood-backend-test/pkg/parameterizable"
	"io/ioutil"
	"net/http"
	"os"
)

type OpenWeatherClient struct {
	apiKey string
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type weatherReturn struct {
	Main `json:"main"`
}

const OpenWeatherURL = "https://api.openweathermap.org/data/2.5/weather?"

func NewOpenWeatherClient() (o *OpenWeatherClient) {
	o = &OpenWeatherClient{os.Getenv("OPEN_WEATHER_API_KEY")}
	return o
}

func (w *OpenWeatherClient) FetchTemperature(parameterizable parameterizable.GETParameterizable) (float64, error) {
	client := &http.Client{}
	urlRequest := OpenWeatherURL + parameterizable.GETParameter() + "&appid=" + w.apiKey
	println(urlRequest)
	req, err := http.NewRequest("GET", urlRequest, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, clients.ClientErr{Status: clients.ErrorFetching, Message: "Error while fetching, got " + resp.Status}
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var ret weatherReturn
	err = json.Unmarshal(byteValue, &ret)
	if err != nil {
		return 0, err
	}

	return KelvinToCelsius(ret.Temperature), nil
}
