package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Struct for weather API response
type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`

	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {

	apiKey := "your_apikey_from_www.weatherapi.com"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Srinagar&aqi=no", apiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	weatherData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON response
	var weather WeatherResponse
	if err := json.Unmarshal(weatherData, &weather); err != nil {
		log.Fatal(err)
	}

	// Print formatted output for Waybar
	fmt.Printf(`{"text": "🌡 %s: %.1f°C"}`, weather.Location.Name, weather.Current.TempC)

}
