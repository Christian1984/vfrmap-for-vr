package settingspanel

import (
	"strconv"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"

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

	// set log level
	// TODO

	// set autosave properties
	// TODO

	// grid and centerContainer
	grid := container.NewGridWithColumns(
		2,
		msfsVersionLabel, msfsVersionSelect,
		msfsAutostartLabel, msfsAutostartCb,
	)
	centerContainer := container.NewCenter(grid)
	return centerContainer
}