package controlpanel

import (
	"os/exec"
	"vfrmap-for-vr/vfrmap/gui/tabs/console"
	"vfrmap-for-vr/vfrmap/server"
	"vfrmap-for-vr/vfrmap/utils"

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

var serverStartedBinding = binding.NewBool()
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

func UpdateServerStarted(value bool) {
	serverStartedBinding.Set(value)
}

func UpdateNewVersionAvailable(value bool) {
	newVersionAvailableBinding.Set(value)
}

func ControlPanel() *fyne.Container {
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

	grid := container.NewGridWithColumns(
		2,
		licenseLabel, licenseValue,
		msfsConnectionLabel, msfsConnectionValue,
		serverStatusLabel, serverStatusValue,
	)
	middle := container.NewCenter(grid)

	// top
	startServerBtn := widget.NewButtonWithIcon("Start FSKneeboard", theme.MediaPlayIcon(), func() {
		console.ConsoleLogLn("Server Started")
		go server.StartFskServer()
	})

	stopServerBtn := widget.NewButtonWithIcon("Stop FSKneeboard", theme.MediaStopIcon(), func() {
		console.ConsoleLogLn("Server Stopped")
		go server.StopFskServer()
	})

	launchSimBtn := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		console.ConsoleLogLn("Launching MSFS...")
	})

	serverStartedBinding.AddListener(binding.NewDataListener(func() {
		serverStarted, _ := serverStartedBinding.Get()

		if (serverStarted) {
			startServerBtn.Disable()
			stopServerBtn.Enable()
		} else {
			startServerBtn.Enable()
			stopServerBtn.Disable()
		}
	}))

	top := container.NewHBox(startServerBtn, stopServerBtn, launchSimBtn)
	
	// bottom
	updateInfoLabel := widget.NewLabel("A new version of FSKneeboard is available.")
	downloadUpdateBtn := widget.NewButtonWithIcon("Download Now", theme.DownloadIcon(), func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://fskneeboard.com/download-latest").Start()
	})
	bottom := container.NewHBox(updateInfoLabel, downloadUpdateBtn)

	newVersionAvailableBinding.AddListener(binding.NewDataListener(func() {
		utils.Println("newVersionAvailable DataListener called!")
		b, _ := newVersionAvailableBinding.Get()
		bottom.Hidden = !b
	}))

	// layout
	border := layout.NewBorderLayout(top, bottom, nil, nil)
	return container.New(border, top, bottom, middle)
}