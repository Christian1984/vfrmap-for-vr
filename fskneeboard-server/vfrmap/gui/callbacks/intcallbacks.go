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

var UpdateAutosaveCallbacks []func(int)

func UpdateAutosave(interval int) {
	runIntCallbacks(interval, UpdateAutosaveCallbacks)
}