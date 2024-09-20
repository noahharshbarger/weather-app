package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"weather-app/pkg/utils"

	"github.com/go-chi/chi"
)

// Retrieve the OpenWeather API key from the environ variables
var apiKey = os.Getenv("OPENWEATHER_API_KEY")

// GetWeather function fetches weather data from the OpenWeather API
func GetWeather(w http.ResponseWriter, r *http.Request) {
	// Get the city parameter from the URL
    city := chi.URLParam(r, "city")
	// Make a GET request to the OpenWeather API
    response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer response.Body.Close()

	// Check if response status is not OK
    if response.StatusCode != http.StatusOK {
        http.Error(w, "Failed to get weather data", response.StatusCode)
        return
    }

	// Decode the JSON response
    var weatherData map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// Parse the weather data
    parsedData := utils.ParseWeatherData(weatherData)

	// Set the response content type to JSON
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(parsedData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}