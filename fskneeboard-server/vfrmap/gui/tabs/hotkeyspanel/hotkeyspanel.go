package hotkeyspanel

import (
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func HotkeysPanel() *fyne.Container {
	logger.LogDebug("Initializing Hotkeys Panel...", false)

	// grid and centerContainer
	label := widget.NewLabel("Test...")
	grid := container.NewGridWithColumns(
		3,
		widget.NewLabel("Test1"), widget.NewLabel("Test2"), widget.NewLabel("Test3"),
		//msfsAutostartLabel, msfsAutostartCb, widget.NewLabel(""),
	)
	vBox := container.NewVBox(label, grid)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebug("Hotkeys Panel initialized", false)

	return centerContainer

}