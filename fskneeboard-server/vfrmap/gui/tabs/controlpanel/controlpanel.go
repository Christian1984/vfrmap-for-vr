package controlpanel

import (
	"os/exec"
	"strconv"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/msfsinterfacing"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var serverStatusBinding = binding.NewString()
var msfsConnectionBinding = binding.NewString()
var licenseBinding = binding.NewString()
var autosaveBinding = binding.NewString()

var msfsStartedBinding = binding.NewBool()
var newVersionAvailableBinding = binding.NewBool()

func UpdateServerStatus(status string) {
	serverStatusBinding.Set(status)
}

func UpdateMsfsConnectionStatus(status string) {
	msfsConnectionBinding.Set(status)
}

func UpdateLicenseStatus(status string) {
	licenseBinding.Set(status)
}

func UpdateMsfsStarted(value bool) {
	msfsStartedBinding.Set(value)
}

func UpdateNewVersionAvailable(value bool) {
	newVersionAvailableBinding.Set(value)
}

func UpdateAutosaveStatus(interval int) {
	intervalString := "Off"

	if interval > 0 {
		intervalString = strconv.Itoa(interval) + " minutes"
	}

	autosaveBinding.Set(intervalString)
}

func ControlPanel() *fyne.Container {
	logger.LogDebug("Initializing Control Panel...", false)

	//middle
	serverStatusLabel := widget.NewLabel("Server Status")
	serverStatusBinding.Set("Not Running")
	serverStatusValue := widget.NewLabelWithData(serverStatusBinding)

	msfsConnectionLabel := widget.NewLabel("Flight Simulator")
	msfsConnectionBinding.Set("Not Connected")
	msfsConnectionValue := widget.NewLabelWithData(msfsConnectionBinding)

	licenseLabel := widget.NewLabel("License")
	licenseBinding.Set("Checking...")
	licenseValue := widget.NewLabelWithData(licenseBinding)

	autosaveLabel := widget.NewLabel("Autosave Interval")
	autosaveBinding.Set("Off")
	autosaveValue := widget.NewLabelWithData(autosaveBinding)

	grid := container.NewGridWithColumns(
		2,
		licenseLabel, licenseValue,
		msfsConnectionLabel, msfsConnectionValue,
		serverStatusLabel, serverStatusValue,
		autosaveLabel, autosaveValue,
	)
	middle := container.NewCenter(grid)

	// top
	launchSimBtn := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		go msfsinterfacing.StartMsfs()
	})

	msfsStartedBinding.AddListener(binding.NewDataListener(func() {
		msfsStarted, _ := msfsStartedBinding.Get()

		if msfsStarted {
			launchSimBtn.Disable()
		} else {
			launchSimBtn.Enable()
		}
	}))
	top := container.NewHBox(launchSimBtn)
	
	// bottom
	updateInfoLabel := widget.NewLabel("A new version of FSKneeboard is available.")
	downloadUpdateBtn := widget.NewButtonWithIcon("Download Now", theme.DownloadIcon(), func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", globals.DownloadLink).Start()
	})
	bottom := container.NewHBox(updateInfoLabel, downloadUpdateBtn)

	newVersionAvailableBinding.AddListener(binding.NewDataListener(func() {
		b, _ := newVersionAvailableBinding.Get()
		bottom.Hidden = !b
	}))

	// layout
	border := layout.NewBorderLayout(top, bottom, nil, nil)
	resContainer := container.New(border, top, bottom, middle)

	logger.LogDebug("Control Panel initialized", false)

	return resContainer
}