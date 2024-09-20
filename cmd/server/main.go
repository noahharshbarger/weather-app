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

// WeatherResponse represents the structure of the weather data to be returned in the response
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
        AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from this origin
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow these HTTP methods
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allow these headers
        ExposedHeaders:   []string{"Link"}, // Expose these headers
        AllowCredentials: true, // Allow credentials
        MaxAge:           300, // Cache the preflight request for 300 seconds
    })

    // Use CORS middleware
    router.Use(c.Handler)

    // Define routes after middleware
    router.Get("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
        // Extract the city parameter from the URL
        city := chi.URLParam(r, "city")
        // Retrieve the OpenWeather API key from environment variables
        apiKey := os.Getenv("OPENWEATHER_API_KEY")

        // Construct the URL for the OpenWeather API request
        url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

        // Make a GET request to the OpenWeather API
        resp, err := http.Get(url)
        if err != nil {
            log.Printf("Failed to fetch weather data: %v", err)
            http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Check if the response status is Unauthorized (401)
        if resp.StatusCode == http.StatusUnauthorized {
            log.Printf("Unauthorized: Invalid API key")
            http.Error(w, "Unauthorized: Invalid API key", http.StatusUnauthorized)
            return
        }

        // Check if the response status is not OK (200)
        if resp.StatusCode != http.StatusOK {
            log.Printf("Non-OK HTTP status: %s", resp.Status)
            http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
            return
        }

        // Decode the JSON response from the API into a map
        var data map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            log.Printf("Failed to decode weather data: %v", err)
            http.Error(w, "Failed to decode weather data", http.StatusInternalServerError)
            return
        }

        // Extract the main weather data
        mainData, ok := data["main"].(map[string]interface{})
        if !ok {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        // Extract the weather array
        weatherArray, ok := data["weather"].([]interface{})
        if !ok || len(weatherArray) == 0 {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        // Extract the first element of the weather array
        weatherData, ok := weatherArray[0].(map[string]interface{})
        if !ok {
            log.Printf("Invalid weather data format: %v", data)
            http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
            return
        }

        // Create a WeatherResponse struct with the extracted data
        weather := WeatherResponse{
            Temperature: mainData["temp"].(float64),
            Description: weatherData["description"].(string),
            Icon:        weatherData["icon"].(string),
        }

        // Set the response header to indicate JSON content
        w.Header().Set("Content-Type", "application/json")
        // Encode the WeatherResponse struct into JSON and write it to the response
        json.NewEncoder(w).Encode(weather)
    })

    // Print a message indicating that the server is running
    fmt.Println("Server is running on port 8080")
    // Start the HTTP server on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}