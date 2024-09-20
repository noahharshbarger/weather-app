package handlers

import (
	"encoding/json"
	"net/http"
	"weather-app/pkg/utils"

	"github.com/go-chi/chi"
)

const apiKey = "197a9a199b3b20586cab551634761e97"

func GetWeather(w http.ResponseWriter, r *http.Request) {
    city := chi.URLParam(r, "city")
    response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        http.Error(w, "Failed to get weather data", response.StatusCode)
        return
    }

    var weatherData map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    parsedData := utils.ParseWeatherData(weatherData)

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(parsedData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}