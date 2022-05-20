package autosave

import (
	"vfrmap-for-vr/simconnect"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
)

func CreateAutosave(s *simconnect.SimConnect, savesToKeep int, verbose bool) {
	return
}

func OpenAutosaveFolder() {
	dialogs.ShowProFeatureInfo("Autosave")
}
