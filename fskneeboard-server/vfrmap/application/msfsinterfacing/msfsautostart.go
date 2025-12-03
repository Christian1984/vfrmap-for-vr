package msfsinterfacing

import (
	"os/exec"
	"time"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/callbacks"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
)

func StartMsfs() {
	callbacks.UpdateMsfsStarted(true)

	failed := false

	switch globals.MsfsVersion {
	case "2020-steam":
		logger.LogInfoVerboseOverride("Starting Steam version of MSFS 2020...", false)
		utils.Println("\nStarting Flight Simulator 2020 via Steam... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", []string{"/C", "start", "steam://run/1250410"}...)
		fserr := cmd.Start()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Steam version of MSFS 2020 could not be started, details: "+fserr.Error(), false)
			utils.Println("\nFlight Simulator 2020 could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			failed = true
		}
	case "2020-winstore":
		logger.LogInfoVerboseOverride("Starting Windows Store version of MSFS 2020...", false)
		utils.Println("\nStarting Flight Simulator 2020... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", []string{"/C", "start", "shell:AppsFolder\\Microsoft.FlightSimulator_8wekyb3d8bbwe!App", "-FastLaunch"}...)
		fserr := cmd.Run()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Windows Store version of MSFS 2020 could not be started, details: "+fserr.Error(), false)
			utils.Println("\nWARNING: Flight Simulator 2020 could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			failed = true
		}
	case "2024-steam":
		logger.LogInfoVerboseOverride("Starting Steam version of MSFS 2024...", false)
		utils.Println("\nStarting Flight Simulator 2024 via Steam... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", []string{"/C", "start", "steam://run/2537590"}...)
		fserr := cmd.Start()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Steam version of MSFS 2024 could not be started, details: "+fserr.Error(), false)
			utils.Println("\nFlight Simulator 2024 could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			failed = true
		}
	case "2024-winstore":
		logger.LogInfoVerboseOverride("Starting Windows Store version of MSFS 2024...", false)
		utils.Println("\nStarting Flight Simulator 2024... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", []string{"/C", "start", "shell:AppsFolder\\Microsoft.Limitless_8wekyb3d8bbwe!App", "-FastLaunch"}...)
		fserr := cmd.Run()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Windows Store version of MSFS 2024 could not be started, details: "+fserr.Error(), false)
			utils.Println("\nWARNING: Flight Simulator 2024 could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")
			failed = true
		}
	default:
		logger.LogInfoVerboseOverride("MSFS autostart disabled!", false)
		logger.LogInfoVerboseOverride("MSFS version not properly configured!", false)
		utils.Println("\nMSFS version not configured in the settings. If you haven't already, please start Flight Simulator manually!")
		failed = true
	}

	if failed {
		callbacks.UpdateMsfsStarted(false)
		dialogs.ShowMsfsAutostartFailedError()
	} else {
		go func() {
			time.Sleep(30 * time.Second)
			callbacks.UpdateMsfsStarted(false)
		}()
	}
}
