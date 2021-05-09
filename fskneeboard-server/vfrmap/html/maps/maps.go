package maps

import (
	"net/http"
)

//go:generate go-bindata -pkg maps -o bindata.go -modtime 1 -prefix "html" "html"

type FS struct {
}

func (_ FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "maps.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("maps.css"))
	case "maps.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("maps.js"))
	case "maps.html":
		w.Header().Set("Content-Type", "text/html")
		w.Write(MustAsset("charts.html"))
	}
}
