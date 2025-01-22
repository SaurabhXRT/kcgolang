package main

import (
	"encoding/json"
	"kcassignment/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := HealthResponse{
			Message: "Server is running",
			Status:  "ok",
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")


	r.HandleFunc("/api/submit", api.SubmitJobHandler).Methods("POST")
	r.HandleFunc("/api/status", api.GetJobStatusHandler).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
