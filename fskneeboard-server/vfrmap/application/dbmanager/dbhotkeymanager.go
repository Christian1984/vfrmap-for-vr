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