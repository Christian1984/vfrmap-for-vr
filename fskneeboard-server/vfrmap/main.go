package main

//go:generate go-bindata -pkg main -o bindata.go -modtime 1 -prefix html html

// build: GOOS=windows GOARCH=amd64 go build -o fskneeboard.exe vfrmap-for-vr/vfrmap

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"vfrmap-for-vr/_vendor/premium/autosave"
	"vfrmap-for-vr/_vendor/premium/charts"
	"vfrmap-for-vr/_vendor/premium/common"
	"vfrmap-for-vr/_vendor/premium/drm"
	"vfrmap-for-vr/simconnect"
	"vfrmap-for-vr/vfrmap/html/fontawesome"
	"vfrmap-for-vr/vfrmap/html/freemium"
	"vfrmap-for-vr/vfrmap/html/leafletjs"
	"vfrmap-for-vr/vfrmap/html/premium"
	"vfrmap-for-vr/vfrmap/websockets"

	updatechecker "github.com/Christian1984/go-update-checker"
)

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

func shutdownWithPromt() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("\nPress ENTER to continue...")
	buf.ReadBytes('\n')

	os.Exit(0)
}

var buildVersion string
var buildTime string
var pro string

var bPro bool
var productName string

var disableTeleport bool
var devMode bool
var steamfs bool
var winstorefs bool
var noupdatecheck bool

var autosaveInterval int

var verbose bool
var httpListen string

func main() {
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.StringVar(&httpListen, "listen", "0.0.0.0:9000", "http listen")
	flag.BoolVar(&disableTeleport, "disable-teleport", false, "disable teleport")
	flag.BoolVar(&devMode, "dev", false, "enable dev mode, i.e. no running msfs required")
	flag.BoolVar(&steamfs, "steamfs", false, "start Flight Simulator via Steam")
	flag.BoolVar(&winstorefs, "winstorefs", false, "start Flight Simulator via Windows Store")
	flag.BoolVar(&noupdatecheck, "noupdatecheck", false, "prevent FSKneeboard from checking the GitHub API for updates")
	flag.IntVar(&autosaveInterval, "autosave", 0, "set autosave interval in minutes")
	flag.Parse()

	bPro = pro == "true"

	productName = "FSKneeboard"
	if bPro {
		productName += " PRO"
	}

	fmt.Printf("\n"+productName+" - Server\n  Website: https://fskneeboard.com\n  Readme:  https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md\n  Issues:  https://github.com/Christian1984/vfrmap-for-vr/issues\n  Version: %s (%s)\n\n", buildVersion, buildTime)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)
	exePath, _ := os.Executable()

	if bPro {
		fmt.Println("=== INFO: License")
		if !drm.Valid() {
			fmt.Println("\nWARNING: You do not have a valid license to run FSKneeboard PRO!")
			fmt.Println("Please purchase a license at https://fskneeboard.com/buy-now and place your fskneeboard.lic-file in the same directory as fskneeboard.exe.")
			shutdownWithPromt()
		} else {
			fmt.Println("Valid license found!")
			fmt.Println("Thanks for purchasing FSKneeboard PRO and supporting the development of this mod!")
			fmt.Println("")
		}
	} else {
		fmt.Println("=== INFO: How to Support the Development of FSKneeboard")
		fmt.Println("Thanks for trying FSKneeboard FREE!")
		fmt.Println("Please checkout https://fskneeboard.com and purchase FSKneeboard PRO to unlock all features the extension has to offer.")
		fmt.Println("")
	}

	if !noupdatecheck {
		uc := updatechecker.New("Christian1984", "vfrmap-for-vr", "FSKneeboard", common.DOWNLOAD_LINK, 3, false)
		uc.CheckForUpdate(buildVersion)

		if uc.UpdateAvailable {
			uc.PrintMessage()
			fmt.Println("")
		}
	}

	// autosave info
	fmt.Println("=== INFO: Autosave")

	if autosaveInterval > 0 {
		fmt.Printf("Autosave Interval set to %d minute(s)...\n", autosaveInterval)
	} else {
		fmt.Println("INFO: Autosave not activated. Run fskneeboard.exe --autosave 5 to automatically save your flights every 5 minutes...")
	}

	if !bPro {
		fmt.Println("PLEASE NOTE: 'Autosave' is a feature available exclusively to FSKneeboard PRO supporters. Please consider supporting the development of FSKneeboard by purchasing a license at https://fskneeboard.com/buy-now/")
	}

	fmt.Println("")

	// starting Flight Simulator
	fmt.Println("=== INFO: Flight Simulator Autostart")

	if steamfs {
		fmt.Println("Starting Flight Simulator via Steam... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start steam://run/1250410")
		fserr := cmd.Start()
		if fserr != nil {
			fmt.Println("Flight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
		}
	} else if winstorefs {
		fmt.Println("Starting Flight Simulator... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start shell:AppsFolder\\Microsoft.FlightSimulator_8wekyb3d8bbwe!App -FastLaunch")
		fserr := cmd.Run()
		if fserr != nil {
			fmt.Println("WARNING: Flight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			fmt.Println("IMPORTANT: If you have purchased MSFS on Steam, please run 'fskneeboard.exe --steamfs' as described in the manual under 'Usage'!")
		}
	} else {
		fmt.Println("FSKneeboard started without autostart options --steamfs or --winstorefs.")
		fmt.Println("If you haven't already, please start Flight Simulator manually!")
	}

	// wait for Flight Simulator
	var s *simconnect.SimConnect
	var err error

	fmt.Print("\nConnecting to Flight Simulator..")

	for true {
		fmt.Print(".")

		s, err = simconnect.New(productName)

		if err != nil {

			if devMode {
				fmt.Println("\nRunning with --dev: Not connected to Flight Simulator!!!")
				break
			}

			time.Sleep(5 * time.Second)
		} else if s != nil {
			fmt.Println("\nConnected to Flight Simulator!")
			break
		}
	}

	ws := websockets.New()

	report := &Report{}
	trafficReport := &TrafficReport{}
	teleportReport := &TeleportRequest{}

	eventSimStartID := simconnect.DWORD(0)
	startupTextEventID := simconnect.DWORD(0)

	if s != nil {
		err = s.RegisterDataDefinition(report)
		if err != nil {
			fmt.Println(err)
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

		chartServer := http.FileServer(http.Dir("./charts"))

		http.HandleFunc("/ws", ws.Serve)
		http.HandleFunc("/freemium/", freemium)
		http.HandleFunc("/premium/", premium)
		http.HandleFunc("/premium/chartsIndex", chartsIndex)
		http.Handle("/leafletjs/", http.StripPrefix("/leafletjs/", leafletjs.FS{}))
		http.Handle("/fontawesome/", http.StripPrefix("/fontawesome/", fontawesome.FS{}))
		http.Handle("/premium/charts/", http.StripPrefix("/premium/charts/", chartServer))
		http.HandleFunc("/", index)

		if devMode {
			testServer := http.FileServer(http.Dir("../fskneeboard-panel/christian1984-ingamepanel-fskneeboard/html_ui/InGamePanels/CustomPanel"))
			http.Handle("/test/", http.StripPrefix("/test/", testServer))
		}

		err := http.ListenAndServe(httpListen, nil)
		if err != nil {
			panic(err)
		}
	}()

	var autosaveTick *time.Ticker

	if autosaveInterval > 0 {
		autosaveTick = time.NewTicker(time.Duration(autosaveInterval) * time.Minute)
	} else {
		autosaveTick = time.NewTicker(9999 * time.Minute)
	}

	simconnectTick := time.NewTicker(100 * time.Millisecond)
	planePositionTick := time.NewTicker(200 * time.Millisecond)
	trafficPositionTick := time.NewTicker(10000 * time.Millisecond)
	systemStateTick := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-systemStateTick.C:
			if s == nil {
				continue
			}

			fmt.Println("Sending RequestSystemState...")
			s.RequestSystemState(1377, "FlightPlan")

		case <-autosaveTick.C:
			if s == nil {
				continue
			}

			if bPro && autosaveInterval > 0 {
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
				fmt.Printf("SIMCONNECT_RECV_ID_EXCEPTION %#v\n", recvErr)

			case simconnect.RECV_ID_OPEN:
				recvOpen := *(*simconnect.RecvOpen)(ppData)
				fmt.Printf(
					"\nFlight Simulator Info:\n  Codename: %s\n  Version: %d.%d (%d.%d)\n  Simconnect: %d.%d (%d.%d)\n\n",
					recvOpen.ApplicationName,
					recvOpen.ApplicationVersionMajor,
					recvOpen.ApplicationVersionMinor,
					recvOpen.ApplicationBuildMajor,
					recvOpen.ApplicationBuildMinor,
					recvOpen.SimConnectVersionMajor,
					recvOpen.SimConnectVersionMinor,
					recvOpen.SimConnectBuildMajor,
					recvOpen.SimConnectBuildMinor,
				)
				fmt.Printf("Ready... Please leave this window open during your Flight Simulator session. Have a safe flight :-)\n\n")

			case simconnect.RECV_ID_EVENT:
				recvEvent := *(*simconnect.RecvEvent)(ppData)

				switch recvEvent.EventID {
				case eventSimStartID:
					fmt.Println("EVENT: SimStart")
				case startupTextEventID:
					// ignore
				default:
					fmt.Println("unknown SIMCONNECT_RECV_ID_EVENT", recvEvent.EventID)
				}
			case simconnect.RECV_ID_WAYPOINT_LIST:
				waypointList := *(*simconnect.RecvFacilityWaypointList)(ppData)
				fmt.Printf("SIMCONNECT_RECV_ID_WAYPOINT_LIST %#v\n", waypointList)

			case simconnect.RECV_ID_AIRPORT_LIST:
				airportList := *(*simconnect.RecvFacilityAirportList)(ppData)
				fmt.Printf("SIMCONNECT_RECV_ID_AIRPORT_LIST %#v\n", airportList)

			case simconnect.RECV_ID_SIMOBJECT_DATA_BYTYPE:
				recvData := *(*simconnect.RecvSimobjectDataByType)(ppData)

				switch recvData.RequestID {
				case s.DefineMap["Report"]:
					report = (*Report)(ppData)

					if verbose {
						fmt.Printf("REPORT: %#v\n", report)
					}

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
					})

				case s.DefineMap["TrafficReport"]:
					trafficReport = (*TrafficReport)(ppData)
					fmt.Printf("TRAFFIC REPORT: %s\n", trafficReport.Inspect())
				}

			case simconnect.RECV_ID_SYSTEM_STATE:
				recvData := *(*simconnect.RecvSystemState)(ppData)
				fmt.Println("Received System State...")
				fmt.Println(string(recvData.String[:]))

			case simconnect.RECV_ID_QUIT:
				fmt.Println("Flight Simulator was shut down. Exiting...")
				shutdownWithPromt()

			default:
				fmt.Println("recvInfo.ID unknown", recvInfo.ID)
			}

		case <-exitSignal:
			fmt.Println("Exiting...")
			if s != nil {
				s.Close()
			}
			os.Exit(0)

		case _ = <-ws.NewConnection:
			// drain and skip

		case m := <-ws.ReceiveMessages:
			fmt.Println("ws.ReceiveMessages!")
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
		fmt.Println("invalid websocket packet", err)
	} else {
		pktType, ok := pkt["type"].(string)
		if !ok {
			fmt.Println("invalid websocket packet", pkt)
			return
		}
		switch pktType {
		case "teleport":
			if disableTeleport {
				fmt.Println("teleport disabled", pkt)
				return
			}

			// validate user input
			lat, ok := pkt["lat"].(float64)
			if !ok {
				fmt.Println("invalid websocket packet", pkt)
				return
			}
			lng, ok := pkt["lng"].(float64)
			if !ok {
				fmt.Println("invalid websocket packet", pkt)
				return
			}
			altitude, ok := pkt["altitude"].(float64)
			if !ok {
				fmt.Println("invalid websocket packet", pkt)
				return
			}

			// teleport
			r := &TeleportRequest{Latitude: lat, Longitude: lng, Altitude: altitude}
			r.SetData(s)
		}
	}
}
