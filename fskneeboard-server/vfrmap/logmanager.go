package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type LogData struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Sender  string `json:"sender,omitempty"`
}

const (
	Debug string = "debug"
	Info         = "info"
	Warn         = "warn"
	Error        = "error"
	Off          = "off"
)

var hasOutputFile = false

func createLogFile() {
	var fileName = time.Now().Local().Format("2006-01-02T15-04-05") + "_fskneeboard.log"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if (err != nil) {
		log.Fatal("Could not create or open log file " + fileName + ", details: " + err.Error())
		return
	}

	log.SetOutput(file)
	hasOutputFile = true
}

func shouldLog(level string) bool {
	var configuredLevel = strings.ToLower(logLevel)
	var thisLevel = strings.ToLower(level)

	if configuredLevel == Debug && (thisLevel == Debug || thisLevel == Info || thisLevel == Warn || thisLevel == Error) {
		return true
	}

	if configuredLevel == Info && (thisLevel == Info || thisLevel == Warn || thisLevel == Error) {
		return true
	}

	if configuredLevel == Warn && (thisLevel == Warn || thisLevel == Error) {
		return true
	}

	if configuredLevel == Error && thisLevel == Error {
		return true
	}

	return false
}

func logMessageWithSender(message string, level string, sender string) {
	if !shouldLog(level) {
		return
	}

	logString := "[" + strings.ToUpper(strings.TrimSpace(level)) + "] " + strings.TrimSpace(message)

	if (strings.TrimSpace(sender) != "") {
		logString += " (from " + strings.TrimSpace(sender) + ")"
	}

	if verbose {
		fmt.Println(logString)
	}

	if hasOutputFile {
		log.Println(logString)
	}
}

func logMessage(message string, level string) {
	logMessageWithSender(message, level, "")
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

	logMessageWithSender(logData.Message, logData.Level, logData.Sender)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("yay"))
}
