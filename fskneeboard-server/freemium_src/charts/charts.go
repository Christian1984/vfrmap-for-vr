package charts

import (
	"net/http"
)

type FileIndex struct {
	Files []string
}

func UpdateIndex() []string {
	return []string{}
}

func Json(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
