package waypoints

import (
	"net/http"
	"vfrmap-for-vr/simconnect"
)

func SendFlightplanResponse(filepath string) {
	/* intentionally left empty */
}

func LocateCurrentFlightplan(s *simconnect.SimConnect, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}
