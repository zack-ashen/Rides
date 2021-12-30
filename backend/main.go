package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"rides/routes"
)

func main() {
	// Initialize Router
	r := mux.NewRouter()

	// Setup Routes
	routes.AuthRouter(r.PathPrefix("/api/auth").Subrouter())

	log.Fatal(http.ListenAndServe(":8000", r))

}