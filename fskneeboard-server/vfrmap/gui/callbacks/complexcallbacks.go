package callbacks

var UpdateServerStatusCallback func(string, string)

func UpdateServerStatus(statusMessage string, url string) {
	if UpdateServerStatusCallback != nil {
		UpdateServerStatusCallback(statusMessage, url)
	}
}
