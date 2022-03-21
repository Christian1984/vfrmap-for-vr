package logger

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vfrmap-for-vr/vfrmap/utils"
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
var logLevel = Off
var isVerbose = false

func Init(level string, verbose bool) {
	logLevel = level
	isVerbose = verbose
}

func CreateLogFile() {
	logspath := filepath.Join(".", "logs")
	err := os.MkdirAll(logspath, os.ModePerm)

	if err != nil {
		log.Fatal("Could not create logs folder, details: " + err.Error())
		return
	}

	var fileName = filepath.Join(logspath, time.Now().Local().Format("2006-01-02T15-04-05") + "_fskneeboard.log")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if (err != nil) {
		log.Fatal("Could not create or open log file " + fileName + ", details: " + err.Error())
		return
	}

	log.SetOutput(file)
	hasOutputFile = true

	log.Println("FSKneeboard Log File, Log-Level: [" + strings.ToUpper(logLevel) + "]")
	log.Println("======================================================================")
}

func ShouldLog(level string) bool {
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

func LogMessage(message string, level string, sender string, verboseOverride bool) {
	logString := "[" + strings.ToUpper(strings.TrimSpace(level)) + "] " + strings.TrimSpace(message)

	if (strings.TrimSpace(sender) != "") {
		logString += " (from " + strings.TrimSpace(sender) + ")"
	}

	if isVerbose || verboseOverride {
		utils.Println(logString)
	}

	if ShouldLog(level) && hasOutputFile {
		log.Println(logString)
	}
}

func LogDebug(message string, verboseOverride bool) {
	LogMessage(message, Debug, "", verboseOverride)
}

func LogInfo(message string, verboseOverride bool) {
	LogMessage(message, Info, "", verboseOverride)
}

func LogWarn(message string, verboseOverride bool) {
	LogMessage(message, Warn, "", verboseOverride)
}

func LogError(message string, verboseOverride bool) {
	LogMessage(message, Error, "", verboseOverride)
}

// controller methods
func LogController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var logData LogData
	sdErr := json.NewDecoder(r.Body).Decode(&logData)
	if sdErr != nil {
		utils.Println("Error in LogController: " + sdErr.Error())
		http.Error(w, sdErr.Error(), http.StatusBadRequest)
		return
	}

	LogMessage(logData.Message, logData.Level, logData.Sender, false)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

func LogLevelController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(logLevel))
}
