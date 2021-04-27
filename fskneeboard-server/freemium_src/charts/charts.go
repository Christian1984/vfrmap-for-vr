package charts

import (
	"net/http"
)

func UpdateIndex() []string {
	return []string{}
}

func Json(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
