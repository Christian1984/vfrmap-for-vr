package waypoints

import (
	"net/http"
)

func GetFlightplan(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}
