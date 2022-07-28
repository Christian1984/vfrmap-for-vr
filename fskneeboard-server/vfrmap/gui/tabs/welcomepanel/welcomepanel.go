package welcomepanel

import (
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func WelcomePanel() *fyne.Container {
	logger.LogDebug("Initializing Welcome Panel...")

	// docsUrl, _ := url.Parse("https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md#troubleshooting")
	// discordUrl, _ := url.Parse("https://discord.fskneeboard.com")

	/*
		introLabel := widget.NewLabel("If you encounter any problems, please try this:")
		introLabel.Alignment = fyne.TextAlignCenter

		dismissTourBtn := widget.NewButton("End Tour", func() {
			logger.LogDebugVerbose("Dismissing gui tour...")
			callbacks.ShowGuiTourChangedCallback(false)
		})

		vBox := container.NewVBox(
			introLabel,
		)

		bottom := container.NewCenter(dismissTourBtn)

		middle := container.NewScroll(vBox)
		border := layout.NewBorderLayout(nil, bottom, nil, nil)
		resContainer := container.New(border, bottom, middle)
	*/

	logger.LogDebugVerboseOverride("Support Panel initialized", false)

	// return resContainer
	return container.NewCenter()
}
