package panelcommons

import (
	"image/color"
	"net/url"
	"strings"
	"time"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/res"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var freeImages []fyne.Resource

func addImageResource(name string) {
	asset, err := res.Asset(name)
	if err == nil {
		resource := fyne.NewStaticResource(name, asset)
		freeImages = append(freeImages, resource)
	}

}

func initImageArray() {
	addImageResource("pro-img-1.jpg")
	addImageResource("pro-img-2.jpg")
	addImageResource("pro-img-3.jpg")
}

var imageIndex = 0

func initImageRotation(image *canvas.Image) {
	if len(freeImages) == 0 || globals.Pro {
		return
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)

			imageIndex = (imageIndex + 1) % len(freeImages)

			image.Resource = freeImages[imageIndex]
			image.Refresh()
		}
	}()
}

func PremiumInfo(message string) *fyne.Container {
	backgroundColor := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})
	infoContainer := container.NewMax(backgroundColor)

	if !globals.Pro {
		textColor := color.RGBA{255, 191, 0, 255}

		freeLabel1 := canvas.NewText("  Thanks For Trying FSKneeboard FREE  ", textColor)
		freeLabel1.TextStyle.Bold = true
		freeLabel1.Alignment = fyne.TextAlignCenter

		freeLabel2 := canvas.NewText("Support the development", textColor)
		freeLabel2.Alignment = fyne.TextAlignCenter

		freeLabel3 := canvas.NewText("and unlock ALL features today!", textColor)
		freeLabel3.Alignment = fyne.TextAlignCenter

		messageLabel := canvas.NewText(message, textColor)
		messageLabel.Alignment = fyne.TextAlignCenter

		initImageArray()

		var freeImage *canvas.Image

		if len(freeImages) > 0 {
			freeImage = canvas.NewImageFromResource(freeImages[0])
			freeImage.FillMode = canvas.ImageFillOriginal
			initImageRotation(freeImage)
		}

		learnMoreUrl, _ := url.Parse("https://fskneeboard.com/compare")
		learnMoreLink := widget.NewHyperlink("Learn more about FSKneeboard PRO", learnMoreUrl)
		learnMoreLink.Alignment = fyne.TextAlignCenter

		orLabel := canvas.NewText("or", textColor)
		orLabel.Alignment = fyne.TextAlignCenter

		buyUrl, _ := url.Parse("https://fskneeboard.com/buy-now")
		buyLink := widget.NewHyperlink("BUY NOW", buyUrl)
		buyLink.Alignment = fyne.TextAlignCenter

		rightVBox := container.NewVBox(
			freeLabel1,
			canvas.NewRectangle(textColor),
		)

		if strings.TrimSpace(message) != "" {
			rightVBox.Add(messageLabel)
			rightVBox.Add(canvas.NewRectangle(textColor))
		}

		rightVBox.Add(freeLabel2)
		rightVBox.Add(freeLabel3)
		rightVBox.Add(freeImage)
		rightVBox.Add(learnMoreLink)
		rightVBox.Add(orLabel)
		rightVBox.Add(buyLink)

		rightCenter := container.NewCenter(rightVBox)
		infoContainer.Add(rightCenter)
	}

	return infoContainer
}
