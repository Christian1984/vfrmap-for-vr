package settingspanel

import (
	"strconv"
	"strings"
	"vfrmap-for-vr/_vendor/premium/autosave"
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

var autosaveOptions = []string{"Off", "1", "5", "10", "15", "30", "60"}
var autosaveBinding = binding.NewString()

var tourStartedBinding = binding.NewBool()

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

func UpdateLogLevelStatus(level string) {
	lowerLevel := strings.ToLower(level)

	if lowerLevel != logger.Debug && lowerLevel != logger.Info && lowerLevel != logger.Warn && lowerLevel != logger.Error {
		lowerLevel = "off"
	}

	loglevelBinding.Set(lowerLevel)
}

func SettingsPanel() *fyne.Container {
	logger.LogDebugVerboseOverride("Initializing Settings Panel...", false)

	// MSFS version select
	msfsVersionLabel := widget.NewLabel("Flight Simulator Version")
	msfsVersionSelect := widget.NewSelect(msfsVersionOptions, func(selected string) {
		msfsVersionBinding.Set(selected)
	})

	msfsVersionBinding.AddListener(binding.NewDataListener(func() {
		selected, _ := msfsVersionBinding.Get()

		globals.SteamFs = false
		globals.WinstoreFs = false

		if selected == msfsVersionOptionWinstore {
			globals.WinstoreFs = true
			logger.LogInfoVerboseOverride("Selected MSFS Version: Windows Store", false)
		} else if selected == msfsVersionOptionSteam {
			globals.SteamFs = true
			logger.LogInfoVerboseOverride("Selected MSFS Version: Steam", false)
		} else {
			return
		}

		if strings.ToLower(selected) != strings.ToLower(msfsVersionSelect.Selected) {
			logger.LogDebugVerboseOverride("msfsVersionBinding changed: ["+selected+"]; updating ui select element...", false)
			msfsVersionSelect.SetSelected(selected)
		} else {
			logger.LogDebugVerboseOverride("msfsVersionBinding change listener: ui select element already up to date => ["+selected+"]", false)
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
		logger.LogInfoVerboseOverride("MSFS Autostart updated: "+strconv.FormatBool(msfsAutostart), false)

		dbmanager.StoreMsfsAutostart()
	}))

	// set autosave properties
	autosaveLabel := widget.NewLabel("Autosave Interval [minutes]")
	autosaveSelect := widget.NewSelect(autosaveOptions, func(selected string) {
		autosaveBinding.Set(selected)
	})
	autosaveOpenFolderBtn := widget.NewButton("Open Autosave Folder", func() {
		autosave.OpenAutosaveFolder()
	})

	autosaveBinding.AddListener(binding.NewDataListener(func() {
		autosaveString, bindingErr := autosaveBinding.Get()

		if bindingErr != nil {
			logger.LogErrorVerboseOverride(bindingErr.Error(), false)
		}

		if autosaveString != "Off" && !globals.Pro {
			autosaveString = "Off"
			dialogs.ShowProFeatureInfo("Autosave")
		}

		for _, v := range autosaveOptions {
			if strings.ToLower(v) == strings.ToLower(autosaveString) {
				if strings.ToLower(autosaveString) != strings.ToLower(autosaveSelect.Selected) {
					logger.LogDebugVerboseOverride("autosaveBinding changed: ["+autosaveString+"]; updating ui select element...", false)
					autosaveSelect.SetSelected(autosaveString)
				} else {
					logger.LogDebugVerboseOverride("autosaveBinding change listener: ui select element already up to date => ["+autosaveString+"]", false)
				}
				break
			}
		}

		autosaveInterval, err := strconv.Atoi(autosaveString)
		if err != nil {
			autosaveInterval = 0
		}

		hasChanged := globals.AutosaveInterval != autosaveInterval

		if hasChanged {
			globals.AutosaveInterval = autosaveInterval
			dbmanager.StoreAutosaveInterval()
		}

		server.UpdateAutosaveInterval(hasChanged)
	}))

	autosaveBinding.Set("Off")

	// set log level
	loglevelLabel := widget.NewLabel("Log Level")
	logevelWarningLabel := widget.NewLabel("WARNING: The Log Level \"Debug\" may result in very large log files!")
	logevelWarningLabel.Hidden = true
	logevelWarningLabel.Alignment = fyne.TextAlignCenter

	loglevelSelect := widget.NewSelect(loglevelOptions, func(selected string) {
		logevelWarningLabel.Hidden = strings.ToLower(selected) != "debug"
		loglevelBinding.Set(selected)
	})
	logsOpenFolderBtn := widget.NewButton("Open Log Folder", func() {
		logger.OpenLogFolder()
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

		if strings.ToLower(loglevelString) != strings.ToLower(loglevelSelect.Selected) {
			logger.LogDebugVerboseOverride("loglevelBinding changed: ["+loglevelString+"]; updating ui select element...", false)
			loglevelSelect.SetSelected(loglevelOptions[matchIndex])
		} else {
			logger.LogDebugVerboseOverride("loglevelBinding change listener: ui select element already up to date => ["+loglevelString+"]", false)
		}

		globals.LogLevel = strings.ToLower(loglevelString)

		dbmanager.StoreLogLevel()

		logger.SetLevel(loglevelString)
		logger.TryCreateLogFile()
	}))

	loglevelBinding.Set(globals.LogLevel)

	restartTourLabel := widget.NewLabel("Ingame Tutorial Tour")
	restartTourBtn := widget.NewButton("Restart Tour", func() {
		logger.LogDebug("Resetting ingame panel tour...")

		tourStartedBinding.Set(false)

		globals.TourIndexStarted = false
		globals.TourMapStarted = false
		globals.TourChartsStarted = false
		globals.TourNotepadStarted = false

		//dbmanager.StoreTourStates()
	})

	/*
		tourStartedBinding.AddListener(binding.NewDataListener(func() {
			tourStarted, _ := tourStartedBinding.Get()
			if tourStarted {
				restartTourBtn.Enable()
			} else {
				restartTourBtn.Disable()
			}
		}))

		tourStartedBinding.Set(true) // remove
	*/

	// grid and centerContainer
	grid := container.NewGridWithColumns(
		3,
		msfsVersionLabel, msfsVersionSelect, widget.NewLabel(""),
		msfsAutostartLabel, msfsAutostartCb, widget.NewLabel(""),
		restartTourLabel, restartTourBtn, widget.NewLabel(""),
		autosaveLabel, autosaveSelect, autosaveOpenFolderBtn,
		loglevelLabel, loglevelSelect, logsOpenFolderBtn,
	)
	vBox := container.NewVBox(grid, logevelWarningLabel)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebugVerboseOverride("Settings Panel initialized", false)

	return centerContainer
}
