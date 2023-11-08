
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-ai-microservices/pkg/model"
	"go-ai-microservices/pkg/inference"
)

// PredictionRequest represents the structure of an incoming prediction request
type PredictionRequest struct {
	Data []float64 `json:"data"`
}

// PredictionResponse represents the structure of an outgoing prediction response
type PredictionResponse struct {
	Prediction float64 `json:"prediction"`
	Error      string  `json:"error,omitempty"`
}

var aiModel *model.SimpleModel

func predictHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PredictionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if aiModel == nil {
		http.Error(w, "AI model not loaded", http.StatusInternalServerError)
		return
	}

	prediction, err := inference.Predict(aiModel, req.Data)
	if err != nil {
		resp := PredictionResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := PredictionResponse{Prediction: prediction}
	json.NewEncoder(w).Encode(resp)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service is healthy!")
}

func main() {
	fmt.Println("Starting AI Microservice...")

	// Initialize a dummy AI model (e.g., a simple linear regression)
	aiModel = model.NewSimpleModel([]float64{0.5, 0.3}, 0.1) // Example weights and bias

	r := mux.NewRouter()
	r.HandleFunc("/predict", predictHandler).Methods("POST")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
