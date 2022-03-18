package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ControlPanel() *fyne.Container {
	consoleText := ""
	console := widget.NewTextGrid()
	console.ShowLineNumbers = true

	consoleScroll := container.NewScroll(console)

	appendToConsole := func(line string) {
		if len(consoleText) > 0 {
			consoleText += "\n"
		}

		consoleText += line
		console.SetText(consoleText)
		consoleScroll.ScrollToBottom()
	}

	startServer := widget.NewButtonWithIcon("Start FSKneeboard", theme.MediaPlayIcon(), func() {
		appendToConsole("Server Started")
	})

	stopServer := widget.NewButtonWithIcon("Stop FSKneeboard", theme.MediaStopIcon(), func() {
		appendToConsole("Server Stopped")
	})

	launchSim := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		appendToConsole("Launching MSFS...")
	})

	top := container.NewHBox(startServer, stopServer, launchSim)
	middle := container.NewMax(consoleScroll)
	border := layout.NewBorderLayout(top, nil, nil, nil)
	
	return container.New(border, top, middle)
}