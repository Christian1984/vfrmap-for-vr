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

	if globals.SteamFs {
		logger.LogInfoVerboseOverride("Starting Steam version of MSFS...", false)
		utils.Println("\nStarting Flight Simulator via Steam... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start steam://run/1250410")
		fserr := cmd.Start()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Steam version of MSFS could not be started, details: "+fserr.Error(), false)
			utils.Println("\nFlight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")

			failed = true
		}
	} else if globals.WinstoreFs {
		logger.LogInfoVerboseOverride("Starting Windows Store version of MSFS...", false)
		utils.Println("\nStarting Flight Simulator... Just sit tight :-)")
		cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C start shell:AppsFolder\\Microsoft.FlightSimulator_8wekyb3d8bbwe!App -FastLaunch")
		fserr := cmd.Run()
		if fserr != nil {
			logger.LogWarnVerboseOverride("Windows Store version of MSFS could not be started, details: "+fserr.Error(), false)
			utils.Println("\nWARNING: Flight Simulator could not be started. Please start Flight Simulator manually! (" + fserr.Error() + ")")

			failed = true
		}
	} else {
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
