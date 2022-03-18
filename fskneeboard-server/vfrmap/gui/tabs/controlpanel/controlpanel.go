package controlpanel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var consoleText string
var console *widget.TextGrid
var consoleScroll *container.Scroll

func ConsoleLog(message string) {
	consoleText += message

	if console != nil {
		console.SetText(consoleText)
	}

	if (consoleScroll != nil) {
		consoleScroll.ScrollToBottom()
	}

}

func ConsoleLogLn(message string) {
	if len(consoleText) > 0 {
		message += "\n"
	}

	ConsoleLog(message)
}

func ControlPanel() *fyne.Container {
	consoleText = ""
	console = widget.NewTextGrid()
	console.ShowLineNumbers = true

	consoleScroll = container.NewScroll(console)

	startServer := widget.NewButtonWithIcon("Start FSKneeboard", theme.MediaPlayIcon(), func() {
		ConsoleLogLn("Server Started")
	})

	stopServer := widget.NewButtonWithIcon("Stop FSKneeboard", theme.MediaStopIcon(), func() {
		ConsoleLogLn("Server Stopped")
	})

	launchSim := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		ConsoleLogLn("Launching MSFS...")
	})

	top := container.NewHBox(startServer, stopServer, launchSim)
	middle := container.NewMax(consoleScroll)
	border := layout.NewBorderLayout(top, nil, nil, nil)
	
	return container.New(border, top, middle)
}