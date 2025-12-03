package settingspanel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"vfrmap-for-vr/_vendor/premium/autosave"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/secrets"
	"vfrmap-for-vr/vfrmap/gui/callbacks"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const msfsVersionOption2020Winstore = "Flight Simulator 2020 (Windows Store)"
const msfsVersionOption2020Steam = "Flight Simulator 2020 (Steam)"
const msfsVersionOption2024Winstore = "Flight Simulator 2024 (Windows Store)"
const msfsVersionOption2024Steam = "Flight Simulator 2024 (Steam)"

var msfsVersionOptions = []string{
	msfsVersionOption2020Winstore,
	msfsVersionOption2020Steam,
	msfsVersionOption2024Winstore,
	msfsVersionOption2024Steam,
}
var msfsVersionBinding = binding.NewString()
var msfsAutostartBinding = binding.NewBool()

var autosaveOptions = []string{"Off", "1", "5", "10", "15", "30", "60"}
var autosaveBinding = binding.NewString()

var interfaceScaleBinding = binding.NewFloat()
var interfaceScaleStringBinding = binding.NewString()

var oaipApiKeyBinding = binding.NewString()
var oaipBypassCacheBinding = binding.NewBool()

var bingMapsApiKeyBinding = binding.NewString()
var googleMapsApiKeyBinding = binding.NewString()

var loglevelOptions = []string{
	strings.Title(logger.Off),
	strings.Title(logger.Error),
	strings.Title(logger.Warn),
	strings.Title(logger.Info),
	strings.Title(logger.Debug),
	strings.Title(logger.Silly),
}
var loglevelBinding = binding.NewString()

func UpdateAutosaveStatus(interval int) {
	intervalString := "Off"

	if interval > 0 {
		intervalString = strconv.Itoa(interval)
	}

	autosaveBinding.Set(intervalString)
}

func UpdateMsfsVersionStatus(version string) {
	switch version {
	case "2020-steam":
		msfsVersionBinding.Set(msfsVersionOption2020Steam)
	case "2020-winstore":
		msfsVersionBinding.Set(msfsVersionOption2020Winstore)
	case "2024-steam":
		msfsVersionBinding.Set(msfsVersionOption2024Steam)
	case "2024-winstore":
		msfsVersionBinding.Set(msfsVersionOption2024Winstore)
	default:
		msfsVersionBinding.Set(msfsVersionOption2020Winstore) // default fallback
	}
}

func UpdateMsfsAutostartStatus(autostart bool) {
	msfsAutostartBinding.Set(autostart)
}

func UpdateInterfaceScale(scale float64) {
	interfaceScaleBinding.Set(scale)
}

func UpdateOpenAipApiKey(apiKey string) {
	oaipApiKeyBinding.Set(apiKey)
}

func UpdateBingMapsApiKey(apiKey string) {
	bingMapsApiKeyBinding.Set(apiKey)
}

func UpdateGoogleMapsApiKey(apiKey string) {
	googleMapsApiKeyBinding.Set(apiKey)
}

func UpdateOpenAipBypassCache(bypassCache bool) {
	oaipBypassCacheBinding.Set(bypassCache)
}

func UpdateLogLevelStatus(level string) {
	lowerLevel := strings.ToLower(level)

	if lowerLevel != logger.Silly && lowerLevel != logger.Debug && lowerLevel != logger.Info && lowerLevel != logger.Warn && lowerLevel != logger.Error {
		lowerLevel = "off"
	}

	loglevelBinding.Set(lowerLevel)
}

func SettingsPanel() *fyne.Container {
	logger.LogDebug("Initializing Settings Panel...")

	// MSFS version select
	msfsVersionLabel := widget.NewLabel("Flight Simulator Version")
	msfsVersionSelect := widget.NewSelect(msfsVersionOptions, func(selected string) {
		msfsVersionBinding.Set(selected)
	})

	msfsVersionBinding.AddListener(binding.NewDataListener(func() {
		selected, _ := msfsVersionBinding.Get()

		switch selected {
		case msfsVersionOption2020Winstore:
			globals.MsfsVersion = "2020-winstore"
			logger.LogInfo("Selected MSFS Version: 2020 Windows Store")
		case msfsVersionOption2020Steam:
			globals.MsfsVersion = "2020-steam"
			logger.LogInfo("Selected MSFS Version: 2020 Steam")
		case msfsVersionOption2024Winstore:
			globals.MsfsVersion = "2024-winstore"
			logger.LogInfo("Selected MSFS Version: 2024 Windows Store")
		case msfsVersionOption2024Steam:
			globals.MsfsVersion = "2024-steam"
			logger.LogInfo("Selected MSFS Version: 2024 Steam")
		default:
			globals.MsfsVersion = "2020-winstore" // default fallback
			logger.LogInfo("Selected MSFS Version: 2020 Windows Store (default)")
		}

		if strings.ToLower(selected) != strings.ToLower(msfsVersionSelect.Selected) {
			logger.LogDebug("msfsVersionBinding changed: [" + selected + "]; updating ui select element...")
			msfsVersionSelect.SetSelected(selected)
		} else {
			logger.LogDebug("msfsVersionBinding change listener: ui select element already up to date => [" + selected + "]")
		}

		dbmanager.StoreMsfsVersion()

		// Trigger callback to update UI
		callbacks.MsfsVersionChangedString(globals.MsfsVersion)
	}))

	msfsVersionBinding.Set(msfsVersionOption2020Winstore)

	// msfs autostart select
	msfsAutostartLabel := widget.NewLabel("Flight Simulator Autostart")
	msfsAutostartCb := widget.NewCheckWithData("Start MSFS When FSKneeboard Starts", msfsAutostartBinding)

	msfsAutostartBinding.AddListener(binding.NewDataListener(func() {
		msfsAutostart, _ := msfsAutostartBinding.Get()
		globals.MsfsAutostart = msfsAutostart
		logger.LogInfo("MSFS Autostart updated: " + strconv.FormatBool(msfsAutostart))

		dbmanager.StoreMsfsAutostart()
	}))

	// set autosave properties
	autosaveLabel := widget.NewLabel("Autosave Interval [minutes]")
	autosaveSelect := widget.NewSelect(autosaveOptions, func(selected string) {
		autosaveBinding.Set(selected)
	})
	autosaveOpenFolderBtn := widget.NewButton("Open Autosave Folder", func() {
		autosave.OpenAutosaveFolder()
	})

	autosaveBinding.AddListener(binding.NewDataListener(func() {
		autosaveString, bindingErr := autosaveBinding.Get()

		if bindingErr != nil {
			logger.LogError(bindingErr.Error())
		}

		if autosaveString != "Off" && !globals.Pro {
			autosaveString = "Off"
			dialogs.ShowProFeatureInfo("Autosave")
		}

		for _, v := range autosaveOptions {
			if strings.ToLower(v) == strings.ToLower(autosaveString) {
				if strings.ToLower(autosaveString) != strings.ToLower(autosaveSelect.Selected) {
					logger.LogDebug("autosaveBinding changed: [" + autosaveString + "]; updating ui select element...")
					autosaveSelect.SetSelected(autosaveString)
				} else {
					logger.LogDebug("autosaveBinding change listener: ui select element already up to date => [" + autosaveString + "]")
				}
				break
			}
		}

		autosaveInterval, err := strconv.Atoi(autosaveString)
		if err != nil {
			autosaveInterval = 0
		}

		hasChanged := globals.AutosaveInterval != autosaveInterval

		if hasChanged {
			globals.AutosaveInterval = autosaveInterval
			dbmanager.StoreAutosaveInterval()
		}

		server.UpdateAutosaveInterval(hasChanged)
	}))

	autosaveBinding.Set("Off")

	// reset ingame tour
	restartTourLabel := widget.NewLabel("Ingame Tutorial Tour")
	restartTourBtn := widget.NewButton("Restart Tour", func() {
		logger.LogDebug("Resetting ingame panel tour...")

		globals.TourIndexStarted = false
		globals.TourMapStarted = false
		globals.TourChartsStarted = false
		globals.TourNotepadStarted = false
		globals.TourGuiStarted = false

		dbmanager.StoreTourStates()

		dialogs.ShowTourRestartedSuccessful()
		callbacks.ShowGuiTourChanged(true)
	})

	// set interface scale
	interfaceScaleLabel := widget.NewLabel("Interface Scale [%]")
	interfaceScaleSlider := widget.NewSliderWithData(globals.InterfaceScaleMin, globals.InterfaceScaleMax, interfaceScaleBinding)
	interfaceScaleSlider.Step = 0.1
	interfaceScaleSliderLabel := widget.NewLabelWithData(interfaceScaleStringBinding)

	interfaceScale2DBtn := widget.NewButton("Optimize for 2D", func() {
		interfaceScaleBinding.Set(globals.DefaultInterfaceScale2D)
	})

	interfaceScaleVRBtn := widget.NewButton("Optimize for VR", func() {
		interfaceScaleBinding.Set(globals.DefaultInterfaceScaleVR)
	})

	interfaceScaleBinding.AddListener(binding.NewDataListener(func() {
		interfaceScale, bindingErr := interfaceScaleBinding.Get()

		if bindingErr != nil {
			logger.LogError(bindingErr.Error())
			return
		}

		percent := fmt.Sprintf("%d%%", int64(math.Floor(interfaceScale*100)))

		interfaceScaleStringBinding.Set(percent)

		hasChanged := globals.InterfaceScale != interfaceScale

		if hasChanged {
			globals.InterfaceScale = interfaceScale
			dbmanager.StoreInterfaceScale()
		}
	}))

	interfaceScaleBinding.Set(1)

	// general settings grid
	generalGrid := container.NewGridWithColumns(
		3,
		msfsVersionLabel, msfsVersionSelect, widget.NewLabel(""),
		msfsAutostartLabel, msfsAutostartCb, widget.NewLabel(""),
		restartTourLabel, restartTourBtn, widget.NewLabel(""),
		autosaveLabel, autosaveSelect, autosaveOpenFolderBtn,
		interfaceScaleLabel, container.NewBorder(nil, nil, nil, interfaceScaleSliderLabel, interfaceScaleSlider), container.NewGridWithColumns(2, interfaceScale2DBtn, interfaceScaleVRBtn),
	)

	//API KEYS
	// oaip api key
	oaipBypassCacheCb := widget.NewCheckWithData("Deactivate & Bypass Local Cache", oaipBypassCacheBinding)
	oaipBypassCacheCb.Disable()

	oaipBypassCacheBinding.AddListener(binding.NewDataListener(func() {
		// if oaipBypassCacheCb.Disabled() {
		// 	oaipBypassCacheBinding.Set(false)
		// }

		oaipBypassCache, _ := oaipBypassCacheBinding.Get()
		globals.OpenAipBypassCache = oaipBypassCache

		dbmanager.StoreOpenAipBypassCache()
	}))

	oaipApiKeyLabel := widget.NewLabel("openAIP.net")
	oaipApiKeyInput := widget.NewEntryWithData(oaipApiKeyBinding)
	oaipApiKeyInput.Validator = nil
	oaipApiKeyInput.PlaceHolder = "SHARED API KEY"

	oaipApiKeyBinding.AddListener(binding.NewDataListener(func() {
		oaipApiKeyRaw, _ := oaipApiKeyBinding.Get()
		oaipApiKey := strings.TrimSpace(oaipApiKeyRaw)
		globals.OpenAipApiKey = oaipApiKey

		server.UpdateCacheApiKeys()

		logger.LogInfo("openAIP API key updated: [" + oaipApiKey + "]")

		dbmanager.StoreOpenAipApiKey()

		if oaipApiKey == "" || oaipApiKey == secrets.API_KEY_OPENAIP {
			//oaipBypassCacheBinding.Set(false)
			oaipBypassCacheCb.Disable()
		} else {
			oaipBypassCacheCb.Enable()
		}
	}))

	// bing maps api key
	bingApiKeyProLabel := widget.NewLabel("")
	bingApiKeyLabel := widget.NewLabel("Bing Maps")
	bingApiKeyInput := widget.NewEntryWithData(bingMapsApiKeyBinding)
	bingApiKeyInput.Validator = nil
	bingApiKeyInput.PlaceHolder = "Your Bing Maps API Key"

	bingMapsApiKeyBinding.AddListener(binding.NewDataListener(func() {
		bingApiKeyRaw, _ := bingMapsApiKeyBinding.Get()
		bingApiKey := strings.TrimSpace(bingApiKeyRaw)
		globals.BingMapsApiKey = bingApiKey

		logger.LogInfo("Bing Maps API key updated: [" + bingApiKey + "]")

		dbmanager.StoreBingMapsApiKey()
	}))

	if !globals.Pro {
		bingApiKeyInput.PlaceHolder = "Requires FSKneeboard PRO"
		//bingApiKeyProLabel.SetText("Requires FSKneeboard PRO")
		bingApiKeyInput.Disable()
		bingMapsApiKeyBinding.Set("")
	}

	// googleMaps maps api key
	// googleMapsApiKeyLabel := widget.NewLabel("Google Maps")
	// googleMapsApiKeyInput := widget.NewEntryWithData(googleMapsApiKeyBinding)
	// googleMapsApiKeyInput.Validator = nil
	// // googleMapsApiKeyInput.PlaceHolder = "SHARED API KEY"

	// googleMapsApiKeyBinding.AddListener(binding.NewDataListener(func() {
	// 	googleMapsApiKeyRaw, _ := googleMapsApiKeyBinding.Get()
	// 	googleMapsApiKey := strings.TrimSpace(googleMapsApiKeyRaw)
	// 	globals.GoogleMapsApiKey = googleMapsApiKey

	// 	logger.LogInfo("GoogleMaps API key updated: [" + googleMapsApiKey + "]")

	// 	dbmanager.StoreGoogleMapsApiKey()
	// }))

	apiKeysGrid := container.NewGridWithColumns(
		3,
		oaipApiKeyLabel, oaipApiKeyInput, oaipBypassCacheCb,
		bingApiKeyLabel, bingApiKeyInput, bingApiKeyProLabel,
		// googleMapsApiKeyLabel, googleMapsApiKeyInput, widget.NewLabel(""),
	)

	// LOG LEVEL
	// set log level
	loglevelLabel := widget.NewLabel("Log Level")
	loglevelWarningLabel := widget.NewLabel("WARNING: The Log Levels \"Debug\" and \"Silly\" may result in very large log files!")
	loglevelWarningLabel.Hidden = true
	loglevelWarningLabel.TextStyle.Italic = true
	// loglevelWarningLabel.Alignment = fyne.TextAlignCenter

	loglevelSelect := widget.NewSelect(loglevelOptions, func(selected string) {
		loglevelWarningLabel.Hidden = strings.ToLower(selected) != "debug" && strings.ToLower(selected) != "silly"
		loglevelBinding.Set(selected)
	})
	logsOpenFolderBtn := widget.NewButton("Open Log Folder", func() {
		logger.OpenLogFolder()
	})

	loglevelBinding.AddListener(binding.NewDataListener(func() {
		loglevelString, _ := loglevelBinding.Get()

		matchIndex := 0

		for index, value := range loglevelOptions {
			if strings.ToLower(loglevelString) == strings.ToLower(value) {
				matchIndex = index
				break
			}
		}

		if strings.ToLower(loglevelString) != strings.ToLower(loglevelSelect.Selected) {
			logger.LogDebug("loglevelBinding changed: [" + loglevelString + "]; updating ui select element...")
			loglevelSelect.SetSelected(loglevelOptions[matchIndex])
		} else {
			logger.LogDebug("loglevelBinding change listener: ui select element already up to date => [" + loglevelString + "]")
		}

		globals.LogLevel = strings.ToLower(loglevelString)

		dbmanager.StoreLogLevel()

		logger.SetLevel(loglevelString)
		logger.TryCreateLogFile()
	}))

	loglevelBinding.Set(globals.LogLevel)

	// general settings grid
	loggingGrid := container.NewGridWithColumns(
		3,
		loglevelLabel, loglevelSelect, logsOpenFolderBtn,
	)

	generalLabel := widget.NewLabel("General")
	generalLabel.TextStyle.Bold = true

	apiKeysLabel := widget.NewLabel("API Keys")
	apiKeysLabel.TextStyle.Bold = true

	apiKeysInfoLabel := widget.NewLabel("Obtain and add your own, private API keys below for better map performance.\nDepending on your internet connection, bypassing the local cache may also impact map performance.")
	apiKeysInfoLabel.TextStyle.Italic = true

	loggingLabel := widget.NewLabel("Logging")
	loggingLabel.TextStyle.Bold = true

	vBox := container.NewVBox(
		widget.NewLabel(""),
		generalLabel,
		widget.NewSeparator(),
		generalGrid,

		widget.NewLabel(""),
		apiKeysLabel,
		widget.NewSeparator(),
		apiKeysInfoLabel,
		apiKeysGrid,

		widget.NewLabel(""),
		loggingLabel,
		widget.NewSeparator(),
		loglevelWarningLabel,
		loggingGrid,
	)
	scroll := container.NewVScroll(vBox)
	maxContainer := container.NewMax(scroll)

	logger.LogDebug("Settings Panel initialized")

	return maxContainer
}
