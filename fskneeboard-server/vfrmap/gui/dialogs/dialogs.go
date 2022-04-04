package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var ParentWindow *fyne.Window

func ShowLicenseError() {
	dialog.ShowInformation("License Not Valid", "FSKneeboard could not find a valid license.", *ParentWindow)
}

func ShowProFeatureInfo(feature string) {
	dialog.ShowInformation("PRO Feature", "PLEASE NOTE: '" + feature + "' is a feature available exclusively to FSKneeboard PRO supporters.\n\nPlease consider supporting the development of FSKneeboard\nby purchasing a license at https://fskneeboard.com/buy-now/", *ParentWindow)
}

func ShowMsfsAutostartFailedError() {
	dialog.ShowInformation("Failed to Start Flight Simulator", "Flight Simulator could not be started. Please see the console output for further details.", *ParentWindow)
}