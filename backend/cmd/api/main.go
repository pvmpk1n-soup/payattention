package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthResponse is the JSON body returned by GET /health.
type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// versionResponse is the JSON body returned by GET /version.
type versionResponse struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

// healthHandler reports that the API is running.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status:  "ok",
		Service: "payattention",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("health: failed to write response: %v", err)
	}
}

// versionHandler reports the API's name and version.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	resp := versionResponse{
		Version: "0.1.0",
		Name:    "PayAttention",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("version failed to write response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /version", versionHandler)

	addr := ":8080"
	log.Printf("PayAttention API listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
