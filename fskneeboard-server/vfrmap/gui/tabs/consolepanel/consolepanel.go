package consolepanel

import (
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var consoleBinding binding.String
var autoScrollBinding binding.Bool
var console *widget.Label
var consoleScroll *container.Scroll

func ClearConsole() {
	if consoleBinding != nil {
		consoleBinding.Set("")
	}
}

func ConsoleLog(message string) {
	if consoleBinding != nil {
		consoleText, err := consoleBinding.Get()

		if err == nil {
			consoleBinding.Set(consoleText + message)
			conditionalAutoScroll()
		}
	}
}

func ConsoleLogLn(message string) {
	ConsoleLog(message + "\n")
}

func ConsolePanel() *fyne.Container {
	logger.LogDebug("Initializing Console Panel...", false)

	// console
	consoleBinding = binding.NewString()
	console = widget.NewLabelWithData(consoleBinding)

	// scroll
	consoleScroll = container.NewScroll(console)

	autoScrollBinding = binding.NewBool()

	autoScrollBinding.AddListener(binding.NewDataListener(func() {
		conditionalAutoScroll()
	}))

	middle := container.NewMax(consoleScroll)
	
	// bottom
	clearLogBtn := widget.NewButtonWithIcon("Clear Console Output", theme.ContentClearIcon(), func() {
		ClearConsole()
	})

	scrollToBottomCb := widget.NewCheckWithData("Enable Autoscroll", autoScrollBinding)

	scrollToBottomCb.SetChecked(true)
	
	bottom := container.NewHBox(clearLogBtn, scrollToBottomCb)

	border := layout.NewBorderLayout(nil, bottom, nil, nil)
	resContainer := container.New(border, bottom, middle)

	logger.LogDebug("Console Panel initialized", false)

	return resContainer
}

func conditionalAutoScroll() {
	autoScroll, _ := autoScrollBinding.Get()

	if (autoScroll && consoleScroll != nil) {
		consoleScroll.ScrollToBottom()
	}
}