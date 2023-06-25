package utils

import (
	"net/http"

	"github.com/rs/cors"
)

// TODO: Build out router to handle CORS requests

// CorsMiddleware returns a new CORS middleware instance with the specified configuration.
func CorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},           // Allow requests from any origin
		AllowedMethods:   []string{"GET", "POST"}, // Allow GET and POST requests
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true, // Allow sending of credentials (cookies, headers)
	})
}

// CorsHandler wraps the provided http.Handler with CORS middleware.
func CorsHandler(handler http.Handler) http.Handler {
	return CorsMiddleware().Handler(handler)
}
