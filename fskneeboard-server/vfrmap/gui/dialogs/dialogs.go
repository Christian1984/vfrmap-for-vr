package dialogs

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var ParentWindow *fyne.Window

func ShowError(message string) {
	dialog.ShowInformation("Something Went Wrong", message, *ParentWindow)
}

func ShowErrorAndExit(message string) {
	dialog.ShowConfirm("Something Went Wrong", message + "\nClick \"Yes\" to EXIT or \"No\" to CONTINUE (may result in an unstable experience)!", func(b bool) {
		if b {
			os.Exit(0)
		}
	}, *ParentWindow)
}

func ShowLicenseError() {
	dialog.ShowInformation("License Not Valid", "FSKneeboard could not find a valid license.", *ParentWindow)
}

func ShowProFeatureInfo(feature string) {
	dialog.ShowInformation("PRO Feature", "PLEASE NOTE: '" + feature + "' is a feature available exclusively to FSKneeboard PRO supporters.\n\nPlease consider supporting the development of FSKneeboard\nby purchasing a license at https://fskneeboard.com/buy-now/", *ParentWindow)
}

func ShowMsfsAutostartFailedError() {
	dialog.ShowInformation("Failed to Start Flight Simulator", "Flight Simulator could not be started. Please see the console output for further details.", *ParentWindow)
}

func ShowMsfsShutdownInfoAndExit() {
	dialog.ShowConfirm("Flight Simulator Shutdown", "Flight Simulator was closed.\nDo you want to close FSKneeboard now?\nPLEASE NOTE: FSKneeboard has to be restarted for each Flight Simulator session!", func(b bool) {
		if b {
			os.Exit(0)
		}
	}, *ParentWindow)
}