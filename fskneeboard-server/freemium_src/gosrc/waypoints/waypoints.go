package waypoints

import (
	"context"
	"net/http"
	"vfrmap-for-vr/simconnect"
)

func SendFlightplanResponse(filepath string) {
	/* intentionally left empty */
}

func LocateCurrentFlightplan(s *simconnect.SimConnect, w http.ResponseWriter, r *http.Request, ctx context.Context, ch chan<- string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
	ch <- "success"
	close(ch)
}
