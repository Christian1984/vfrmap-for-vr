package callbacks

func runStringCallback(value string, callback func(string)) {
	if callback != nil {
		go callback(value)
	}
}

var UpdateServerStatusCallback func(string)

func UpdateServerStatus(status string) {
	runStringCallback(status, UpdateServerStatusCallback)
}

var UpdateMsfsConnectionStatusCallback func(string)

func UpdateMsfsConnectionStatus(status string) {
	runStringCallback(status, UpdateMsfsConnectionStatusCallback)
}

var UpdateLicenseStatusCallback func(string)

func UpdateLicenseStatus(status string) {
	runStringCallback(status, UpdateLicenseStatusCallback)
}