package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LogData struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Sender string `json:"sender,omitempty"`
}

// controller methods
func logController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var logData LogData
	sdErr := json.NewDecoder(r.Body).Decode(&logData)
	if sdErr != nil {
		fmt.Println("Error in logController: " + sdErr.Error())
		http.Error(w, sdErr.Error(), http.StatusBadRequest)
		return
	}

	if verbose || logData.Level == "INFO" || logData.Level == "WARNING" || logData.Level == "ERROR" {
		fmt.Println("[" + strings.TrimSpace(logData.Level) + "] " + strings.TrimSpace(logData.Message) + " (from " + logData.Sender + ")")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("yay"))
}
