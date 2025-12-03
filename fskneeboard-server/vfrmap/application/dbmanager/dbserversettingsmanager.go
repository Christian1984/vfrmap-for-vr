package dbmanager

import (
	"fmt"
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
)

// msfs version
func StoreMsfsVersion() {
	DbWriteSettings("msfsVersion", globals.MsfsVersion)
}

func LoadMsfsVersion() {
	version := DbReadSettings("msfsVersion")

	if version == "" {
		// Default to MSFS 2020 Windows Store
		version = "2020-winstore"
	}

	// Validate the version string
	switch version {
	case "2020-steam", "2020-winstore", "2024-steam", "2024-winstore":
		globals.MsfsVersion = version
	default:
		globals.MsfsVersion = "2020-winstore" // fallback to default
	}
} // msfs autostart
func StoreMsfsAutostart() {
	DbWriteSettings("msfsAutostart", strconv.FormatBool(globals.MsfsAutostart))
}

func LoadMsfsAutostart() {
	autostart, _ := strconv.ParseBool(DbReadSettings("msfsAutostart"))
	globals.MsfsAutostart = autostart
}

// tour
func StoreTourStates() {
	DbWriteSettings("tourIndexStarted", strconv.FormatBool(globals.TourIndexStarted))
	DbWriteSettings("tourMapStarted", strconv.FormatBool(globals.TourMapStarted))
	DbWriteSettings("tourChartsStarted", strconv.FormatBool(globals.TourChartsStarted))
	DbWriteSettings("tourNotepadStarted", strconv.FormatBool(globals.TourNotepadStarted))
	DbWriteSettings("tourGuiStarted", strconv.FormatBool(globals.TourGuiStarted))
}

func LoadTourStates() {
	tourIndexStarted, _ := strconv.ParseBool(DbReadSettings("tourIndexStarted"))
	globals.TourIndexStarted = tourIndexStarted

	tourMapStarted, _ := strconv.ParseBool(DbReadSettings("tourMapStarted"))
	globals.TourMapStarted = tourMapStarted

	tourNotepadStarted, _ := strconv.ParseBool(DbReadSettings("tourNotepadStarted"))
	globals.TourNotepadStarted = tourNotepadStarted

	tourChartsStarted, _ := strconv.ParseBool(DbReadSettings("tourChartsStarted"))
	globals.TourChartsStarted = tourChartsStarted

	tourGuiStarted, _ := strconv.ParseBool(DbReadSettings("tourGuiStarted"))
	globals.TourGuiStarted = tourGuiStarted
}

// autosave
func StoreAutosaveInterval() {
	DbWriteSettings("autosaveInterval", strconv.Itoa(globals.AutosaveInterval))
}

func LoadAutosaveInterval() {
	interval, _ := strconv.Atoi(DbReadSettings("autosaveInterval"))
	globals.AutosaveInterval = interval
}

// loglevel
func StoreLogLevel() {
	DbWriteSettings("loglevel", globals.LogLevel)
}

func LoadLogLevel() string {
	res := DbReadSettings("loglevel")

	if res == "" {
		res = "off"
	}

	return res
}

// interface scale
func StoreInterfaceScale() {
	interfaceScaleString := fmt.Sprintf("%f", globals.InterfaceScale)
	DbWriteSettings("interfacescale", interfaceScaleString)
}

func LoadInterfaceScale() {
	res := DbReadSettings("interfacescale")

	floatValue, err := strconv.ParseFloat(res, 64)

	if res == "" || err != nil || floatValue < 0 {
		floatValue = 1
	}

	globals.InterfaceScale = floatValue
}

func StoreInterfaceScalePromptShown() {
	DbWriteSettings("interfacescaleShowPrompt", strconv.FormatBool(globals.InterfaceScalePromptShown))
}

func LoadInterfaceScalePromptShown() {
	interfacescaleShowPrompt, _ := strconv.ParseBool(DbReadSettings("interfacescaleShowPrompt"))
	globals.InterfaceScalePromptShown = interfacescaleShowPrompt
}

// openAPI
func StoreOpenAipApiKey() {
	DbWriteSettings("oaipApiKey", strings.TrimSpace(globals.OpenAipApiKey))
}

func LoadOpenAipApiKey() {
	openAipApiKey := strings.TrimSpace(DbReadSettings("oaipApiKey"))
	globals.OpenAipApiKey = openAipApiKey
}

func StoreOpenAipBypassCache() {
	DbWriteSettings("oaipBypassCache", strconv.FormatBool(globals.OpenAipBypassCache))
}

func LoadOpenAipBypassCache() {
	openAipBypassCache, _ := strconv.ParseBool(DbReadSettings("oaipBypassCache"))
	globals.OpenAipBypassCache = openAipBypassCache
}

// bing
func StoreBingMapsApiKey() {
	DbWriteSettings("bingMapsApiKey", strings.TrimSpace(globals.BingMapsApiKey))
}

func LoadBingMapsApiKey() {
	if !globals.Pro {
		return
	}

	bingMapsApiKey := strings.TrimSpace(DbReadSettings("bingMapsApiKey"))
	globals.BingMapsApiKey = bingMapsApiKey
}

// google maps
func StoreGoogleMapsApiKey() {
	DbWriteSettings("googleMapsApiKey", strings.TrimSpace(globals.GoogleMapsApiKey))
}

func LoadGoogleMapsApiKey() {
	googleMapsApiKey := strings.TrimSpace(DbReadSettings("googleMapsApiKey"))
	globals.GoogleMapsApiKey = googleMapsApiKey
}
