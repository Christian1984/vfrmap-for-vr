package server

//go:generate go-bindata -pkg server -o bindata.go -modtime 1 -prefix ../html/webdist ../html/webdist

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
	"vfrmap-for-vr/_vendor/premium/autosave"
	"vfrmap-for-vr/_vendor/premium/charts"
	"vfrmap-for-vr/_vendor/premium/notepad"
	"vfrmap-for-vr/_vendor/premium/waypoints"
	"vfrmap-for-vr/simconnect"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/secrets"
	"vfrmap-for-vr/vfrmap/gui/callbacks"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/html/fontawesome"
	"vfrmap-for-vr/vfrmap/html/freemium"
	"vfrmap-for-vr/vfrmap/html/leafletjs"
	"vfrmap-for-vr/vfrmap/html/premium"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server/hotkeys"
	"vfrmap-for-vr/vfrmap/server/tour"
	"vfrmap-for-vr/vfrmap/utils"
	"vfrmap-for-vr/vfrmap/websockets"

	"github.com/Christian1984/go-maptilecache"
)

var started = false
var autosaveTick = time.NewTicker(9999 * time.Minute)

var mockV = 0.005
var mockH = math.Pi / 4

const keepTurnDirChance = 0.995
const deltaMockH = 0.05

var mockLat = 50.8694
var mockLng = 7.1389
var turnLeft = true

const trailDataHdCapacity = 3000
const trailDataSdResolution = 20

var trail = Trail{
	TrailDataHd: [][]float64{},
	TrailDataSd: [][]float64{},
}

type Trail struct {
	TrailDataHd [][]float64 `json:"TrailDataHd"`
	TrailDataSd [][]float64 `json:"TrailDataSd"`
}

type Report struct {
	simconnect.RecvSimobjectDataByType
	Title         [256]byte `name:"TITLE"`
	Altitude      float64   `name:"INDICATED ALTITUDE" unit:"feet"` // PLANE ALTITUDE or PLANE ALT ABOVE GROUND
	Latitude      float64   `name:"PLANE LATITUDE" unit:"degrees"`
	Longitude     float64   `name:"PLANE LONGITUDE" unit:"degrees"`
	Heading       float64   `name:"PLANE HEADING DEGREES TRUE" unit:"degrees"`
	Airspeed      float64   `name:"AIRSPEED INDICATED" unit:"knot"`
	AirspeedTrue  float64   `name:"AIRSPEED TRUE" unit:"knot"`
	VerticalSpeed float64   `name:"VERTICAL SPEED" unit:"ft/min"`
	Flaps         float64   `name:"TRAILING EDGE FLAPS LEFT ANGLE" unit:"degrees"`
	Trim          float64   `name:"ELEVATOR TRIM PCT" unit:"percent"`
	RudderTrim    float64   `name:"RUDDER TRIM PCT" unit:"percent"`
	WindDirection float64   `name:"AMBIENT WIND DIRECTION" unit:"degrees"`
	WindVelocity  float64   `name:"AMBIENT WIND VELOCITY" unit:"knots"`
}

type MapServiceUrls struct {
	CacheUrl  string
	RemoteUrl string
}

var OsmUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35302/maptilecache/osm/{s}/{z}/{y}/{x}/",
	RemoteUrl: "http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
}

var OtmUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35303/maptilecache/otm/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png",
}

var StamenBwUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35304/maptilecache/stamenbw/{s}/{z}/{y}/{x}/",
	RemoteUrl: "http://{s}.tile.stamen.com/toner/{z}/{x}/{y}.png",
}

var StamenTUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35305/maptilecache/stament/{s}/{z}/{y}/{x}/",
	RemoteUrl: "http://{s}.tile.stamen.com/terrain/{z}/{x}/{y}.png",
}

var StamenWUrls = MapServiceUrls{
	CacheUrl:  "",
	RemoteUrl: "http://{s}.tile.stamen.com/watercolor/{z}/{x}/{y}.png",
}

var CartoD = MapServiceUrls{
	CacheUrl:  "http://localhost:35307/maptilecache/cartod/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://cartodb-basemaps-{s}.global.ssl.fastly.net/dark_all/{z}/{x}/{y}.png",
}

var Ofm = MapServiceUrls{
	CacheUrl:  "http://localhost:35308/maptilecache/ofm/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://nwy-tiles-api.prod.newaydata.com/tiles/{z}/{x}/{y}.png",
}

var OaipAirportsUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35309/maptilecache/oaip-airports/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://api.tiles.openaip.net/api/data/airports/{z}/{x}/{y}.png?apiKey={apiKey}",
}

var OaipAirspacesUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35310/maptilecache/oaip-airspaces/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://api.tiles.openaip.net/api/data/airspaces/{z}/{x}/{y}.png?apiKey={apiKey}",
}

var OaipNavaidsUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35311/maptilecache/oaip-navaids/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://api.tiles.openaip.net/api/data/navaids/{z}/{x}/{y}.png?apiKey={apiKey}",
}

var OaipReportingUrls = MapServiceUrls{
	CacheUrl:  "http://localhost:35312/maptilecache/oaip-reportingpoints/{s}/{z}/{y}/{x}/",
	RemoteUrl: "https://api.tiles.openaip.net/api/data/reporting-points/{z}/{x}/{y}.png?apiKey={apiKey}",
}

type MapServiceUrlsDto struct {
	Osm      string `json:"osm"`
	Otm      string `json:"otm"`
	StamenBw string `json:"stamenbw"`
	StamenT  string `json:"stament"`
	StamenW  string `json:"stamenw"`
	CartoD   string `json:"cartod"`
	//Ofm           string `json:"ofm"` // does not work because of airac cycle
	OaipAirports  string `json:"oaipAirports"`
	OaipAirspaces string `json:"oaipAirspaces"`
	OaipNavaids   string `json:"oaipNavaids"`
	OaipReporting string `json:"oaipReporting"`
}

func (r *Report) RequestData(s *simconnect.SimConnect) {
	logger.LogSilly("requesting simconnect data...")

	defineID := s.GetDefineID(r)
	logger.LogSilly("simconnect defineID: " + strconv.FormatUint(uint64(defineID), 10))

	requestID := defineID
	logger.LogSilly("simconnect requestID: " + strconv.FormatUint(uint64(requestID), 10))

	err := s.RequestDataOnSimObjectType(requestID, defineID, 0, simconnect.SIMOBJECT_TYPE_USER)
	if (err != nil) {
		logger.LogError("could not request simconnect data, reason: " + err.Error())
	}
}

func (r *Report) MockData() {
	if rand.Float32() > keepTurnDirChance {
		turnLeft = !turnLeft
	}

	dir := 1.0
	if turnLeft {
		dir = -1.0
	}

	mockH += dir * (rand.Float64()*3*deltaMockH - deltaMockH)
	hdgDeg := mockH * 180 / math.Pi

	mockLng += mockV * math.Sin(mockH)
	mockLat += mockV * math.Cos(mockH)

	//fmt.Println("hdg:", hdgDeg, ", lat/lng:", mockLat, "/", mockLng)

	r.Latitude = mockLat
	r.Longitude = mockLng
	r.Heading = hdgDeg

	r.Altitude = 5000
	r.Airspeed = 500
	r.AirspeedTrue = 500
	r.WindDirection = 10
}

type TrafficReport struct {
	simconnect.RecvSimobjectDataByType
	AtcID           [64]byte `name:"ATC ID"`
	AtcFlightNumber [8]byte  `name:"ATC FLIGHT NUMBER"`
	Altitude        float64  `name:"PLANE ALTITUDE" unit:"feet"`
	Latitude        float64  `name:"PLANE LATITUDE" unit:"degrees"`
	Longitude       float64  `name:"PLANE LONGITUDE" unit:"degrees"`
	Heading         float64  `name:"PLANE HEADING DEGREES TRUE" unit:"degrees"`
}

func (r *TrafficReport) RequestData(s *simconnect.SimConnect) {
	defineID := s.GetDefineID(r)
	requestID := defineID
	s.RequestDataOnSimObjectType(requestID, defineID, 0, simconnect.SIMOBJECT_TYPE_AIRCRAFT)
}

func (r *TrafficReport) Inspect() string {
	return fmt.Sprintf(
		"%s GPS %.6f %.6f @ %.0f feet %.0f°",
		r.AtcID,
		r.Latitude,
		r.Longitude,
		r.Altitude,
		r.Heading,
	)
}

type TeleportRequest struct {
	simconnect.RecvSimobjectDataByType
	Latitude  float64 `name:"PLANE LATITUDE" unit:"degrees"`
	Longitude float64 `name:"PLANE LONGITUDE" unit:"degrees"`
	Altitude  float64 `name:"PLANE ALTITUDE" unit:"feet"`
}

func (r *TeleportRequest) SetData(s *simconnect.SimConnect) {
	defineID := s.GetDefineID(r)

	buf := [3]float64{
		r.Latitude,
		r.Longitude,
		r.Altitude,
	}

	size := simconnect.DWORD(3 * 8) // 2 * 8 bytes
	s.SetDataOnSimObject(defineID, simconnect.OBJECT_ID_USER, 0, 0, size, unsafe.Pointer(&buf[0]))
}

func trailDataController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodDelete {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		responseJson, jsonErr := json.Marshal(trail)

		if jsonErr != nil {
			logger.LogError(jsonErr.Error())
			http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJson))
	} else {
		trail.TrailDataHd = [][]float64{}
		trail.TrailDataSd = [][]float64{}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(""))
	}
}

var prevLat float64
var prevLng float64

const MSFS_ZERO_LAT = 0.000407
const MSFS_ZERO_LNG = 0.013975

func (r *Report) process(ws *websockets.Websocket) {
	addToTrail := true
	logger.LogSilly("trail lat: " + fmt.Sprintf("%f", r.Latitude) + ", lng: " + fmt.Sprintf("%f", r.Longitude))

	if r.Latitude == 0 && r.Longitude == 0 {
		logger.LogSilly("lat and lng are 0, ignoring...")
		addToTrail = false
	}

	if math.Abs(r.Latitude-MSFS_ZERO_LAT) < 1e-6 && math.Abs(r.Longitude-MSFS_ZERO_LNG) < 1e-6 {
		logger.LogSilly("lat and lng are very close to MSFS defaults! ignoring...")
		addToTrail = false
	}

	// check if position has ever changed
	if prevLat < 1e-9 {
		prevLat = r.Latitude
	}

	if prevLng < 1e-9 {
		prevLng = r.Longitude
	}

	if addToTrail {
		if math.Abs(r.Latitude-prevLat) > 1e-6 || math.Abs(r.Longitude-prevLng) > 1e-6 {
			logger.LogSilly("lat and/or lng are far enough apart, adding them to history")

			td := []float64{
				math.Round(r.Latitude*1000000) / 1000000,
				math.Round(r.Longitude*1000000) / 1000000,
			}
			trail.TrailDataHd = append(trail.TrailDataHd, td)

			if len(trail.TrailDataSd) == 0 {
				trail.TrailDataSd = append(trail.TrailDataSd, td)
			}

			if len(trail.TrailDataHd) > trailDataHdCapacity {
				trail.TrailDataSd = append(trail.TrailDataSd, trail.TrailDataHd[0])
				trail.TrailDataHd = trail.TrailDataHd[trailDataSdResolution:]
			}
		} else {
			logger.LogSilly("lat and/or lng are very close, NOT adding them to history")
			addToTrail = false
		}
	}

	prevLat = r.Latitude
	prevLng = r.Longitude

	ws.Broadcast(map[string]interface{}{
		"type":           "plane",
		"latitude":       r.Latitude,
		"longitude":      r.Longitude,
		"altitude":       fmt.Sprintf("%.0f", r.Altitude),
		"heading":        r.Heading,
		"airspeed":       fmt.Sprintf("%.0f", r.Airspeed),
		"airspeed_true":  fmt.Sprintf("%.0f", r.AirspeedTrue),
		"vertical_speed": fmt.Sprintf("%.0f", r.VerticalSpeed),
		"flaps":          fmt.Sprintf("%.0f", r.Flaps),
		"trim":           fmt.Sprintf("%.1f", r.Trim),
		"rudder_trim":    fmt.Sprintf("%.1f", r.RudderTrim),
		"wind_direction": fmt.Sprintf("%.0f", r.WindDirection),
		"wind_velocity":  fmt.Sprintf("%.0f", r.WindVelocity),
		"add_to_trail":   addToTrail,
	})

}

func ShutdownWithPrompt() {
	if globals.Quietshutdown {
		os.Exit(0)
	} else {
		dialogs.ShowMsfsShutdownInfoAndExit()
	}
}

func UpdateAutosaveInterval(verbose bool) {
	if autosaveTick != nil {
		logger.LogDebug("Autosave interval updated: Stopping old timer...")
		autosaveTick.Stop()
	}

	if verbose {
		utils.Println("=== INFO: Autosave")
	}

	if globals.AutosaveInterval > 0 {
		if verbose {
			utils.Printf("Autosave Interval set to %d minute(s)...\n", globals.AutosaveInterval)
		}

		logger.LogInfo("Autosave interval updated: Creating new ticker with an interval of " + strconv.Itoa(globals.AutosaveInterval) + " minutes")
		autosaveTick = time.NewTicker(time.Duration(globals.AutosaveInterval) * time.Minute)
	} else {
		if verbose {
			utils.Println("Autosave deactivated. Please configure the autosave interval in the settings section.")
		}

		logger.LogInfo("Autosave interval disabled: Creating new ticker with an interval of 9999 minutes")
		autosaveTick = time.NewTicker(9999 * time.Minute)
	}

	utils.Println("")

	callbacks.UpdateAutosaveStatus(globals.AutosaveInterval)
}

func initCache(ttl time.Duration, root string, provider string, port string, url string, apiKey string, forwardHeaders bool, expectedParams []string, sharedMemoryCache *maptilecache.SharedMemoryCache) (*maptilecache.Cache, error) {
	cacheConfig := maptilecache.CacheConfig{
		Host:              "0.0.0.0",
		Port:              port,
		Route:             []string{root, provider},
		UrlScheme:         url,
		StructureParams:   expectedParams,
		ApiKey:            apiKey,
		ForwardHeaders:    forwardHeaders,
		TimeToLive:        ttl,
		SharedMemoryCache: sharedMemoryCache,
		DebugLogger:       logger.LogDebug,
		InfoLogger:        logger.LogInfo,
		WarnLogger:        logger.LogWarn,
		ErrorLogger:       logger.LogError,
		StatsLogDelay:     globals.MaptileCacheStatsLogDelay,
	}

	c, err := maptilecache.New(cacheConfig)

	if err == nil {
		if globals.WipeMaptileCaches {
			c.WipeCache()
		} else {
			go func() {
				c.ValidateCache()
				c.PreloadMemoryMap()
			}()
		}
	} else {
		logger.LogError("An error was raised during the initialization of maptilecache [" + provider + "], reason: " + err.Error())
	}

	return c, err
}

func UpdateCacheApiKeys() {
	logger.LogDebug("Enter UpdateCacheApiKeys() with " + globals.OpenAipApiKey)

	oaipApiKey := globals.OpenAipApiKey
	oaipLog := oaipApiKey
	if oaipApiKey == "" {
		oaipApiKey = secrets.API_KEY_OPENAIP
		oaipLog = "DEFAULT API KEY"
	}

	for _, cache := range globals.OpenAipCaches {
		logger.LogDebug("Update api key on cache " + cache.UrlScheme + " from " + cache.ApiKey + " to " + oaipLog)
		cache.ApiKey = oaipApiKey
	}
}

func initMaptileCache() {
	ttl := globals.MaptileCacheTimeToLiveDefault

	globalRoot := "maptilecache"

	sharedMemoryCacheConfig := maptilecache.SharedMemoryCacheConfig{
		MaxSizeBytes:          globals.MaptileCacheMaxMemoryUsage,
		EnsureMaxSizeInterval: 1 * time.Minute,
		DebugLogger:           logger.LogDebug,
		InfoLogger:            logger.LogInfo,
		WarnLogger:            logger.LogWarn,
		ErrorLogger:           logger.LogError,
	}
	sharedMemoryCache := maptilecache.NewSharedMemoryCache(sharedMemoryCacheConfig)

	initCache(ttl, globalRoot, "osm", "35302", OsmUrls.RemoteUrl, "", true, []string{}, sharedMemoryCache)
	initCache(ttl, globalRoot, "otm", "35303", OtmUrls.RemoteUrl, "", true, []string{}, sharedMemoryCache)
	initCache(ttl, globalRoot, "stamenbw", "35304", StamenBwUrls.RemoteUrl, "", true, []string{}, sharedMemoryCache)
	initCache(ttl, globalRoot, "stament", "35305", StamenTUrls.RemoteUrl, "", true, []string{}, sharedMemoryCache)
	initCache(ttl, globalRoot, "cartod", "35307", CartoD.RemoteUrl, "", true, []string{}, sharedMemoryCache)

	initCache(ttl, globalRoot, "ofm", "35308", Ofm.RemoteUrl, "", true, []string{"path"}, sharedMemoryCache)

	var oaipCaches = []*maptilecache.Cache{}

	oaipApiKey := strings.TrimSpace(globals.OpenAipApiKey)
	oaipApiKeyLog := oaipApiKey
	if oaipApiKey == "" {
		oaipApiKey = secrets.API_KEY_OPENAIP
		oaipApiKeyLog = "DEFAULT API KEY"
	}

	logger.LogDebug("Initializing OAIP caches with api key: [" + oaipApiKeyLog + "]")

	oaipAirports, oaipAirportsErr := initCache(ttl, globalRoot, "oaip-airports", "35309", OaipAirportsUrls.RemoteUrl, oaipApiKey, false, []string{}, sharedMemoryCache)
	if oaipAirportsErr == nil {
		oaipCaches = append(oaipCaches, oaipAirports)
	}

	oaipAirspaces, oaipAirspacesErr := initCache(ttl, globalRoot, "oaip-airspaces", "35310", OaipAirspacesUrls.RemoteUrl, oaipApiKey, false, []string{}, sharedMemoryCache)
	if oaipAirspacesErr == nil {
		oaipCaches = append(oaipCaches, oaipAirspaces)
	}

	oaipNavaids, oaipNavaidsErr := initCache(ttl, globalRoot, "oaip-navaids", "35311", OaipNavaidsUrls.RemoteUrl, oaipApiKey, false, []string{}, sharedMemoryCache)
	if oaipNavaidsErr == nil {
		oaipCaches = append(oaipCaches, oaipNavaids)
	}

	oaipReporting, oaipReportingErr := initCache(ttl, globalRoot, "oaip-reportingpoints", "35312", OaipReportingUrls.RemoteUrl, oaipApiKey, false, []string{}, sharedMemoryCache)
	if oaipReportingErr == nil {
		oaipCaches = append(oaipCaches, oaipReporting)
	}

	globals.OpenAipCaches = oaipCaches

	//initCache(ttl, globalRoot, "oaip-obstacles", "35313", "https://api.tiles.openaip.net/api/data/obstacles/{z}/{x}/{y}.png?apiKey={apiKey}", globals.MaptileCacheOaipApiKey, false, []string{}, sharedMemoryCache)
}

func serveMapServiceUrls(w http.ResponseWriter, r *http.Request) {
	logger.LogDebug("serveMapServiceUrls called!")

	oaipAirportsUrl := OaipAirportsUrls.CacheUrl
	oaipAirspacesUrl := OaipAirspacesUrls.CacheUrl
	oaipNavaidsUrl := OaipNavaidsUrls.CacheUrl
	oaipReportingUrl := OaipReportingUrls.CacheUrl

	if globals.OpenAipBypassCache && globals.OpenAipApiKey != "" && globals.OpenAipApiKey != secrets.API_KEY_OPENAIP {
		logger.LogDebug("serving openAIP remote urls")

		oaipAirportsUrl = strings.ReplaceAll(OaipAirportsUrls.RemoteUrl, "{apiKey}", globals.OpenAipApiKey)
		oaipAirspacesUrl = strings.ReplaceAll(OaipAirspacesUrls.RemoteUrl, "{apiKey}", globals.OpenAipApiKey)
		oaipNavaidsUrl = strings.ReplaceAll(OaipNavaidsUrls.RemoteUrl, "{apiKey}", globals.OpenAipApiKey)
		oaipReportingUrl = strings.ReplaceAll(OaipReportingUrls.RemoteUrl, "{apiKey}", globals.OpenAipApiKey)
	} else {
		logger.LogDebug("serving openAIP cache urls")
	}

	mapServiceUrls := MapServiceUrlsDto{
		Osm:      OsmUrls.CacheUrl,
		Otm:      OtmUrls.CacheUrl,
		StamenBw: StamenBwUrls.CacheUrl,
		StamenT:  StamenTUrls.CacheUrl,
		StamenW:  StamenWUrls.RemoteUrl, // cache was buggy
		CartoD:   CartoD.CacheUrl,

		OaipAirports:  oaipAirportsUrl,
		OaipAirspaces: oaipAirspacesUrl,
		OaipNavaids:   oaipNavaidsUrl,
		OaipReporting: oaipReportingUrl,
	}

	responseJson, jsonErr := json.Marshal(mapServiceUrls)

	if jsonErr != nil {
		logger.LogError(jsonErr.Error())
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responseJson))
}

func StartFskServer() {
	logger.LogDebug("Server startup: enter StartFskServer()...")

	if globals.Pro && !globals.DrmValid {
		dialogs.ShowLicenseError()
		utils.Println("WARNING: Cannot start FSKneeboard server, reason: no valid license found!")
		logger.LogWarnVerboseOverride("Cannot start server, reason: no valid license found", false)
		return
	}

	if started {
		logger.LogWarn("Server startup: Server has already been started... Returning...")
		return
	}

	started = true

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)

	logger.LogDebug("Server startup: Waiting 1 s...")
	time.Sleep(1 * time.Second)

	exePath, _ := os.Executable()

	// wait for Flight Simulator
	var s *simconnect.SimConnect
	var err error

	logger.LogInfoVerboseOverride("Waiting for Flight Simulator...", false)
	utils.Println("")
	utils.Print("Waiting for Flight Simulator..")
	callbacks.UpdateServerStatus("Waiting for Flight Simulator...", "")
	callbacks.UpdateMsfsConnectionStatus("Connecting...")

	for true {
		utils.Print(".")

		s, err = simconnect.New(globals.ProductName)

		if err != nil {

			if globals.DevMode {
				logger.LogInfoVerboseOverride("Running with --dev: Not connected to Flight Simulator!", false)
				utils.Println("")
				utils.Println("Running with --dev: Not connected to Flight Simulator!!!")
				utils.Println("")

				callbacks.UpdateMsfsConnectionStatus("Not Connected (dev Mode)")

				break
			}

			time.Sleep(5 * time.Second)
		} else if s != nil {
			logger.LogInfoVerboseOverride("Connected to Flight Simulator!", false)
			utils.Println("")
			utils.Println("Connected to Flight Simulator!")
			utils.Println("")
			callbacks.UpdateMsfsConnectionStatus("Connected")
			break
		}
	}

	logger.LogDebug("Server startup: initialize web sockets")
	ws := websockets.New()

	hotkeysWs := websockets.New()
	hotkeys.Ws = hotkeysWs

	notepadWs := websockets.New()
	globals.Notepad = notepad.New(notepadWs)

	report := &Report{}
	trafficReport := &TrafficReport{}
	teleportReport := &TeleportRequest{}

	eventSimStartID := simconnect.DWORD(0)
	startupTextEventID := simconnect.DWORD(0)

	logger.LogDebug("Server startup: initialize simconnect data definitions...")
	if s != nil {
		err = s.RegisterDataDefinition(report)
		if err != nil {
			utils.Println(err)
			logger.LogErrorVerboseOverride("s.RegisterDataDefinition(report) failed, reason: "+err.Error(), false)
			dialogs.ShowErrorAndExit("Communication with Flight Simulator failed!")
		}

		err = s.RegisterDataDefinition(trafficReport)
		if err != nil {
			utils.Println(err)
			logger.LogErrorVerboseOverride("s.RegisterDataDefinition(trafficReport) failed, reason: "+err.Error(), false)
			dialogs.ShowErrorAndExit("Communication with Flight Simulator failed!")
		}

		err = s.RegisterDataDefinition(teleportReport)
		if err != nil {
			utils.Println(err)
			logger.LogErrorVerboseOverride("s.RegisterDataDefinition(teleportReport) failed, reason: "+err.Error(), false)
			dialogs.ShowErrorAndExit("Communication with Flight Simulator failed!")
		}

		eventSimStartID = s.GetEventID()

		startupTextEventID = s.GetEventID()
		s.ShowText(simconnect.TEXT_TYPE_PRINT_WHITE, 15, startupTextEventID, ">> FSKneeboard connected <<")
	}

	go func() {
		logger.LogDebug("Server startup: initialize charts index...")
		charts.UpdateIndex()

		getContentType := func(requestedResource string) string {
			split := strings.Split(requestedResource, ".")
			extension := split[len(split)-1]

			switch extension {
			case "css":
				return "text/css"
			case "js":
				return "text/javascript"
			case "html":
				return "text/html"
			case "map":
				return "application/json"
			default:
				return "text/plain"
			}
		}

		setHeaders := func(contentType string, w http.ResponseWriter) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
			w.Header().Set("Content-Type", contentType)
		}

		sendResponse := func(w http.ResponseWriter, r *http.Request, filePath string, requestedResource string, asset []byte, assetErr error) {
			contentType := getContentType(requestedResource)
			logger.LogDebug("Resource [" + requestedResource + "] with MIME-type [" + contentType + "] requested!")

			setHeaders(contentType, w)

			if info, err := os.Stat(filePath); os.IsNotExist(err) {
				logger.LogDebug("Resource [" + requestedResource + "] not found in local file system! Serving embedded resource...")

				if assetErr != nil {
					logger.LogError("Embedded resource [" + requestedResource + "] not found!")
					http.Error(w, assetErr.Error(), http.StatusNotFound)
				} else {
					logger.LogDebug(fmt.Sprintf("Serving embedded resource [%s], %d Bytes total", requestedResource, len(asset)))
					w.Write(asset)
				}
			} else {
				logger.LogDebug(fmt.Sprintf("Serving resource [%s] from local file system, %d Bytes total", requestedResource, info.Size()))
				http.ServeFile(w, r, filePath)
			}
		}

		index := func(w http.ResponseWriter, r *http.Request) {
			logger.LogDebug("Request to index handler: " + r.URL.Path)
			requestedResource := strings.TrimPrefix(r.URL.Path, "/")
			if requestedResource == "" {
				requestedResource = "index.html"
			} else if requestedResource == "favicon.ico" {
				w.Write([]byte{})
				return
			}
			filePath := filepath.Join(filepath.Dir(exePath), "vfrmap", "html", "webdist", requestedResource)
			asset, assetErr := Asset(filepath.Base(filePath))
			sendResponse(w, r, filePath, requestedResource, asset, assetErr)
		}

		freemium := func(w http.ResponseWriter, r *http.Request) {
			logger.LogDebug("Request to freemium handler: " + r.URL.Path)

			requestedResource := strings.TrimPrefix(r.URL.Path, "/freemium/")
			filePath := filepath.Join(filepath.Dir(exePath), "vfrmap", "html", "freemium", "maps", "webdist", requestedResource)
			asset, assetErr := freemium.Asset(requestedResource)
			sendResponse(w, r, filePath, requestedResource, asset, assetErr)
		}

		premium := func(w http.ResponseWriter, r *http.Request) {
			logger.LogDebug("Request to premium handler: " + r.URL.Path)

			requestedResource := strings.TrimPrefix(r.URL.Path, "/premium/")
			filePath := filepath.Join(filepath.Dir(exePath), "_vendor", "premium", "webdist", requestedResource)
			asset, assetErr := premium.Asset(requestedResource)
			sendResponse(w, r, filePath, requestedResource, asset, assetErr)
		}

		chartsIndex := func(w http.ResponseWriter, r *http.Request) {
			setHeaders("application/json", w)
			charts.Json(w, r)
		}

		flightplan := func(w http.ResponseWriter, r *http.Request) {
			notifier, ok := w.(http.CloseNotifier)
			if !ok {
				utils.Println("Expected http.ResponseWriter to be an http.CloseNotifier")
				http.Error(w, "Expected http.ResponseWriter to be an http.CloseNotifier", http.StatusInternalServerError)
				return
			}

			ctx, cancel := context.WithCancel((context.Background()))
			ch := make(chan string)

			go waypoints.LocateCurrentFlightplan(s, w, r, ctx, ch)

			select {
			case <-ch:
				cancel()
				return
			case <-time.After(time.Second * 10):
				http.Error(w, "Server busy", http.StatusInternalServerError)
			case <-notifier.CloseNotify():
				utils.Println("Client has disconnected.")
			}
			cancel()
			<-ch
		}

		logger.LogDebug("Server startup: initialize map tile cache...")
		initMaptileCache()

		logger.LogDebug("Server startup: initialize chart server...")
		chartServer := http.FileServer(http.Dir("./charts"))

		logger.LogDebug("Server startup: initialize endpoints...")

		http.HandleFunc("/ws", ws.Serve)
		http.HandleFunc("/hotkeysWs", hotkeysWs.Serve)
		http.HandleFunc("/notepadWs", notepadWs.Serve)
		http.HandleFunc("/hotkey/", hotkeys.ServeMasterHotkey)
		http.HandleFunc("/mapserviceurls/", serveMapServiceUrls)
		http.HandleFunc("/tour/", tour.ServeTourStatus)
		http.HandleFunc("/log/", logger.LogController)
		http.HandleFunc("/loglevel/", logger.LogLevelController)
		http.HandleFunc("/data/", dbmanager.DataController)
		http.HandleFunc("/dataSet/", dbmanager.DataSetController)
		http.HandleFunc("/traildata/", trailDataController)
		http.HandleFunc("/freemium/", freemium)
		http.HandleFunc("/premium/", premium)
		http.HandleFunc("/premium/chartsIndex", chartsIndex)
		http.HandleFunc("/premium/flightplan", flightplan)
		http.Handle("/leafletjs/", http.StripPrefix("/leafletjs/", leafletjs.FS{}))
		http.Handle("/fontawesome/", http.StripPrefix("/fontawesome/", fontawesome.FS{}))
		http.Handle("/premium/charts/", http.StripPrefix("/premium/charts/", chartServer))
		http.HandleFunc("/", index)

		logger.LogDebug("Server startup: initialize endpoints done!")

		if globals.DevMode {
			logger.LogDebug("Server startup: initialize dev mode fileserver...")
			testServer := http.FileServer(http.Dir("../fskneeboard-panel/christian1984-ingamepanel-fskneeboard/html_ui/InGamePanels/FSKneeboardPanel"))
			http.Handle("/test/", http.StripPrefix("/test/", testServer))
		}

		logger.LogDebug("Server startup: starting server...")
		callbacks.UpdateServerStatus("Ready", "")
		logger.LogDebug("Server startup: updated GUI info to >Ready<!")

		go func() {
			err := http.ListenAndServe(globals.HttpListen, nil)
			if err != nil {
				logger.LogErrorVerbose("FSKneeboard Server could not be started! Reason: " + err.Error())
				callbacks.UpdateServerStatus("Could not start server...", "")
				dialogs.ShowErrorAndExit("FSKneeboard Server could not be started!\nPlease close ALL running instances of FSKneeboard an try again.")
			}
		}()

		// connect tablet etc.
		time.Sleep(2 * time.Second)
		logger.LogDebug("Server startup: obtaining public ip address...")

		ip, addr_err := utils.GetOutboundIP()
		server_addr_arr := strings.Split(globals.HttpListen, ":")
		port := server_addr_arr[len(server_addr_arr)-1]

		if addr_err != nil {
			logger.LogError("Server startup: Could not obtain public IP address... Reason: " + addr_err.Error())
		} else if ip != nil {
			logger.LogDebug("Server startup: processing and displaying public ip address...")

			if ip.To4() != nil {
				connectInfo := "http://" + ip.To4().String() + ":" + port
				logger.LogInfo("FSKneeboard available at: " + connectInfo)
				utils.Println("=== INFO: Connecting Your Tablet")
				utils.Println("Besides using the FSKneeboard ingame panel from within Flight Simulator you can also connect to FSKneeboard with your tablet or web browser. To do so please enter follwing IP address and port into the address bar.")
				utils.Println("FSKneeboard Server-Address: " + connectInfo)
				utils.Println("")

				callbacks.UpdateServerStatus("Ready at", connectInfo)
				logger.LogDebug("Server startup: processing and displaying public ip address done!")
			} else {
				logger.LogWarn("Server startup: could not obtain public ip address")

			}

		}
	}()

	logger.LogDebug("Server startup: initialize tickers")
	simconnectTick := time.NewTicker(100 * time.Millisecond)
	planePositionTick := time.NewTicker(200 * time.Millisecond)
	trafficPositionTick := time.NewTicker(10000 * time.Millisecond)

	loggedGetNextDispatchError := false

	for {
		select {
		case <-autosaveTick.C:
			if s == nil {
				continue
			}

			if globals.Pro && globals.AutosaveInterval > 0 {
				autosave.CreateAutosave(s, 5, false)
			}

		case <-planePositionTick.C:
			if globals.MockData {
				report.MockData()
				report.process(ws)
			} else if s != nil {
				report.RequestData(s)
			}

		case <-trafficPositionTick.C:
			if s == nil {
				continue
			}

		case <-simconnectTick.C:
			if s == nil {
				continue
			}

			ppData, r1, err := s.GetNextDispatch()
			if (err != nil) {
				logger.LogSilly(fmt.Sprintf("SimConnect error -> GetNextDispatch error: %d %s", r1, err))
			}

			if r1 < 0 {
				if uint32(r1) == simconnect.E_FAIL {
					// skip error, means no new messages?
					continue
				} else {
					if !loggedGetNextDispatchError {
						logger.LogError(fmt.Sprintf("GetNextDispatch error: %d %s", r1, err))
						loggedGetNextDispatchError = true
					}

					continue
				}
			}

			recvInfo := *(*simconnect.Recv)(ppData)

			switch recvInfo.ID {
			case simconnect.RECV_ID_EXCEPTION:
				recvErr := *(*simconnect.RecvException)(ppData)
				logger.LogError(fmt.Sprintf("SIMCONNECT_RECV_ID_EXCEPTION %#v\n", recvErr))

			case simconnect.RECV_ID_OPEN:
				recvOpen := *(*simconnect.RecvOpen)(ppData)
				fsInfo := fmt.Sprintf("\nFlight Simulator Info:\n  Codename: %s\n  Version: %d.%d (%d.%d)\n  Simconnect: %d.%d (%d.%d)",
					strings.Trim(string(recvOpen.ApplicationName[:]), "\x00"),
					recvOpen.ApplicationVersionMajor,
					recvOpen.ApplicationVersionMinor,
					recvOpen.ApplicationBuildMajor,
					recvOpen.ApplicationBuildMinor,
					recvOpen.SimConnectVersionMajor,
					recvOpen.SimConnectVersionMinor,
					recvOpen.SimConnectBuildMajor,
					recvOpen.SimConnectBuildMinor)
				logger.LogInfo("Connected to MSFS, details:"+fsInfo)
				utils.Println(fsInfo + "\n")
				utils.Printf("Ready... Please leave this window open during your Flight Simulator session. Have a safe flight :-)\n\n")

			case simconnect.RECV_ID_EVENT:
				recvEvent := *(*simconnect.RecvEvent)(ppData)

				switch recvEvent.EventID {
				case eventSimStartID:
					utils.Println("EVENT: SimStart")
				case startupTextEventID:
					// ignore
				default:
					utils.Println("unknown SIMCONNECT_RECV_ID_EVENT", recvEvent.EventID)
				}
			case simconnect.RECV_ID_WAYPOINT_LIST:
				waypointList := *(*simconnect.RecvFacilityWaypointList)(ppData)
				utils.Printf("SIMCONNECT_RECV_ID_WAYPOINT_LIST %#v\n", waypointList)

			case simconnect.RECV_ID_AIRPORT_LIST:
				airportList := *(*simconnect.RecvFacilityAirportList)(ppData)
				utils.Printf("SIMCONNECT_RECV_ID_AIRPORT_LIST %#v\n", airportList)

			case simconnect.RECV_ID_SIMOBJECT_DATA_BYTYPE:
				recvData := *(*simconnect.RecvSimobjectDataByType)(ppData)

				switch recvData.RequestID {
				case s.DefineMap["Report"]:
					report = (*Report)(ppData)

					/*
						if verbose {
							utils.Printf("REPORT: %#v\n", report)
						}
					*/

					report.process(ws)

				case s.DefineMap["TrafficReport"]:
					trafficReport = (*TrafficReport)(ppData)
					utils.Printf("TRAFFIC REPORT: %s\n", trafficReport.Inspect())
				}

			case simconnect.RECV_ID_SYSTEM_STATE:
				recvData := *(*simconnect.RecvSystemState)(ppData)

				filepathRaw := string(recvData.String[:])
				filepathReplace := strings.ReplaceAll(filepathRaw, string([]byte{0}), " ")
				filepath := strings.TrimSpace(filepathReplace)
				waypoints.SendFlightplanResponse(filepath)

			case simconnect.RECV_ID_QUIT:
				utils.Println("Flight Simulator was shut down. Exiting...")
				callbacks.UpdateServerStatus("Not Running", "")
				callbacks.UpdateMsfsConnectionStatus("Not Connected")

				ShutdownWithPrompt()

			default:
				utils.Println("recvInfo.ID unknown", recvInfo.ID)
			}

		case <-exitSignal:
			utils.Println("Exiting...")
			if s != nil {
				s.Close()
			}
			os.Exit(0)

		case _ = <-ws.NewConnection:
			// drain and skip

		case m := <-ws.ReceiveMessages:
			//utils.Println("ws.ReceiveMessages!")
			if s == nil {
				continue
			}
			handleClientMessage(m, s)
		}
	}
}

func handleClientMessage(m websockets.ReceiveMessage, s *simconnect.SimConnect) {
	var pkt map[string]interface{}
	if err := json.Unmarshal(m.Message, &pkt); err != nil {
		utils.Println("invalid websocket packet", err)
	} else {
		pktType, ok := pkt["type"].(string)
		if !ok {
			utils.Println("invalid websocket packet", pkt)
			return
		}
		switch pktType {
		case "teleport":
			if globals.DisableTeleport {
				utils.Println("teleport disabled", pkt)
				return
			}

			// validate user input
			lat, ok := pkt["lat"].(float64)
			if !ok {
				utils.Println("invalid websocket packet", pkt)
				return
			}
			lng, ok := pkt["lng"].(float64)
			if !ok {
				utils.Println("invalid websocket packet", pkt)
				return
			}
			altitude, ok := pkt["altitude"].(float64)
			if !ok {
				utils.Println("invalid websocket packet", pkt)
				return
			}

			// teleport
			r := &TeleportRequest{Latitude: lat, Longitude: lng, Altitude: altitude}
			r.SetData(s)
		}
	}
}
