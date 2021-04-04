package premium

import (
    "net/http"
)

//go:generate go-bindata -pkg premium -o bindata.go -modtime 1 -prefix "../../../_vendor/premium" "../../../_vendor/premium"

type FS struct {
}

func (_ FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "waypoints.js":
            w.Header().Set("Content-Type", "text/javascript")
            w.Write(MustAsset("waypoints.js"))
        case "charts.html":
            w.Header().Set("Content-Type", "text/html")
            w.Write(MustAsset("charts.html"))
        case "charts.css":
            w.Header().Set("Content-Type", "text/css")
            w.Write(MustAsset("charts.css"))
    }
}
