package welcomepanel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WelcomePanel() *fyne.Container {
	logger.LogDebug("Initializing Welcome Panel...")

	// docsUrl, _ := url.Parse("https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md#troubleshooting")
	// discordUrl, _ := url.Parse("https://discord.fskneeboard.com")

	introLabel := widget.NewLabel("If you encounter any problems, please try this:")
	introLabel.Alignment = fyne.TextAlignCenter

	dismissTourBtn := widget.NewButton("End Tour", func() {
		logger.LogDebug("Dismissing gui tour...")

		globals.TourIndexStarted = false
		dbmanager.StoreTourStates()

		callbacks.GuiTourStartedChangedCallback()
	})

	vBox := container.NewVBox(
		introLabel,
		dismissTourBtn,
	)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebugVerboseOverride("Support Panel initialized", false)

	return centerContainer
}
