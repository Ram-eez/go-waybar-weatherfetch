package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`

	Current struct {
		TempCurrent float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Srinagar&aqi=yes", apiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	weatherData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Raw Response:", string(weatherData))

	var weather WeatherResponse
	if err := json.Unmarshal(weatherData, &weather); err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather.Location.Name, weather.Current.TempCurrent)
}
