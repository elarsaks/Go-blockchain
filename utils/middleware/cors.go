package utils

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CorsMiddleware returns a new mux.MiddlewareFunc that applies CORS middleware.
func CorsMiddleware() mux.MiddlewareFunc {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},           // Allow requests from any origin
		AllowedMethods:   []string{"GET", "POST"}, // Allow GET and POST requests
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true, // Allow sending of credentials (cookies, headers)
	}

	corsMiddleware := cors.New(corsOptions).Handler

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") // Add Access-Control-Allow-Origin header
			corsMiddleware.ServeHTTP(w, r)
		})
	}
}
