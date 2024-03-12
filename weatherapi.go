package weatherapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const weatherBaseUrl = "https://api.weather.yandex.ru/v2/informers?lat=52.339204&lon=35.350873"

var weatherKey string

func InitWehather(key string) {
	weatherKey = key
}

func getBodyByUrl(url string) []byte {

	client := &http.Client{}
	request, err := http.NewRequest(
		"GET", url, nil,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Add("X-Yandex-API-Key", weatherKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}

func (weather *WeatherQueryT) GetWeather() {
	body := getBodyByUrl(weatherBaseUrl)
	err := json.Unmarshal(body, weather)
	if err != nil {
		fmt.Printf("Error unmarshalling weather type: %s", err.Error())
	}
}
