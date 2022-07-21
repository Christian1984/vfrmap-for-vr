package dbmanager

import (
	"strconv"
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
