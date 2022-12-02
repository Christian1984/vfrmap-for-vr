package supportpanel

import (
	"net/url"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SupportPanel() *fyne.Container {
	logger.LogDebug("Initializing Support Panel...")

	docsUrl, _ := url.Parse("https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md#troubleshooting")
	discordUrl, _ := url.Parse("https://discord.fskneeboard.com")

	introLabel := widget.NewLabel("If you encounter any problems, please try this:")
	introLabel.Alignment = fyne.TextAlignCenter

	docsLabel := widget.NewLabel("Step 1: Please check out the FSKneeboard manual, especially the troubleshooting section here...")
	docsLink := widget.NewHyperlink("Read the FSKneeboard Troubleshooting Guide", docsUrl)

	docsLabel.Alignment = fyne.TextAlignCenter
	docsLink.Alignment = fyne.TextAlignCenter

	discordLabel := widget.NewLabel("Step 2: You are always welcome to join us on Discord and ask questions or leave feedback...")
	discordLink := widget.NewHyperlink("Join the FSKneeboard Discord Server", discordUrl)

	discordLabel.Alignment = fyne.TextAlignCenter
	discordLink.Alignment = fyne.TextAlignCenter

	vBox := container.NewVBox(
		introLabel,
		widget.NewLabel(""),
		docsLabel,
		docsLink,
		widget.NewLabel(""),
		discordLabel,
		discordLink,
	)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebugVerbose("Support Panel initialized")

	return centerContainer

}
