package controlpanel

import (
	"vfrmap-for-vr/vfrmap/gui/tabs/console"
	"vfrmap-for-vr/vfrmap/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var serverStatusValue *widget.Label
var msfsConnectionValue *widget.Label
var licenseValue *widget.Label

func updateValue(labelWidget *widget.Label, value string) {
	if labelWidget != nil {
		labelWidget.SetText(value)
	}
}

func UpdateServerStatus(status string) {
	updateValue(serverStatusValue, status)
}

func UpdateMsfsConnectionStatus(status string) {
	updateValue(msfsConnectionValue, status)
}

func UpdateLicenseStatus(status string) {
	updateValue(licenseValue, status)
}

func ControlPanel() *fyne.Container {
	//middle
	serverStatusLabel := widget.NewLabel("Server Status")
	serverStatusValue = widget.NewLabel("Not Running")

	msfsConnectionLabel := widget.NewLabel("Flight Simulator")
	msfsConnectionValue = widget.NewLabel("Not Connected")

	licenseLabel := widget.NewLabel("License")
	licenseValue = widget.NewLabel("Not Valid")

	grid := container.NewGridWithColumns(2, 
		serverStatusLabel, serverStatusValue,
		msfsConnectionLabel, msfsConnectionValue,
		licenseLabel, licenseValue)
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

	/*test := widget.NewButtonWithIcon("Count", theme.UploadIcon(), func() {
		console.ConsoleLogLn("Counting")

		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(1 * time.Second)
				go console.ConsoleLogLn("Counter: " + strconv.Itoa(i))
			}
		}()
	})*/

	top := container.NewHBox(startServerBtn, stopServerBtn, launchSimBtn)
	
	// bottom
	// TODO

	// layout
	border := layout.NewBorderLayout(top, nil, nil, nil)
	return container.New(border, top, middle)
}