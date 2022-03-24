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

func ControlPanel() *fyne.Container {
	//middle
	// TODO

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
	return container.New(border, top)
}