package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
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

// create a slice of json objects to store the data no type is defined
var data []interface{}

func loadData() error {
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	return nil
}

// return a random RiskPredictionResponse object from the data
func randomObjectHandler() (RiskPredictionResponse, error) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(data))
	randomObject := data[randomIndex]

	// turn random json object into a RiskPredictionResponse object
	var response RiskPredictionResponse
	bytes, err := json.Marshal(randomObject)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(bytes, &response); err != nil {
		return response, err
	}

	return response, nil

}

// StartServer starts the HTTP server on the specified port
func StartServer() {

	// load the data from the JSON file
	if err := loadData(); err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

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

	response, _ := randomObjectHandler()
	// print the response as RiskPredictionResponse
	fmt.Printf("%+v\n", response)

}
