package models

// WeatherData struct holds the weather data
type WeatherData struct {
    Temperature string `json:"temperature"`
    Description string `json:"description"`
    Icon        string `json:"icon"`
    City        string `json:"city"`
}