package controlpanel

import (
	"vfrmap-for-vr/vfrmap/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var consoleBinding binding.String
var autoScroll = true
var console *widget.Label
var consoleScroll *container.Scroll

func ClearConsole() {
	if (consoleBinding != nil) {
		consoleBinding.Set("")
	}
}

func ConsoleLog(message string) {
	if (consoleBinding != nil) {
		consoleText, err := consoleBinding.Get()

		if (err == nil) {
			consoleBinding.Set(consoleText + message)

			if (consoleScroll != nil && autoScroll) {
				consoleScroll.ScrollToBottom()
			}
		}
	}
}

func ConsoleLogLn(message string) {
	ConsoleLog(message + "\n")
}

func ControlPanel() *fyne.Container {
	// console
	consoleBinding = binding.NewString()
	console = widget.NewLabelWithData(consoleBinding)
	consoleScroll = container.NewScroll(console)

	middle := container.NewMax(consoleScroll)
	
	// top
	startServerBtn := widget.NewButtonWithIcon("Start FSKneeboard", theme.MediaPlayIcon(), func() {
		ConsoleLogLn("Server Started")
		go server.StartFskServer()
	})

	stopServerBtn := widget.NewButtonWithIcon("Stop FSKneeboard", theme.MediaStopIcon(), func() {
		ConsoleLogLn("Server Stopped")
		go server.StopFskServer()
	})

	launchSimBtn := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		ConsoleLogLn("Launching MSFS...")
	})

	top := container.NewHBox(startServerBtn, stopServerBtn, launchSimBtn)
	
	// bottom
	clearLogBtn := widget.NewButtonWithIcon("Clear Console Output", theme.ContentClearIcon(), func() {
		ClearConsole()
	})

	scrollToBottomCb := widget.NewCheck("Enable Autoscroll", func(checked bool) {
		autoScroll = checked

		if (autoScroll) {
			consoleScroll.ScrollToBottom()
		}
	})

	scrollToBottomCb.SetChecked(true)
	
	bottom := container.NewHBox(clearLogBtn, scrollToBottomCb)

	// layout
	border := layout.NewBorderLayout(top, bottom, nil, nil)	
	return container.New(border, top, bottom, middle)
}