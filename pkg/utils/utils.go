package utils

func ParseWeatherData(weatherData map[string]interface{}) map[string]interface{} {
    mainData := weatherData["main"].(map[string]interface{})
    weatherInfo := weatherData["weather"].([]interface{})[0].(map[string]interface{})
    return map[string]interface{}{
        "temperature": mainData["temp"],
        "description": weatherInfo["description"],
        "icon":        weatherInfo["icon"],
        "city":        weatherData["name"],
    }
}