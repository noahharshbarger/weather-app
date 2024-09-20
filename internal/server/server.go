package server

import (
	"weather-app/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
    Router *chi.Mux
}

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

    r.Get("/weather/{city}", handlers.GetWeather)

    return &Server{Router: r}
}