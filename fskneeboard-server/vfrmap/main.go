package main

// build: GOOS=windows GOARCH=amd64 go build -o fskneeboard.exe vfrmap-for-vr/vfrmap

import (
	"flag"
	"strconv"

	"vfrmap-for-vr/_vendor/premium/common"
	"vfrmap-for-vr/_vendor/premium/drm"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/msfsinterfacing"
	"vfrmap-for-vr/vfrmap/gui"
	"vfrmap-for-vr/vfrmap/gui/callbacks"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/gui/tabs/consolepanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/controlpanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/hotkeyspanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/settingspanel"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server"
	"vfrmap-for-vr/vfrmap/utils"

	updatechecker "github.com/Christian1984/go-update-checker"
)

var buildVersion string
var buildTime string
var pro string

var noupdatecheck bool

func initFsk() {
	utils.Println("Initializing FSKneeboard Core Application...")

	utils.Printf("\n"+globals.ProductName+" - Server\n  Website: https://fskneeboard.com\n  Discord: https://discord.fskneeboard.com\n  Readme:  https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md\n  Issues:  https://github.com/Christian1984/vfrmap-for-vr/issues\n  Version: %s (%s)\n\n", buildVersion, buildTime)

	callbacks.NewVersionAvailable(false)

	if !noupdatecheck {
		logger.LogInfo("Running Update-Check...", false)

		uc := updatechecker.New("Christian1984", "vfrmap-for-vr", "FSKneeboard", common.DOWNLOAD_LINK, 3, false)
		uc.CheckForUpdate(buildVersion)

		if uc.UpdateAvailable {
			callbacks.NewVersionAvailable(true)

			logger.LogInfo("New Version found:\n" + uc.Message, false)
			
			utils.Println(uc.Message)
			utils.Println("")
		} else {
			logger.LogInfo("Could not find a new version!", false)
		}
	} else {
		callbacks.NewVersionAvailable(false)
	}

	// connect to bolt db
	utils.Println("=== INFO: Local FSKneeboard Database Connection")
	db_err := dbmanager.DbConnect()
	if db_err != nil {
		utils.Println("")

		logger.LogError("WARNING: Cannot connect to local FSKneeboard database. Please make sure that there's no other instance of FSKneeboard running! Shutting down...", true)
		dialogs.ShowErrorAndExit("Cannot connect to local FSKneeboard database. Please make sure that there's no other instance of FSKneeboard running!")
		} else {
		logger.LogInfo("Established connection with local FSKneeboard database!", false)

		utils.Println("Established connection with local FSKneeboard database!")
		dbmanager.DbInit()

		utils.Println("")
	}

	// load settings
	// loglevel
	if globals.LogLevel != "off" {
		dbmanager.StoreLogLevel() // store loglevel if set through args
	} else {
		globals.LogLevel = dbmanager.LoadLogLevel() // only load if off so far
	}

	callbacks.UpdateLogLevelStatus(globals.LogLevel)

	// autosave interval
	dbmanager.LoadAutosaveInterval()
	callbacks.UpdateAutosaveStatus(globals.AutosaveInterval)

	// msfs version
	dbmanager.LoadMsfsVersion()
	callbacks.MsfsVersionChanged(globals.SteamFs)

	// msfs autostart
	dbmanager.LoadMsfsAutostart()
	callbacks.MsfsAutostartChanged(globals.MsfsAutostart)

	// load hotkeys
	// master hotkey
	dbmanager.LoadMasterHotkey()
	callbacks.UpdateMasterHotkey(globals.MasterHotkey.ShiftKey, globals.MasterHotkey.CtrlKey, globals.MasterHotkey.AltKey, globals.MasterHotkey.Key)

	// check license
	if globals.Pro {
		logger.LogInfo("FSKneeboard PRO started. Checking license information...", false)

		utils.Println("=== INFO: License")
		drmData := drm.New()
		globals.DrmValid = drmData.Valid()

		if !globals.DrmValid {
			utils.Println("WARNING: You do not have a valid license to run FSKneeboard PRO!")
			utils.Println("Please purchase a license at https://fskneeboard.com/buy-now and place your fskneeboard.lic-file in the same directory as fskneeboard.exe.")

			logger.LogWarn("No valid license found, details: email [" + drmData.Email() + "]", false)

			callbacks.UpdateLicenseStatus("Invalid")
			dialogs.ShowLicenseError()

			return
		} else {
			utils.Println("Valid license found! This copy of FSKneeboard is licensed to: " + drmData.Email())
			utils.Println("Thanks for purchasing FSKneeboard PRO and supporting the development of this mod!")
			utils.Println("")

			logger.LogInfo("Valid license found, details: email [" + drmData.Email() + "]", false)
			callbacks.UpdateLicenseStatus("Valid")
		}
	} else {
		logger.LogInfo("FSKneeboard FREE started...", false)

		utils.Println("=== INFO: How to Support the Development of FSKneeboard")
		utils.Println("Thanks for trying FSKneeboard FREE!")
		utils.Println("Please checkout https://fskneeboard.com and purchase FSKneeboard PRO to unlock all features the extension has to offer.")
		utils.Println("")

		callbacks.UpdateLicenseStatus("TRIAL (FSKneeboard FREE)")
	}

	// starting Flight Simulator
	utils.Println("=== INFO: Flight Simulator Autostart")
	if (globals.MsfsAutostart) {
		msfsinterfacing.StartMsfs()
	} else {
		logger.LogInfo("MSFS autostart disabled!", false)
		utils.Println("MSFS autostart disabled! Please configure your version of Flight Simulator and enable autostart in the settings section.")
		utils.Println("")
	}

	callbacks.ProVersionChanged(globals.Pro)

	// starting FSKneeboard Server
	go server.StartFskServer()
}

func registerGuiCallbacks() {
	utils.GuiPrintCallback = consolepanel.ConsoleLog

	callbacks.UpdateServerStatusCallback = controlpanel.UpdateServerStatus
	callbacks.UpdateMsfsConnectionStatusCallback = controlpanel.UpdateMsfsConnectionStatus
	callbacks.UpdateLicenseStatusCallback = controlpanel.UpdateLicenseStatus
	callbacks.ProVersionChangedCallback = controlpanel.UpdateProVersionStatus

	callbacks.UpdateLogLevelStatusCallback = settingspanel.UpdateLogLevelStatus

	callbacks.UpdateAutosaveStatusCallbacks = append(callbacks.UpdateAutosaveStatusCallbacks, controlpanel.UpdateAutosaveStatus)
	callbacks.UpdateAutosaveStatusCallbacks = append(callbacks.UpdateAutosaveStatusCallbacks, settingspanel.UpdateAutosaveStatus)

	callbacks.UpdateMsfsStartedCallback = controlpanel.UpdateMsfsStarted
	callbacks.NewVersionAvailableCallback = controlpanel.UpdateNewVersionAvailable

	callbacks.MsfsVersionChangedCallback = settingspanel.UpdateMsfsVersionStatus
	callbacks.MsfsAutostartChangedCallback = settingspanel.UpdateMsfsAutostartStatus

	callbacks.UpdateMasterHotkeyCallback = hotkeyspanel.UpdateMasterHotkeyStatus
}

func main() {
	globals.Pro = pro == "true"

	globals.ProductName = "FSKneeboard"
	if globals.Pro {
		globals.ProductName += " PRO"
		globals.DownloadLink = globals.DownloadLinkPro
	} else {
		globals.ProductName += " FREE"
		globals.DownloadLink = globals.DownloadLinkFree
	}

	// flags to respect always
	flag.BoolVar(&globals.DevMode, "dev", false, "enable dev mode, i.e. no running msfs required")
	flag.StringVar(&globals.HttpListen, "listen", "0.0.0.0:9000", "http listen")
	flag.BoolVar(&noupdatecheck, "noupdatecheck", false, "prevent FSKneeboard from checking the GitHub API for updates")
	flag.BoolVar(&globals.Verbose, "verbose", false, "verbose output")

	// flags to compare against stored values
	flag.StringVar(&globals.LogLevel, "log", "off", "set log level (debug | info | error)")

	// TODO: flags to check if required
	flag.BoolVar(&globals.Quietshutdown, "quietshutdown", false, "prevent FSKneeboard from showing a \"Press ENTER to continue...\" prompt after disconnecting from MSFS")

	flag.Parse()

	// init logger
	logger.Init(globals.LogLevel, globals.Verbose)
	logger.TryCreateLogFile()
	
	/*
	logger.LogMessage("OFF-Test", logger.Off, "", false)
	logger.LogDebug("DEBUG-Test", false)
	logger.LogInfo("INFO-Test", false)
	logger.LogWarn("WARN-Test", false)
	logger.LogError("ERROR-Test", false)
	*/


	logger.LogInfo("FSKneeboard started with params\n" + 
		"\tverbose:          " + strconv.FormatBool(globals.Verbose) + "\n" +
		"\tlog:              " + globals.LogLevel + "\n" +
		"\tlisten:           " + globals.HttpListen + "\n" +
		"\tdev:              " + strconv.FormatBool(globals.DevMode) + "\n" +
		"\tnoupdatecheck:    " + strconv.FormatBool(noupdatecheck) + "\n" +
		"\tquietshutdown:    " + strconv.FormatBool(globals.Quietshutdown) + "\n", false)

	gui.InitGui()
	registerGuiCallbacks()

	initFsk()

	gui.ShowAndRun()
}
