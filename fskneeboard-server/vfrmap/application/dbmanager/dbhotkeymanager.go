package dbmanager

import (
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/hotkeys"
)

func SanitizeHotkey(key string) string {
	if len(key) != 1 {
		return ""
	}

	return strings.ToLower(key)
}

// master hotkey
func StoreMasterHotkeyShiftModifier() {
	DbWriteSettings("hotkeyMasterShiftModifier", strconv.FormatBool(globals.MasterHotkey.ShiftKey))
}

func StoreMasterHotkeyCtrlModifier() {
	DbWriteSettings("hotkeyMasterCtrlModifier", strconv.FormatBool(globals.MasterHotkey.CtrlKey))
}

func StoreMasterHotkeyAltModifier() {
	DbWriteSettings("hotkeyMasterAltModifier", strconv.FormatBool(globals.MasterHotkey.AltKey))
}

func StoreMasterHotkeyKey() {
	DbWriteSettings("hotkeyMasterKey", SanitizeHotkey(globals.MasterHotkey.Key))
}

func LoadMasterHotkey() {
	shiftModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMasterShiftModifier"))
	ctrlModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMasterCtrlModifier"))
	altModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMasterAltModifier"))
	key := SanitizeHotkey(DbReadSettings("hotkeyMasterKey"))

	globals.MasterHotkey = hotkeys.New(shiftModifier, ctrlModifier, altModifier, key)
}

// maps hotkey
func StoreMapsHotkeyShiftModifier() {
	DbWriteSettings("hotkeyMapsShiftModifier", strconv.FormatBool(globals.MapsHotkey.ShiftKey))
}

func StoreMapsHotkeyCtrlModifier() {
	DbWriteSettings("hotkeyMapsCtrlModifier", strconv.FormatBool(globals.MapsHotkey.CtrlKey))
}

func StoreMapsHotkeyAltModifier() {
	DbWriteSettings("hotkeyMapsAltModifier", strconv.FormatBool(globals.MapsHotkey.AltKey))
}

func StoreMapsHotkeyKey() {
	DbWriteSettings("hotkeyMapsKey", SanitizeHotkey(globals.MapsHotkey.Key))
}

func LoadMapsHotkey() {
	shiftModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMapsShiftModifier"))
	ctrlModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMapsCtrlModifier"))
	altModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyMapsAltModifier"))
	key := SanitizeHotkey(DbReadSettings("hotkeyMapsKey"))

	globals.MapsHotkey = hotkeys.New(shiftModifier, ctrlModifier, altModifier, key)
}