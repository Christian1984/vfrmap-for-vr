package server

//go:generate go-bindata -pkg server -o bindata.go -modtime 1 -prefix ../html ../html

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"
	"vfrmap-for-vr/freemium_src/autosave"
	"vfrmap-for-vr/freemium_src/charts"
	"vfrmap-for-vr/freemium_src/notepad"
	"vfrmap-for-vr/freemium_src/waypoints"
	"vfrmap-for-vr/simconnect"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/callbacks"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/html/fontawesome"
	"vfrmap-for-vr/vfrmap/html/freemium"
	"vfrmap-for-vr/vfrmap/html/leafletjs"
	"vfrmap-for-vr/vfrmap/html/premium"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
	"vfrmap-for-vr/vfrmap/websockets"
)

var started = false

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

func (r *Report) RequestData(s *simconnect.SimConnect) {
	defineID := s.GetDefineID(r)
	requestID := defineID
	s.RequestDataOnSimObjectType(requestID, defineID, 0, simconnect.SIMOBJECT_TYPE_USER)
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

type Hotkey struct {
	AltKey   bool `json:"altkey"`
	CtrlKey  bool `json:"ctrlkey"`
	ShiftKey bool `json:"shiftkey"`
	KeyCode  int  `json:"keycode"`
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

func ShutdownWithPrompt() {
	if !globals.Quietshutdown {
		buf := bufio.NewReader(os.Stdin)
		utils.Print("\nPress ENTER to continue...") //TODO
		buf.ReadBytes('\n')
	}

	os.Exit(0)
}

func StartFskServer() {
	callbacks.UpdateServerStarted(true)

	if globals.Pro && !globals.DrmValid {
		dialogs.ShowLicenseError()
		utils.Println("WARNING: Cannot start FSKneeboard server, reason: no valid license found!")
		logger.LogWarn("Cannot start server, reason: no valid license found", false)
		callbacks.UpdateServerStarted(false)
		return
	}

	if started {
		return
	}

	started = true

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)
	
	time.Sleep(1 * time.Second)

	exePath, _ := os.Executable()

	// wait for Flight Simulator
	var s *simconnect.SimConnect
	var err error

	logger.LogInfo("Waiting for Flight Simulator...", false)
	utils.Println("")
	utils.Print("Waiting for Flight Simulator..")
	callbacks.UpdateMsfsConnectionStatus("Connecting...")

	for true {
		utils.Print(".")

		s, err = simconnect.New(globals.ProductName)

		if err != nil {

			if globals.DevMode {
				logger.LogInfo("Running with --dev: Not connected to Flight Simulator!", false)
				utils.Println("")
				utils.Println("Running with --dev: Not connected to Flight Simulator!!!")
				utils.Println("")

				callbacks.UpdateMsfsConnectionStatus("Not Connected (dev Mode)")

				break
			}

			time.Sleep(5 * time.Second)
		} else if s != nil {
			logger.LogInfo("Connected to Flight Simulator!", false)
			utils.Println("")
			utils.Println("Connected to Flight Simulator!")
			utils.Println("")
			callbacks.UpdateMsfsConnectionStatus("Connected")
			break
		}
	}

	ws := websockets.New()
	notepadWs := websockets.New()
	globals.Notepad = notepad.New(notepadWs, globals.Verbose)

	report := &Report{}
	trafficReport := &TrafficReport{}
	teleportReport := &TeleportRequest{}

	eventSimStartID := simconnect.DWORD(0)
	startupTextEventID := simconnect.DWORD(0)

	if s != nil {
		err = s.RegisterDataDefinition(report)
		if err != nil {
			utils.Println(err)
			panic(err)
		}

		err = s.RegisterDataDefinition(trafficReport)
		if err != nil {
			panic(err)
		}

		err = s.RegisterDataDefinition(teleportReport)
		if err != nil {
			panic(err)
		}

		eventSimStartID = s.GetEventID()

		startupTextEventID = s.GetEventID()
		s.ShowText(simconnect.TEXT_TYPE_PRINT_WHITE, 15, startupTextEventID, ">> FSKneeboard connected <<")
	}

	go func() {
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

		sendResponse := func(w http.ResponseWriter, r *http.Request, filePath string, requestedResource string, asset []byte) {
			contentType := getContentType(requestedResource)
			setHeaders(contentType, w)

			if _, err = os.Stat(filePath); os.IsNotExist(err) {
				w.Write(asset)
			} else {
				http.ServeFile(w, r, filePath)
			}
		}

		index := func(w http.ResponseWriter, r *http.Request) {
			requestedResource := strings.TrimPrefix(r.URL.Path, "/")
			if requestedResource == "" {
				requestedResource = "index.html"
			} else if requestedResource == "favicon.ico" {
				w.Write([]byte{})
				return
			}
			filePath := filepath.Join(filepath.Dir(exePath), "vfrmap", "html", requestedResource)
			sendResponse(w, r, filePath, requestedResource, MustAsset(filepath.Base(filePath)))
		}

		hotkey := func(w http.ResponseWriter, r *http.Request) {
			keycode := -1
			altkey := false
			shiftkey := false
			ctrlkey := false

			switch globals.Hotkey {
			case 1:
				altkey = true
				keycode = 70
			case 2:
				altkey = true
				keycode = 75
			case 3:
				altkey = true
				keycode = 84
			case 4:
				shiftkey = true
				ctrlkey = true
				keycode = 70
			case 5:
				shiftkey = true
				ctrlkey = true
				keycode = 75
			case 6:
				shiftkey = true
				ctrlkey = true
				keycode = 84
			}

			hotkey := Hotkey{altkey, shiftkey, ctrlkey, keycode}
			responseJson, jsonErr := json.Marshal(hotkey)

			if jsonErr != nil {
				utils.Println(jsonErr.Error())
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

		freemium := func(w http.ResponseWriter, r *http.Request) {
			requestedResource := strings.TrimPrefix(r.URL.Path, "/freemium/")
			filePath := filepath.Join(filepath.Dir(exePath), "vfrmap", "html", "freemium", "maps", requestedResource)
			sendResponse(w, r, filePath, requestedResource, freemium.MustAsset(requestedResource))
		}

		premium := func(w http.ResponseWriter, r *http.Request) {
			requestedResource := strings.TrimPrefix(r.URL.Path, "/premium/")
			filePath := filepath.Join(filepath.Dir(exePath), "_vendor", "premium", requestedResource)
			sendResponse(w, r, filePath, requestedResource, premium.MustAsset(requestedResource))
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

		chartServer := http.FileServer(http.Dir("./charts"))

		http.HandleFunc("/ws", ws.Serve)
		http.HandleFunc("/notepadWs", notepadWs.Serve)
		http.HandleFunc("/hotkey/", hotkey)
		http.HandleFunc("/log/", logger.LogController)
		http.HandleFunc("/loglevel/", logger.LogLevelController)
		http.HandleFunc("/data/", dbmanager.DataController)
		http.HandleFunc("/dataSet/", dbmanager.DataSetController)
		http.HandleFunc("/freemium/", freemium)
		http.HandleFunc("/premium/", premium)
		http.HandleFunc("/premium/chartsIndex", chartsIndex)
		http.HandleFunc("/premium/flightplan", flightplan)
		http.Handle("/leafletjs/", http.StripPrefix("/leafletjs/", leafletjs.FS{}))
		http.Handle("/fontawesome/", http.StripPrefix("/fontawesome/", fontawesome.FS{}))
		http.Handle("/premium/charts/", http.StripPrefix("/premium/charts/", chartServer))
		http.HandleFunc("/", index)

		if globals.DevMode {
			testServer := http.FileServer(http.Dir("../fskneeboard-panel/christian1984-ingamepanel-fskneeboard/html_ui/InGamePanels/FSKneeboardPanel"))
			http.Handle("/test/", http.StripPrefix("/test/", testServer))
		}

		// connect tablet etc.
		ip, addr_err := utils.GetOutboundIP()
		server_addr_arr := strings.Split(globals.HttpListen, ":")
		port := server_addr_arr[len(server_addr_arr) - 1]

		if (addr_err == nil && ip != nil) {
			connectInfo := ip.To4().String() + ":" + port
			logger.LogInfo("FSKneeboard available at: " + connectInfo, false)
			utils.Println("=== INFO: Connecting Your Tablet")
			utils.Println("Besides using the FSKneeboard ingame panel from within Flight Simulator you can also connect to FSKneeboard with your tablet or web browser. To do so please enter follwing IP address and port into the address bar.")
			utils.Println("FSKneeboard Server-Address: " + connectInfo)
			utils.Println("")

			callbacks.UpdateServerStatus("Ready at " + connectInfo)
		}

		err := http.ListenAndServe(globals.HttpListen, nil)
		if err != nil {
			callbacks.UpdateServerStarted(false)
			panic(err)
		}
	}()

	var autosaveTick *time.Ticker

	if globals.AutosaveInterval > 0 {
		autosaveTick = time.NewTicker(time.Duration(globals.AutosaveInterval) * time.Minute)
	} else {
		autosaveTick = time.NewTicker(9999 * time.Minute)
	}

	simconnectTick := time.NewTicker(100 * time.Millisecond)
	planePositionTick := time.NewTicker(200 * time.Millisecond)
	trafficPositionTick := time.NewTicker(10000 * time.Millisecond)

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
			if s == nil {
				continue
			}

			report.RequestData(s)

		case <-trafficPositionTick.C:
			if s == nil {
				continue
			}

		case <-simconnectTick.C:
			if s == nil {
				continue
			}

			ppData, r1, err := s.GetNextDispatch()

			if r1 < 0 {
				if uint32(r1) == simconnect.E_FAIL {
					// skip error, means no new messages?
					continue
				} else {
					panic(fmt.Errorf("GetNextDispatch error: %d %s", r1, err))
				}
			}

			recvInfo := *(*simconnect.Recv)(ppData)

			switch recvInfo.ID {
			case simconnect.RECV_ID_EXCEPTION:
				recvErr := *(*simconnect.RecvException)(ppData)
				utils.Printf("SIMCONNECT_RECV_ID_EXCEPTION %#v\n", recvErr)

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
				logger.LogInfo("Connected to MSFS, details:" + fsInfo, false)
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

					ws.Broadcast(map[string]interface{}{
						"type":           "plane",
						"latitude":       report.Latitude,
						"longitude":      report.Longitude,
						"altitude":       fmt.Sprintf("%.0f", report.Altitude),
						"heading":        int(report.Heading),
						"airspeed":       fmt.Sprintf("%.0f", report.Airspeed),
						"airspeed_true":  fmt.Sprintf("%.0f", report.AirspeedTrue),
						"vertical_speed": fmt.Sprintf("%.0f", report.VerticalSpeed),
						"flaps":          fmt.Sprintf("%.0f", report.Flaps),
						"trim":           fmt.Sprintf("%.1f", report.Trim),
						"rudder_trim":    fmt.Sprintf("%.1f", report.RudderTrim),
						"wind_direction": fmt.Sprintf("%.0f", report.WindDirection),
						"wind_velocity":  fmt.Sprintf("%.0f", report.WindVelocity),
					})

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
			utils.Println("ws.ReceiveMessages!")
			if s == nil {
				continue
			}
			handleClientMessage(m, s)
		}
	}
}

func StopFskServer() {
	//TODO
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