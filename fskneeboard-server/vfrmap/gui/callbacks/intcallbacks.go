package callbacks

func runIntCallback(value int, callback func(int)) {
	if callback != nil {
		callback(value)
	}
}

func runIntCallbacks(value int, callbacks []func(int)) {
	for _, callback := range callbacks {
		callback(value)
	}
}

var UpdateAutosaveStatusCallbacks []func(int)

func UpdateAutosaveStatus(interval int) {
	runIntCallbacks(interval, UpdateAutosaveStatusCallbacks)
}