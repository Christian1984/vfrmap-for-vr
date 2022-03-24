package gui

import (
	"vfrmap-for-vr/vfrmap/gui/tabs/console"
	"vfrmap-for-vr/vfrmap/gui/tabs/controlpanel"
	"vfrmap-for-vr/vfrmap/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window

func InitGui() {
	utils.Println("Starting FSKneeboard GUI...")

	a := app.New()
	w = a.NewWindow("FSKneeboard")

	tabs := container.NewAppTabs(
		container.NewTabItem("Control Panel", controlpanel.ControlPanel()),
		container.NewTabItem("Settings", widget.NewLabel("//TODO")),
		container.NewTabItem("Hotkeys", widget.NewLabel("//TODO")),
		container.NewTabItem("PDF Import", widget.NewLabel("//TODO")),
		container.NewTabItem("Console", console.Console()),
	)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 600))
}

func ShowAndRun() {
	if w != nil {
		w.ShowAndRun()
	}
}