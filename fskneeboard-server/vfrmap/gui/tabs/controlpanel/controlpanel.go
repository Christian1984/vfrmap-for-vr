package controlpanel

import (
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/msfsinterfacing"
	"vfrmap-for-vr/vfrmap/gui/tabs/panelcommons"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var serverStatusBinding = binding.NewString()
var serverAddressBinding = binding.NewString()
var msfsConnectionBinding = binding.NewString()
var licenseBinding = binding.NewString()
var autosaveBinding = binding.NewString()

var msfsStartedBinding = binding.NewBool()
var newVersionAvailableBinding = binding.NewBool()

func UpdateServerStatus(statusMessage string, url string) {
	serverStatusBinding.Set(statusMessage)
	serverAddressBinding.Set(url)
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

func ControlPanel() *fyne.Container {
	logger.LogDebug("Initializing Control Panel...")

	//middle
	serverStatusLabel := widget.NewLabel("Server Status")
	serverStatusBinding.Set("Not Running")
	serverStatusValue := widget.NewLabelWithData(serverStatusBinding)
	serverAddressValue := widget.NewHyperlink("Test", nil)
	serverStatusHBox := container.NewHBox(serverStatusValue, serverAddressValue)

	serverAddressBinding.AddListener(binding.NewDataListener(func() {
		serverAddress, err := serverAddressBinding.Get()

		if err == nil {
			trimmedServerAddress := strings.TrimSpace(serverAddress)

			if trimmedServerAddress == "" {
				serverAddressValue.SetURL(nil)
				serverAddressValue.Hide()
			} else {
				serverUrl, parseErr := url.Parse(serverAddress)

				if parseErr != nil {
					logger.LogWarnVerbose("Could not parse server address from string " + serverAddress + ", reason: " + parseErr.Error())
				} else {
					serverAddressValue.SetURL(serverUrl)
					serverAddressValue.Show()
				}
			}

			serverAddressValue.SetText(trimmedServerAddress)
		}
	}))

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
		serverStatusLabel, serverStatusHBox,
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
	right := panelcommons.PremiumInfo("")
	right.Hidden = globals.Pro

	// layout
	border := layout.NewBorderLayout(top, bottom, nil, right)
	resContainer := container.New(border, top, bottom, right, middle)

	logger.LogDebugVerboseOverride("Control Panel initialized", false)

	return resContainer
}
