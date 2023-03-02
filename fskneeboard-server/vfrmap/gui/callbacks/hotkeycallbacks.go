package callbacks

func runHotkeyCallback(shiftModifier bool, ctrlModifier bool, altModifier bool, key string, callback func(bool, bool, bool, string)) {
	if callback != nil {
		callback(shiftModifier, ctrlModifier, altModifier, key)
	}
}

var UpdateMasterHotkeyCallback func(bool, bool, bool, string)

func UpdateMasterHotkey(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	runHotkeyCallback(shiftModifier, ctrlModifier, altModifier, key, UpdateMasterHotkeyCallback)
}

var UpdateMapsHotkeyCallback func(bool, bool, bool, string)

func UpdateMapsHotkey(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	runHotkeyCallback(shiftModifier, ctrlModifier, altModifier, key, UpdateMapsHotkeyCallback)
}