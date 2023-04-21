package callbacks

func runFloat64Callback(value float64, callback func(float64)) {
	if callback != nil {
		callback(value)
	}
}

func runFloat64Callbacks(value float64, callback func(float64)) {
	callback(value)
}

var UpdateInterfaceScaleCallback func(float64)

func UpdateInterfaceScale(float64erval float64) {
	runFloat64Callbacks(float64erval, UpdateInterfaceScaleCallback)
}