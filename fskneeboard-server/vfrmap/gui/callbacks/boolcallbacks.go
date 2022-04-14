package callbacks

func runBoolCallback(value bool, callback func(bool)) {
	if callback != nil {
		callback(value)
	}
}

var UpdateServerStartedCallback func(bool)

func UpdateServerStarted(status bool) {
	runBoolCallback(status, UpdateServerStartedCallback)
}

var UpdateMsfsStartedCallback func(bool)

func UpdateMsfsStarted(status bool) {
	runBoolCallback(status, UpdateMsfsStartedCallback)
}

var NewVersionAvailableCallback func(bool)

func NewVersionAvailable(status bool) {
	runBoolCallback(status, NewVersionAvailableCallback)
}

var MsfsVersionChangedCallback func(bool)

func MsfsVersionChanged(steam bool) {
	runBoolCallback(steam, MsfsVersionChangedCallback)
}

var MsfsAutostartChangedCallback func(bool)

func MsfsAutostartChanged(autostart bool) {
	runBoolCallback(autostart, MsfsAutostartChangedCallback)
}

/*var ServerAutostartChangedCallback func(bool)

func ServerAutostartChanged(autostart bool) {
	runBoolCallback(autostart, ServerAutostartChangedCallback)
}*/