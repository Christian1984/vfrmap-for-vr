package gui

import (
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/gui/tabs/console"
	"vfrmap-for-vr/vfrmap/gui/tabs/controlpanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/settingspanel"
	"vfrmap-for-vr/vfrmap/logger"
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

	logger.LogDebug("Loading icon...", false)
	r, err := fyne.LoadResourceFromPath("icon.png")
	if err == nil {
		logger.LogDebug("Icon loaded", false)
		a.SetIcon(r)
	} else {
		logger.LogWarn("Icon could not be loaded!", false)
	}

	title := "FSKneeboard "

	if globals.Pro {
		title += "PRO"
	} else {
		title += "FREE"
	}

	w = a.NewWindow(title)

	logger.LogDebug("Initializing tabs...", false)
	tabs := container.NewAppTabs(
		container.NewTabItem("Control Panel", controlpanel.ControlPanel()),
		container.NewTabItem("Settings", settingspanel.SettingsPanel()),
		container.NewTabItem("Hotkeys", widget.NewLabel("//TODO")),
		container.NewTabItem("PDF Import", widget.NewLabel("//TODO")),
		container.NewTabItem("Console", console.Console()),
	)

	logger.LogDebug("Tabs initialized", false)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 600))

	dialogs.ParentWindow = &w
}

func ShowAndRun() {
	if w != nil {
		logger.LogDebug("Showing window...", false)
		w.ShowAndRun()
	}
}