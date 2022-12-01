package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type LogData struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Sender  string `json:"sender,omitempty"`
}

const (
	Silly string = "silly"
	Debug        = "debug"
	Info         = "info"
	Warn         = "warn"
	Error        = "error"
	Off          = "off"
)

var hasOutputFile = false
var logLevel = Off
var isVerbose = false

func Init(level string, verbose bool) {
	SetLevel(level)
	SetVerbose(verbose)
}

func SetLevel(level string) {
	logLevel = strings.ToLower(level)
}

func SetVerbose(verbose bool) {
	isVerbose = verbose
}

func TryCreateLogFile() {
	if !hasOutputFile && ShouldLog(logLevel) {
		CreateLogFile()
		LogDebug("Logfile created!")
		LogSystemStats()
	}
}

func LogSystemStats() {
	start := time.Now()

	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	h, _ := host.Info()
	c, _ := cpu.Info()
	d, _ := disk.Usage("/")

	LogInfo("SYSTEM INFO")
	LogInfo(fmt.Sprintf("Virtual Memory: %s\n", v))
	LogInfo(fmt.Sprintf("Swap Memory: %s\n", s))
	LogInfo(fmt.Sprintf("Host: %s\n", h))
	LogInfo(fmt.Sprintf("CPU: %s\n", c))
	LogInfo(fmt.Sprintf("Disk: %s\n", d))

	duration := time.Since(start)
	LogInfo("END SYSTEM INFO, took " + duration.String())
}

func OpenLogFolder() {
	LogDebug("Trying to open log folder...")

	err := utils.OpenExplorer("logs")

	if err != nil {
		LogErrorVerbose("Could not open log folder")
		dialogs.ShowError("Log folder could not be opened! Reason: " + err.Error())
	}
}

func CreateLogFile() {
	logspath := filepath.Join(".", "logs")
	err := os.MkdirAll(logspath, os.ModePerm)

	if err != nil {
		log.Fatal("Could not create logs folder, details: " + err.Error())
		return
	}

	var fileName = filepath.Join(logspath, time.Now().Local().Format("2006-01-02T15-04-05")+"_fskneeboard.log")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
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

	if configuredLevel == Silly && (thisLevel == Silly || thisLevel == Debug || thisLevel == Info || thisLevel == Warn || thisLevel == Error) {
		return true
	}

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

	if strings.TrimSpace(sender) != "" {
		logString += " (from " + strings.TrimSpace(sender) + ")"
	}

	if isVerbose || verboseOverride {
		utils.Println(logString)
	}

	if ShouldLog(level) && hasOutputFile {
		log.Println(logString)
	}
}

// SILLY
func LogSillyVerboseOverride(message string, verboseOverride bool) {
	LogMessage(message, Silly, "", verboseOverride)
}

func LogSillyVerbose(message string) {
	LogSillyVerboseOverride(message, true)
}

func LogSilly(message string) {
	LogSillyVerboseOverride(message, false)
}

// DEBUG
func LogDebugVerboseOverride(message string, verboseOverride bool) {
	LogMessage(message, Debug, "", verboseOverride)
}

func LogDebugVerbose(message string) {
	LogDebugVerboseOverride(message, true)
}

func LogDebug(message string) {
	LogDebugVerboseOverride(message, false)
}

// INFO
func LogInfoVerboseOverride(message string, verboseOverride bool) {
	LogMessage(message, Info, "", verboseOverride)
}

func LogInfoVerbose(message string) {
	LogInfoVerboseOverride(message, true)
}

func LogInfo(message string) {
	LogInfoVerboseOverride(message, false)
}

// WARN
func LogWarnVerboseOverride(message string, verboseOverride bool) {
	LogMessage(message, Warn, "", verboseOverride)
}

func LogWarnVerbose(message string) {
	LogWarnVerboseOverride(message, true)
}

func LogWarn(message string) {
	LogWarnVerboseOverride(message, false)
}

// ERROR
func LogErrorVerboseOverride(message string, verboseOverride bool) {
	LogMessage(message, Error, "", verboseOverride)
}

func LogErrorVerbose(message string) {
	LogErrorVerboseOverride(message, true)
}

func LogError(message string) {
	LogErrorVerboseOverride(message, false)
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
