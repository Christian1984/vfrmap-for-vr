package settingspanel

import (
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const msfsVersionOptionWinstore = "Windows Store Version"
const msfsVersionOptionSteam = "Steam Version"

var msfsVersionOptions = []string{msfsVersionOptionWinstore, msfsVersionOptionSteam}
var msfsVersionBinding = binding.NewString()

var msfsAutostartBinding = binding.NewBool()
var serverAutostartBinding = binding.NewBool()

var autosaveOptions = []string{"Off", "1", "5", "10", "15", "30", "60"}
var autosaveBinding = binding.NewString()

var loglevelOptions = []string{
	strings.Title(logger.Off),
	strings.Title(logger.Debug),
	strings.Title(logger.Info),
	strings.Title(logger.Warn),
	strings.Title(logger.Error),
}
var loglevelBinding = binding.NewString()

func UpdateAutosaveStatus(interval int) {
	intervalString := "Off"

	if interval > 0 {
		intervalString = strconv.Itoa(interval)
	}

	autosaveBinding.Set(intervalString)
}

func UpdateMsfsVersionStatus(steam bool) {
	if steam {
		msfsVersionBinding.Set(msfsVersionOptionSteam)
	} else {
		msfsVersionBinding.Set(msfsVersionOptionWinstore)
	}
}

func UpdateMsfsAutostartStatus(autostart bool) {
	msfsAutostartBinding.Set(autostart)
}

func UpdateServerAutostartStatus(autostart bool) {
	serverAutostartBinding.Set(autostart)
}

func UpdateLogLevelStatus(level string) {
	lowerLevel := strings.ToLower(level)

	if lowerLevel != logger.Debug && lowerLevel != logger.Info && lowerLevel != logger.Warn && lowerLevel != logger.Error {
		lowerLevel = "off"
	}

	loglevelBinding.Set(lowerLevel)
}

func SettingsPanel() *fyne.Container {
	logger.LogDebug("Initializing Settings Panel...", false)

	// MSFS version select
	msfsVersionLabel := widget.NewLabel("Flight Simulator Version")
	msfsVersionSelect := widget.NewSelect(msfsVersionOptions, func(selected string) {
		msfsVersionBinding.Set(selected)
	})

	msfsVersionBinding.AddListener(binding.NewDataListener(func() {
		selected, _ := msfsVersionBinding.Get()

		msfsVersionSelect.SetSelected(selected)

		globals.SteamFs = false
		globals.WinstoreFs = false

		if selected == msfsVersionOptionWinstore {
			globals.WinstoreFs = true
			logger.LogInfo("Selected MSFS Version: Windows Store", false)
		} else if selected == msfsVersionOptionSteam {
			globals.SteamFs = true
			logger.LogInfo("Selected MSFS Version: Steam", false)
		}

		dbmanager.StoreMsfsVersion()
	}))

	msfsVersionBinding.Set(msfsVersionOptionWinstore)

	// msfs autostart select
	msfsAutostartLabel := widget.NewLabel("Flight Simulator Autostart")
	msfsAutostartCb := widget.NewCheckWithData("Start MSFS when FSKneeboard starts", msfsAutostartBinding)

	msfsAutostartBinding.AddListener(binding.NewDataListener(func() {
		msfsAutostart, _ := msfsAutostartBinding.Get()
		globals.MsfsAutostart = msfsAutostart
		logger.LogInfo("MSFS Autostart updated: " + strconv.FormatBool(msfsAutostart), false)

		dbmanager.StoreMsfsAutostart()
	}))

	// server autostart select
	serverAutostartLabel := widget.NewLabel("FSKneeboard Server Autostart")
	serverAutostartCb := widget.NewCheckWithData("Start Server when FSKneeboard starts", serverAutostartBinding)

	serverAutostartBinding.AddListener(binding.NewDataListener(func() {
		serverAutostart, _ := serverAutostartBinding.Get()
		globals.ServerAutostart = serverAutostart
		logger.LogInfo("Server Autostart updated: " + strconv.FormatBool(serverAutostart), false)

		dbmanager.StoreServerAutostart()
	}))

	// set autosave properties
	autosaveLabel := widget.NewLabel("Autosave Interval [minutes]")
	autosaveSelect := widget.NewSelect(autosaveOptions, func(selected string) {
		autosaveBinding.Set(selected)
	})

	autosaveBinding.AddListener(binding.NewDataListener(func() {
		autosaveString, _ := autosaveBinding.Get()

		if autosaveString != "Off" && !globals.Pro {
			autosaveString = "Off"
			dialogs.ShowProFeatureInfo("Autosave")
		}

		autosaveSelect.SetSelected(autosaveString)

		autosaveInterval, err := strconv.Atoi(autosaveString)
		if err != nil {
			autosaveInterval = 0
		}

		globals.AutosaveInterval = autosaveInterval
		dbmanager.StoreAutosaveInterval()

		server.UpdateAutosaveInterval()
	}))

	autosaveBinding.Set("Off")

	// set log level
	loglevelLabel := widget.NewLabel("Log Level")
	loglevelSelect := widget.NewSelect(loglevelOptions, func(selected string) {
		loglevelBinding.Set(selected)
	})

	loglevelBinding.AddListener(binding.NewDataListener(func() {
		loglevelString, _ := loglevelBinding.Get()

		matchIndex := 0

		for index, value := range loglevelOptions {
			if strings.ToLower(loglevelString) == strings.ToLower(value) {
				matchIndex = index
				break
			}
		}

		loglevelSelect.SetSelected(loglevelOptions[matchIndex])

		globals.LogLevel = strings.ToLower(loglevelString)

		dbmanager.StoreLogLevel()

		logger.SetLevel(loglevelString)
		logger.TryCreateLogFile()
	}))

	loglevelBinding.Set(globals.LogLevel)

	// grid and centerContainer
	//empty := widget.NewLabel("")
	grid := container.NewGridWithColumns(
		2,
		msfsVersionLabel, msfsVersionSelect,
		msfsAutostartLabel, msfsAutostartCb,
		serverAutostartLabel, serverAutostartCb,
		autosaveLabel, autosaveSelect,
		loglevelLabel, loglevelSelect,
	)
	centerContainer := container.NewCenter(grid)

	logger.LogDebug("Settings Panel initialized", false)

	return centerContainer
}