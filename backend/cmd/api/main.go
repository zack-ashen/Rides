package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"rides/routes"
)

func main() {
	// Initialize Router
	r := mux.NewRouter()

	// Setup Routes
	routes.AuthRouter(r.PathPrefix("/api/auth").Subrouter())

	handler := handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:8000"}),
	)
	log.Fatal(http.ListenAndServe(":8000", handler(r)))
}
