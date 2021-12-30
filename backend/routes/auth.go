package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("whats up")
}

// Authenticate Organization
func org(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("whats up")
}

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/org", org).Methods("POST")
}