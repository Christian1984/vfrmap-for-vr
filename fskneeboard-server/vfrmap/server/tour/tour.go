package tour

import (
	"net/http"
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/websockets"
)

var Ws *websockets.Websocket

func ServeTourStatus(w http.ResponseWriter, r *http.Request) {
	logger.LogDebug("ServeTourStatus called!")

	if r.Method != http.MethodGet {
		logger.LogError("Method " + r.Method + " not allowed!")
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	pathArr := strings.Split(r.URL.Path, "/")

	if len(pathArr) < 3 {
		logger.LogError("Tour status requested with invalid path: " + r.URL.Path)
		http.Error(w, "Tour status requested with invalid path! Requested status not found", http.StatusNotFound)
	}

	response := ""

	switch pathArr[2] {
	case "indexStarted":
		response = strconv.FormatBool(globals.TourIndexStarted)
		globals.TourIndexStarted = true
		break
	case "mapStarted":
		response = strconv.FormatBool(globals.TourMapStarted)
		globals.TourMapStarted = true
		break
	case "chartsStarted":
		response = strconv.FormatBool(globals.TourChartsStarted)
		globals.TourChartsStarted = true
		break
	case "notepadStarted":
		response = strconv.FormatBool(globals.TourNotepadStarted)
		globals.TourNotepadStarted = true
		break
	default:
		logger.LogError("Tour status requested invalid: " + pathArr[1])
		http.Error(w, "Tour status requested with invalid path! Requested status not found.", http.StatusNotFound)
	}

	dbmanager.StoreTourStates()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
