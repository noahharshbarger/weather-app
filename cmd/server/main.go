package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
    Temperature float64 `json:"temperature"`
    Description string  `json:"description"`
    Icon        string  `json:"icon"`
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Create a new router
    router := chi.NewRouter()

    // CORS configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    })

    // Use CORS middleware
    router.Use(c.Handler)

    // Define routes after middleware
    router.Get("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
        city := chi.URLParam(r, "city")
        apiKey := os.Getenv("OPENWEATHER_API_KEY")

        // Log the API key for debugging (remove this in production)
        log.Printf("Using API Key: %s", apiKey)

        url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

        resp, err := http.Get(url)
        if err != nil {
            log.Printf("Failed to fetch weather data: %v", err)
            http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusUnauthorized {
            log.Printf("Unauthorized: Invalid API key")
            http.Error(w, "Unauthorized: Invalid API key", http.StatusUnauthorized)
            return
        }

        if resp.StatusCode != http.StatusOK {
            log.Printf("Non-OK HTTP status: %s", resp.Status)
            http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
            return
        }

        var data map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            log.Printf("Failed to decode weather data: %v", err)
            http.Error(w, "Failed to decode weather data", http.StatusInternalServerError)
            return
        }

        mainData, ok := data["main"].(map[string]interface{})
        if !ok {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        weatherArray, ok := data["weather"].([]interface{})
        if !ok || len(weatherArray) == 0 {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        weatherData, ok := weatherArray[0].(map[string]interface{})
        if !ok {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        weather := WeatherResponse{
            Temperature: mainData["temp"].(float64),
            Description: weatherData["description"].(string),
            Icon:        weatherData["icon"].(string),
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(weather)
    })

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
