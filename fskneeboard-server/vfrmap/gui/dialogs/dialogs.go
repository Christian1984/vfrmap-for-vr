package dialogs

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var ParentWindow *fyne.Window

func ShowProgress(message string) *dialog.ProgressDialog {
	return dialog.NewProgress("Please wait...", message, *ParentWindow)
}

func ShowInformation(message string) {
	dialog.ShowInformation("Info", message, *ParentWindow)
}

func ShowError(message string) {
	dialog.ShowInformation("Something Went Wrong", message, *ParentWindow)
}

func ShowErrorAndExit(message string) {
	dialog.ShowConfirm("Something Went Wrong", message+"\nClick \"Yes\" to EXIT or \"No\" to CONTINUE (may result in an unstable experience)!", func(b bool) {
		if b {
			os.Exit(0)
		}
	}, *ParentWindow)
}

func ShowLicenseError() {
	dialog.ShowInformation("License Not Valid", "FSKneeboard could not find a valid license.", *ParentWindow)
}

func ShowProFeatureInfo(feature string) {
	dialog.ShowInformation("PRO Feature", "PLEASE NOTE: '"+feature+"' is a feature available exclusively to FSKneeboard PRO supporters.\n\nPlease consider supporting the development of FSKneeboard\nby purchasing a license at https://fskneeboard.com/buy-now/", *ParentWindow)
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

func ShowTourRestartedSuccessful() {
	dialog.ShowInformation("Tour Restarted", "The tutorial tour for the FSKneeboard ingame panel was restarted.\nPlease close and re-open the ingame panel to take the tour.", *ParentWindow)
}

func ShowImporterDownloadPrompt(showFreeUserInfo bool, callback func(bool)) {
	freeUserInfo := ""

	if showFreeUserInfo {
		freeUserInfo = "\n\n==============================\nFREE USERS, PLEASE NOTE:\n\nThe charts imported with the PDF Importer Tool are supposed to be viewed with the\nFSKneeboard Charts Viewer, which is only available to PRO supporters.\nFREE users may download and use the tool (and the source code),\nbut cannot view the imported charts inside FSKneeboard!\n==============================\n"
	}

	dialog.ShowConfirm("PDF Importer", "FSKneeboard's PDF Importer Tool is a separate application that needs to be download first."+freeUserInfo+"\nDo you want FSKneeboard to download the importer tool and continue?\n\n(Alternatively, you may refer to the FSKneeboard documentation to install the FSKneeboard Importer Tool manually.)", func(b bool) {
		callback(b)
	}, *ParentWindow)
}
