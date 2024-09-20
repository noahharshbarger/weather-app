package utils

// ParseWeatherData parses the weather data and returns a map
func ParseWeatherData(weatherData map[string]interface{}) map[string]interface{} {
	// Extract the main weather data
    mainData := weatherData["main"].(map[string]interface{})
	// Extract the weather array
    weatherInfo := weatherData["weather"].([]interface{})[0].(map[string]interface{})
	// Return a map with the extracted data
    return map[string]interface{}{
        "temperature": mainData["temp"],
        "description": weatherInfo["description"],
        "icon":        weatherInfo["icon"],
        "city":        weatherData["name"],
    }
}