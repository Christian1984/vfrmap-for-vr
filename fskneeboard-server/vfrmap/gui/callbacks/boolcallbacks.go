package callbacks

func runBoolCallback(value bool, callback func(bool)) {
	if callback != nil {
		callback(value)
	}
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

var ShowGuiTourChangedCallback func(bool)

func ShowGuiTourChanged(show bool) {
	runBoolCallback(show, ShowGuiTourChangedCallback)
}

var OpenAipBypassCacheChangedCallback func(bool)

func OpenAipBypassCacheChanged(show bool) {
	runBoolCallback(show, OpenAipBypassCacheChangedCallback)
}
