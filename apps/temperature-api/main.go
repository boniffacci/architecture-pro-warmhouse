package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`
	Timestamp   string  `json:"timestamp"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SensorID    string  `json:"sensor_id"`
	Description string  `json:"description"`
	SensorType  string  `json:"sensor_type"`
}


func main() {
	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/temperature/", temperatureByIDHandler)

	log.Println("temperature-api running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// /temperature?location=Living%20Room&sensorId=1
func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	writeTemperatureResponse(w, location, sensorID)
}


func temperatureByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("temperatureByIDHandler called with path:", r.URL.Path)
	sensorID := strings.TrimPrefix(r.URL.Path, "/temperature/")
	location := "Unknown"

	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	}

	writeTemperatureResponse(w, location, sensorID)
}

func writeTemperatureResponse(w http.ResponseWriter, location, sensorID string) {
	rand.Seed(time.Now().UnixNano())
	temp := 18.0 + rand.Float64()*(30.0-18.0)

	response := TemperatureResponse{
		Value:       temp,
		Unit:        "°C",
		Timestamp:   time.Now().Format(time.RFC3339),
		Location:    location,
		Status:      "ok",
		SensorID:    sensorID,
		SensorType:  "",
		Description: "Random temperature reading",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
