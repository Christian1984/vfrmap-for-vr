package main

// build: GOOS=windows GOARCH=amd64 go build -o fskneeboard.exe vfrmap-for-vr/vfrmap

import (
	"flag"
	"os/exec"
	"strconv"

	"vfrmap-for-vr/_vendor/premium/common"
	"vfrmap-for-vr/_vendor/premium/drm"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server"
	"vfrmap-for-vr/vfrmap/utils"

	updatechecker "github.com/Christian1984/go-update-checker"
)

var buildVersion string
var buildTime string
var pro string

var disableTeleport bool
var steamfs bool
var winstorefs bool
var noupdatecheck bool

var logLevel string

func initFsk() {
	globals.Pro = pro == "true"

	globals.ProductName = "FSKneeboard"
	if globals.Pro {
		globals.ProductName += " PRO"
	}

	utils.Printf("\n"+globals.ProductName+" - Server\n  Website: https://fskneeboard.com\n  Discord: https://discord.fskneeboard.com\n  Readme:  https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md\n  Issues:  https://github.com/Christian1984/vfrmap-for-vr/issues\n  Version: %s (%s)\n\n", buildVersion, buildTime)

	if globals.Pro {
		logger.LogInfo("FSKneeboard PRO started. Checking license information...", false)

		utils.Println("=== INFO: License")		
		drmData := drm.New()
		if !drmData.Valid() {
			utils.Println("\nWARNING: You do not have a valid license to run FSKneeboard PRO!")
			utils.Println("Please purchase a license at https://fskneeboard.com/buy-now and place your fskneeboard.lic-file in the same directory as fskneeboard.exe.")
			logger.LogWarn("No valid license found, details: email [" + drmData.Email() + "] - Shutting down!", false)
			server.ShutdownWithPrompt()
		} else {
			utils.Println("Valid license found! This copy of FSKneeboard is licensed to: " + drmData.Email())
			utils.Println("Thanks for purchasing FSKneeboard PRO and supporting the development of this mod!")
			utils.Println("")

			logger.LogInfo("Valid license found, details: email [" + drmData.Email() + "]", false)
		}
	} else {
		logger.LogInfo("FSKneeboard FREE started...", false)

		utils.Println("=== INFO: How to Support the Development of FSKneeboard")
		utils.Println("Thanks for trying FSKneeboard FREE!")
		utils.Println("Please checkout https://fskneeboard.com and purchase FSKneeboard PRO to unlock all features the extension has to offer.")
		utils.Println("")
	}

	if !noupdatecheck {
		logger.LogInfo("Running Update-Check...", false)

		uc := updatechecker.New("Christian1984", "vfrmap-for-vr", "FSKneeboard", common.DOWNLOAD_LINK, 3, false)
		uc.CheckForUpdate(buildVersion)

		if uc.UpdateAvailable {
			logger.LogInfo("New Version found:\n" + uc.Message, false)
			
			utils.Println(uc.Message)
			utils.Println("")
		} else {
			logger.LogInfo("Could not find a new version!", false)
		}
	}

	// autosave info
	utils.Println("=== INFO: Autosave")

	if globals.AutosaveInterval > 0 {
		utils.Printf("Autosave Interval set to %d minute(s)...\n", globals.AutosaveInterval)
		logger.LogInfo("Autosave Interval set to " + strconv.Itoa(globals.AutosaveInterval) + " minute(s)", false)
	} else {
		utils.Println("Autosave not activated. Run fskneeboard.exe --autosave 5 to automatically save your flights every 5 minutes...")
		logger.LogInfo("Autosave not activated", false)
	}

	if globals.Pro {
		utils.Println("PLEASE NOTE: 'Autosave' is a feature available exclusively to FSKneeboard PRO supporters. Please consider supporting the development of FSKneeboard by purchasing a license at https://fskneeboard.com/buy-now/")
	}

	utils.Println("")

	// hotkey info
	utils.Println("=== INFO: Hotkey")

	if globals.Hotkey != 0 {
		key := "F"
		mod := "[ALT]"

		switch globals.Hotkey {
		case 2:
			key = "K"
		case 3:
			key = "T"
		case 4:
			key = "F"
			mod = "[CTRL]+[SHIFT]"
		case 5:
			key = "K"
			mod = "[CTRL]+[SHIFT]"
		case 6:
			key = "T"
			mod = "[CTRL]+[SHIFT]"
		}

		logger.LogInfo("Hotkey set to " + mod + "+" + key, false)
		utils.Println("Hotkey set to " + mod + "+" + key)
	} else {
		utils.Println("Hotkey not configured. Run fskneeboard.exe --hotkey 1 to enable [ALT]+F as your hotkey to toggle the ingame panel's visibility. Please refer to the readme for other hotkey options.")
	}

	utils.Println("")

	// connect to bolt db
	utils.Println("=== INFO: Local FSKneeboard Database Connection")
	db_err := dbmanager.DbConnect()
	if db_err != nil {
		utils.Println("")

		logger.LogError("WARNING: Cannot connect to local FSKneeboard database. Please make sure that there's no other instance of FSKneeboard running! Shutting down...", true)
		server.ShutdownWithPrompt()
		} else {
		logger.LogInfo("Established connection with local FSKneeboard database!", false)

		utils.Println("Established connection with local FSKneeboard database!")
		dbmanager.DbInit()

		utils.Println("")
	}

	// starting Flight Simulator
	utils.Println("=== INFO: Flight Simulator Autostart")

	if steamfs {
		logger.LogInfo("Starting Steam version of MSFS...", false)
		utils.Println("Starting Flight Simulator via Steam... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start steam://run/1250410")
		fserr := cmd.Start()
		if fserr != nil {
			logger.LogWarn("Steam version of MSFS could not be started, details: " + fserr.Error(), false)
			utils.Println("Flight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
		}
	} else if winstorefs {
		logger.LogInfo("Starting Windows Store version of MSFS...", false)
		utils.Println("Starting Flight Simulator... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start shell:AppsFolder\\Microsoft.FlightSimulator_8wekyb3d8bbwe!App -FastLaunch")
		fserr := cmd.Run()
		if fserr != nil {
			logger.LogWarn("Windows Store version of MSFS could not be started, details: " + fserr.Error(), false)
			utils.Println("WARNING: Flight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			utils.Println("IMPORTANT: If you have purchased MSFS on Steam, please run 'fskneeboard.exe --steamfs' as described in the manual under 'Usage'!")
		}
	} else {
		logger.LogInfo("MSFS autostart disabled!", false)
		utils.Println("FSKneeboard started without autostart options --steamfs or --winstorefs.")
		utils.Println("If you haven't already, please start Flight Simulator manually!")
	}
}

func main() {
	gui.InitGui()

	flag.BoolVar(&globals.Verbose, "verbose", false, "verbose output")
	flag.StringVar(&globals.HttpListen, "listen", "0.0.0.0:9000", "http listen")
	flag.StringVar(&logLevel, "log", "off", "set log level (debug | info | error | off)")
	flag.BoolVar(&globals.DisableTeleport, "disable-teleport", false, "disable teleport")
	flag.BoolVar(&globals.DevMode, "dev", false, "enable dev mode, i.e. no running msfs required")
	flag.BoolVar(&steamfs, "steamfs", false, "start Flight Simulator via Steam")
	flag.BoolVar(&winstorefs, "winstorefs", false, "start Flight Simulator via Windows Store")
	flag.BoolVar(&noupdatecheck, "noupdatecheck", false, "prevent FSKneeboard from checking the GitHub API for updates")
	flag.BoolVar(&globals.Quietshutdown, "quietshutdown", false, "prevent FSKneeboard from showing a \"Press ENTER to continue...\" prompt after disconnecting from MSFS")

	flag.IntVar(&globals.AutosaveInterval, "autosave", 0, "set autosave interval in minutes")
	flag.IntVar(&globals.Hotkey, "hotkey", 0, "select a hotkey to toggle the ingame panel's visibility. 1 => [ALT]+F, 2 => [ALT]+K, 3 => [ALT]+T, 4 => [CTRL]+[SHIFT]+F, 5 => [CTRL]+[SHIFT]+K, 6 => [CTRL]+[SHIFT]+T")

	flag.Parse()

	logger.Init(logLevel, globals.Verbose)

	if logger.ShouldLog(logLevel) {
		logger.CreateLogFile()
		logger.LogDebug("Logfile created!", false)
	}
	
	/*
	logger.LogMessage("OFF-Test", logger.Off, "", false)
	logger.LogDebug("DEBUG-Test", false)
	logger.LogInfo("INFO-Test", false)
	logger.LogWarn("WARN-Test", false)
	logger.LogError("ERROR-Test", false)
	*/


	logger.LogInfo("FSKneeboard started with params\n" + 
		"\tverbose:          " + strconv.FormatBool(globals.Verbose) + "\n" +
		"\tlisten:           " + globals.HttpListen + "\n" +
		"\tlog:              " + logLevel + "\n" +
		"\tdisable-teleport: " + strconv.FormatBool(disableTeleport) + "\n" +
		"\tdev:              " + strconv.FormatBool(globals.DevMode) + "\n" +
		"\tsteamfs:          " + strconv.FormatBool(steamfs) + "\n" +
		"\twinstorefs:       " + strconv.FormatBool(winstorefs) + "\n" +
		"\tnoupdatecheck:    " + strconv.FormatBool(noupdatecheck) + "\n" +
		"\tquietshutdown:    " + strconv.FormatBool(globals.Quietshutdown) + "\n" +
		"\tautosave:         " + strconv.Itoa(globals.AutosaveInterval) + "\n" +
		"\thotkey:           " + strconv.Itoa(globals.Hotkey) + "\n", false)

	go server.StartFskServer()
	initFsk()
	gui.ShowAndRun()
}
