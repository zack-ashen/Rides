package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"rides/middleware"
	"rides/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type OrgCredentials struct {
	OrgID    string `json:"orgID"`
	Password string `json:"password"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("whats up")
}

// Register Organization
func RegisterOrg(org models.Organization) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(org.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	org.Password = string(hashedPassword)
	models.CreateOrg(org)
}

// Authenticate Organization
func org(w http.ResponseWriter, r *http.Request) {
	var creds OrgCredentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	org, err := models.FindOrg(creds.OrgID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Organization does not exist"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("%v", err)
		}

		w.Write(jsonResp)
		log.Printf("%v", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(org.Password), []byte(creds.Password))
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Incorrect password"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("%v", err)
		}

		w.Write(jsonResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["token"] = models.OrgToken(org)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("%v", err)
	}
	w.Write(jsonResp)
}

// Routes for Authentication '/api/auth/...'
func AuthRouter(r *mux.Router) {
	r.HandleFunc("/hello", middleware.Auth(hello))
	r.HandleFunc("/org", org).Methods("POST")
}
