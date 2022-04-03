package callbacks

func runBoolCallback(value bool, callback func(bool)) {
	if callback != nil {
		go callback(value)
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