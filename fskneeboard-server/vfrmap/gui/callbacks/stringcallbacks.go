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

var MsfsVersionChangedStringCallback func(string)

func MsfsVersionChangedString(version string) {
	runStringCallback(version, MsfsVersionChangedStringCallback)
}

var UpdateOpenAipApiCallback func(string)

func UpdateOpenAipApi(apiKey string) {
	runStringCallback(apiKey, UpdateOpenAipApiCallback)
}

var UpdateBingMapsApiCallback func(string)

func UpdateBingMapsApi(apiKey string) {
	runStringCallback(apiKey, UpdateBingMapsApiCallback)
}

var UpdateGoogleMapsApiCallback func(string)

func UpdateGoogleMapsApi(apiKey string) {
	runStringCallback(apiKey, UpdateGoogleMapsApiCallback)
}
