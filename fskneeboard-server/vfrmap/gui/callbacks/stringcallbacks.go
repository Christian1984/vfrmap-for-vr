package callbacks

func runStringCallback(value string, callback func(string)) {
	if callback != nil {
		callback(value)
	}
}

var UpdateMsfsConnectionStatusCallback func(string)

func UpdateMsfsConnectionStatus(status string) {
	runStringCallback(status, UpdateMsfsConnectionStatusCallback)
}

var UpdateLicenseStatusCallback func(string)

func UpdateLicenseStatus(status string) {
	runStringCallback(status, UpdateLicenseStatusCallback)
}

var UpdateLogLevelStatusCallback func(string)

func UpdateLogLevelStatus(status string) {
	runStringCallback(status, UpdateLogLevelStatusCallback)
}
