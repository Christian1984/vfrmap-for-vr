package controlpanel

import (
	"image/color"
	"net/url"
	"os/exec"
	"strconv"
	"time"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/msfsinterfacing"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var serverStatusBinding = binding.NewString()
var msfsConnectionBinding = binding.NewString()
var licenseBinding = binding.NewString()
var autosaveBinding = binding.NewString()

var msfsStartedBinding = binding.NewBool()
var newVersionAvailableBinding = binding.NewBool()

var freeImages []fyne.Resource

func UpdateServerStatus(status string) {
	serverStatusBinding.Set(status)
}

func UpdateMsfsConnectionStatus(status string) {
	msfsConnectionBinding.Set(status)
}

func UpdateLicenseStatus(status string) {
	licenseBinding.Set(status)
}

func UpdateMsfsStarted(value bool) {
	msfsStartedBinding.Set(value)
}

func UpdateNewVersionAvailable(value bool) {
	newVersionAvailableBinding.Set(value)
}

func UpdateAutosaveStatus(interval int) {
	intervalString := "Off"

	if interval > 0 {
		intervalString = strconv.Itoa(interval) + " minutes"
	}

	autosaveBinding.Set(intervalString)
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

func ControlPanel() *fyne.Container {
	logger.LogDebug("Initializing Control Panel...", false)

	//middle
	serverStatusLabel := widget.NewLabel("Server Status")
	serverStatusBinding.Set("Not Running")
	serverStatusValue := widget.NewLabelWithData(serverStatusBinding)

	msfsConnectionLabel := widget.NewLabel("Flight Simulator")
	msfsConnectionBinding.Set("Not Connected")
	msfsConnectionValue := widget.NewLabelWithData(msfsConnectionBinding)

	licenseLabel := widget.NewLabel("License")
	licenseBinding.Set("Checking...")
	licenseValue := widget.NewLabelWithData(licenseBinding)

	autosaveLabel := widget.NewLabel("Autosave Interval")
	autosaveBinding.Set("Off")
	autosaveValue := widget.NewLabelWithData(autosaveBinding)

	grid := container.NewGridWithColumns(
		2,
		licenseLabel, licenseValue,
		msfsConnectionLabel, msfsConnectionValue,
		serverStatusLabel, serverStatusValue,
		autosaveLabel, autosaveValue,
	)
	middle := container.NewCenter(grid)

	// top
	launchSimBtn := widget.NewButtonWithIcon("Launch Flight Simulator", theme.UploadIcon(), func() {
		go msfsinterfacing.StartMsfs()
	})

	msfsStartedBinding.AddListener(binding.NewDataListener(func() {
		msfsStarted, _ := msfsStartedBinding.Get()

		if msfsStarted {
			launchSimBtn.Disable()
		} else {
			launchSimBtn.Enable()
		}
	}))
	top := container.NewHBox(launchSimBtn)
	
	// bottom
	updateInfoLabel := widget.NewLabel("A new version of FSKneeboard is available.")
	downloadUpdateBtn := widget.NewButtonWithIcon("Download Now", theme.DownloadIcon(), func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", globals.DownloadLink).Start()
	})
	bottom := container.NewHBox(updateInfoLabel, downloadUpdateBtn)

	newVersionAvailableBinding.AddListener(binding.NewDataListener(func() {
		b, _ := newVersionAvailableBinding.Get()
		bottom.Hidden = !b
	}))

	//right
	textColor := color.RGBA{255, 191, 0, 255}

	freeLabel1 := canvas.NewText("  Thanks For Trying FSKneeboard FREE  ", textColor)
	freeLabel1.TextStyle.Bold = true
	freeLabel1.Alignment = fyne.TextAlignCenter
	
	freeLabel2 := canvas.NewText("Support the development", textColor)
	freeLabel2.Alignment = fyne.TextAlignCenter
	
	freeLabel3 := canvas.NewText("and unlock ALL features today!", textColor)
	freeLabel3.Alignment = fyne.TextAlignCenter

	freeImage1, err := fyne.LoadResourceFromPath("res/pro-img-1.jpg")
	if err == nil {
		freeImages = append(freeImages, freeImage1)
	}

	if !globals.Pro {
		freeImage2, err := fyne.LoadResourceFromPath("res/pro-img-2.jpg")
		if err == nil {
			freeImages = append(freeImages, freeImage2)
		}
	
		freeImage3, err := fyne.LoadResourceFromPath("res/pro-img-3.jpg")
		if err == nil {
			freeImages = append(freeImages, freeImage3)
		}
	}

	freeImage := canvas.NewImageFromResource(freeImage1)
	freeImage.FillMode = canvas.ImageFillOriginal
	initImageRotation(freeImage)

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
		freeLabel2,
		freeLabel3,
		freeImage,
		learnMoreLink,
		orLabel,
		buyLink,
	)

	rightCenter := container.NewCenter(rightVBox)

	// background-color
	backgroundColor := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})
	right := container.NewMax(backgroundColor, rightCenter)

	right.Hidden = globals.Pro

	// layout
	border := layout.NewBorderLayout(top, bottom, nil, right)
	resContainer := container.New(border, top, bottom, right, middle)

	logger.LogDebug("Control Panel initialized", false)

	return resContainer
}