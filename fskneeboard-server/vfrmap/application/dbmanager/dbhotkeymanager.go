package dbmanager

import (
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
)

func SanitizeHotkey(key string) string {
	if len(key) != 1 {
		return ""
	}

	return strings.ToUpper(key)
}

func StoreMasterHotkeyShiftModifier() {
	DbWriteSettings("hotkeyMasterShiftModifier", strconv.FormatBool(globals.HotkeysMasterShiftModifier))
}

func StoreMasterHotkeyCtrlModifier() {
	DbWriteSettings("hotkeyMasterCtrlModifier", strconv.FormatBool(globals.HotkeysMasterCtrlModifier))
}

func StoreMasterHotkeyAltModifier() {
	DbWriteSettings("hotkeyMasterAltModifier", strconv.FormatBool(globals.HotkeysMasterAltModifier))
}

func StoreMasterHotkeyKey() {
	DbWriteSettings("hotkeyMasterKey", SanitizeHotkey(globals.HotkeysMasterKey))
}

func LoadMasterHotkey() {
	globals.HotkeysMasterShiftModifier, _ = strconv.ParseBool(DbReadSettings("hotkeyMasterShiftModifier"))
	globals.HotkeysMasterCtrlModifier, _ = strconv.ParseBool(DbReadSettings("hotkeyMasterCtrlModifier"))
	globals.HotkeysMasterAltModifier, _ = strconv.ParseBool(DbReadSettings("hotkeyMasterAltModifier"))
	globals.HotkeysMasterKey = SanitizeHotkey(DbReadSettings("hotkeyMasterKey"))
}