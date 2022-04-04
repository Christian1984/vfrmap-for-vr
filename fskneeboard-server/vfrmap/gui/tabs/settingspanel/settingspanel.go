package settingspanel

import (
	"strconv"
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

func UpdateAutosaveStatus(interval int) {
	intervalString := "Off"

	if interval > 0 {
		intervalString = strconv.Itoa(interval)
	}

	autosaveBinding.Set(intervalString)
}

func SettingsPanel() *fyne.Container {
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
	}))

	msfsVersionBinding.Set(msfsVersionOptionWinstore)

	// autostart select
	msfsAutostartLabel := widget.NewLabel("Flight Simulator Autostart")
	msfsAutostartCb := widget.NewCheckWithData("Start MSFS when FSKneeboard starts", msfsAutostartBinding)

	msfsAutostartBinding.AddListener(binding.NewDataListener(func() {
		msfsAutostart, _ := msfsAutostartBinding.Get()
		globals.MsfsAutostart = msfsAutostart
		logger.LogInfo("MSFS Autostart updated: " + strconv.FormatBool(msfsAutostart), false)
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
		server.UpdateAutosaveInterval()
	}))

	autosaveBinding.Set("Off")

	// set log level
	// TODO

	// grid and centerContainer
	//empty := widget.NewLabel("")
	grid := container.NewGridWithColumns(
		2,
		msfsVersionLabel, msfsVersionSelect,
		msfsAutostartLabel, msfsAutostartCb,
		autosaveLabel, autosaveSelect,
	)
	centerContainer := container.NewCenter(grid)
	return centerContainer
}