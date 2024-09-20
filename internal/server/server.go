package server

import (
	"weather-app/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Server struct holds the router instance
type Server struct {
    Router *chi.Mux
}

// NewServer initializes a new server instance
func NewServer() *Server {
    r := chi.NewRouter()

    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    }))

	// Define the route for getting weather info
    r.Get("/weather/{city}", handlers.GetWeather)

	// Return a new server instance
    return &Server{Router: r}
}