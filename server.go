package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RiskPredictionRequest struct {
	Age           int    `json:"age"`
	Location      string `json:"location"`
	VehicleType   string `json:"vehicleType"`
	AnnualMileage int    `json:"annualMileage"`
}

// RiskPredictionResponse defines the structure of the outgoing response payload
type RiskPredictionResponse struct {
	RiskScore       int      `json:"riskScore"`
	RiskLevel       string   `json:"riskLevel"`
	Recommendations []string `json:"recommendations"`
}

// StartServer starts the HTTP server on the specified port
func StartServer() {
	http.HandleFunc("/api/risk-prediction", RiskPredictionHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func RiskPredictionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// parse the JSON request body
	var req RiskPredictionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	response := RiskPredictionResponse{
		RiskScore: 85,
		RiskLevel: "High",
		Recommendations: []string{
			"Install a vehicle tracking system.",
			"Enroll in a defensive driving course.",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
