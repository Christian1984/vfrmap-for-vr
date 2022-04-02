package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var ParentWindow *fyne.Window

func ShowLicenseError() {
	dialog.ShowInformation("License Not Valid", "FSKneeboard could not find a valid license.", *ParentWindow)
}