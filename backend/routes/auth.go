package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
