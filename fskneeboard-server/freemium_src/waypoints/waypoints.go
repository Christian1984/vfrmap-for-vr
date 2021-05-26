package waypoints

import (
	"net/http"
)

func GetFlightplan(w http.ResponseWriter, r *http.Request, filepath string) {
	w.Write([]byte("{}"))
}
