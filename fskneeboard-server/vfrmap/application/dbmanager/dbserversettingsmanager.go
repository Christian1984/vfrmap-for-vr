package dbmanager

import (
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
)

// msfs version
func StoreMsfsVersion() {
	if globals.WinstoreFs && !globals.SteamFs {
		DbWriteSettings("msfsVersion", "winstore")
	} else if !globals.WinstoreFs && globals.SteamFs {
		DbWriteSettings("msfsVersion", "steam")
	}
}

func LoadMsfsVersion() {
	res := DbReadSettings("msfsVersion")

	if res == "steam" {
		globals.SteamFs = true
		globals.WinstoreFs = false
	} else {
		globals.WinstoreFs = true
		globals.SteamFs = false
	}
}

// msfs autostart
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