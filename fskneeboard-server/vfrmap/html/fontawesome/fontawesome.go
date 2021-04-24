package fontawesome

import (
	"net/http"
)

//go:generate go-bindata -pkg fontawesome -o bindata.go -modtime 1 -prefix "../../../_vendor/fontawesome" "../../../_vendor/fontawesome" "../../../_vendor/fontawesome/css" "../../../_vendor/fontawesome/webfonts" "../../../_vendor/fontawesome/js"

type FS struct {
}

func (_ FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "js/all.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("js/all.js"))
	case "js/all.min.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("js/all.min.js"))
	case "css/all.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("css/all.css"))
	case "css/all.min.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("css/all.min.css"))
	case "webfonts/fa-brands-400.eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
		w.Write(MustAsset("webfonts/fa-brands-400.eot"))
	case "webfonts/fa-regular-400.eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
		w.Write(MustAsset("webfonts/fa-regular-400.eot"))
	case "webfonts/fa-solid-900.eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
		w.Write(MustAsset("webfonts/fa-solid-900.eot"))
	case "webfonts/fa-brands-400.svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(MustAsset("webfonts/fa-brands-400.svg"))
	case "webfonts/fa-regular-400.svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(MustAsset("webfonts/fa-regular-400.svg"))
	case "webfonts/fa-solid-900.svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(MustAsset("webfonts/fa-solid-900.svg"))
	case "webfonts/fa-brands-400.ttf":
		w.Header().Set("Content-Type", "font/ttf")
		w.Write(MustAsset("webfonts/fa-brands-400.ttf"))
	case "webfonts/fa-regular-400.ttf":
		w.Header().Set("Content-Type", "font/ttf")
		w.Write(MustAsset("webfonts/fa-regular-400.ttf"))
	case "webfonts/fa-solid-900.ttf":
		w.Header().Set("Content-Type", "font/ttf")
		w.Write(MustAsset("webfonts/fa-solid-900.ttf"))
	case "webfonts/fa-brands-400.woff":
		w.Header().Set("Content-Type", "font/woff")
		w.Write(MustAsset("webfonts/fa-brands-400.woff"))
	case "webfonts/fa-regular-400.woff":
		w.Header().Set("Content-Type", "font/woff")
		w.Write(MustAsset("webfonts/fa-regular-400.woff"))
	case "webfonts/fa-solid-900.woff":
		w.Header().Set("Content-Type", "font/woff")
		w.Write(MustAsset("webfonts/fa-solid-900.woff"))
	case "webfonts/fa-brands-400.woff2":
		w.Header().Set("Content-Type", "font/woff2")
		w.Write(MustAsset("webfonts/fa-brands-400.woff2"))
	case "webfonts/fa-regular-400.woff2":
		w.Header().Set("Content-Type", "font/woff2")
		w.Write(MustAsset("webfonts/fa-regular-400.woff2"))
	case "webfonts/fa-solid-900.woff2":
		w.Header().Set("Content-Type", "font/woff2")
		w.Write(MustAsset("webfonts/fa-solid-900.woff2"))
	}
}
