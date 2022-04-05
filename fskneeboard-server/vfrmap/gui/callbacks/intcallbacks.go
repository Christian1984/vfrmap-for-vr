package callbacks

func runIntCallback(value int, callback func(int)) {
	if callback != nil {
		go callback(value)
	}
}

func runIntCallbacks(value int, callbacks []func(int)) {
	for _, callback := range callbacks {
		go callback(value)
	}
}

var UpdateAutosaveStatusCallbacks []func(int)

func UpdateAutosaveStatus(interval int) {
	runIntCallbacks(interval, UpdateAutosaveStatusCallbacks)
}