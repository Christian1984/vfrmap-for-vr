package premium

import (
	"net/http"
)

//go:generate go-bindata -pkg premium -o bindata.go -modtime 1 -prefix "../../../_vendor/premium" "../../../_vendor/premium"

type FS struct {
}

func (_ FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "common.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("common.css"))
	case "waypoints.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("waypoints.js"))
	case "charts.html":
		w.Header().Set("Content-Type", "text/html")
		w.Write(MustAsset("charts.html"))
	case "charts.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("charts.js"))
	case "charts.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("charts.css"))
	case "notepad.html":
		w.Header().Set("Content-Type", "text/html")
		w.Write(MustAsset("notepad.html"))
	case "notepad.css":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("notepad.css"))
	case "notepad.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("notepad.js"))
	}
}
