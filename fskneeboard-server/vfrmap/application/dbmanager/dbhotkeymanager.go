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

// charts hotkey
func StoreChartsHotkeyShiftModifier() {
	DbWriteSettings("hotkeyChartsShiftModifier", strconv.FormatBool(globals.ChartsHotkey.ShiftKey))
}

func StoreChartsHotkeyCtrlModifier() {
	DbWriteSettings("hotkeyChartsCtrlModifier", strconv.FormatBool(globals.ChartsHotkey.CtrlKey))
}

func StoreChartsHotkeyAltModifier() {
	DbWriteSettings("hotkeyChartsAltModifier", strconv.FormatBool(globals.ChartsHotkey.AltKey))
}

func StoreChartsHotkeyKey() {
	DbWriteSettings("hotkeyChartsKey", SanitizeHotkey(globals.ChartsHotkey.Key))
}

func LoadChartsHotkey() {
	shiftModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyChartsShiftModifier"))
	ctrlModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyChartsCtrlModifier"))
	altModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyChartsAltModifier"))
	key := SanitizeHotkey(DbReadSettings("hotkeyChartsKey"))

	globals.ChartsHotkey = hotkeys.New(shiftModifier, ctrlModifier, altModifier, key)
}

// notepad hotkey
func StoreNotepadHotkeyShiftModifier() {
	DbWriteSettings("hotkeyNotepadShiftModifier", strconv.FormatBool(globals.NotepadHotkey.ShiftKey))
}

func StoreNotepadHotkeyCtrlModifier() {
	DbWriteSettings("hotkeyNotepadCtrlModifier", strconv.FormatBool(globals.NotepadHotkey.CtrlKey))
}

func StoreNotepadHotkeyAltModifier() {
	DbWriteSettings("hotkeyNotepadAltModifier", strconv.FormatBool(globals.NotepadHotkey.AltKey))
}

func StoreNotepadHotkeyKey() {
	DbWriteSettings("hotkeyNotepadKey", SanitizeHotkey(globals.NotepadHotkey.Key))
}

func LoadNotepadHotkey() {
	shiftModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyNotepadShiftModifier"))
	ctrlModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyNotepadCtrlModifier"))
	altModifier, _ := strconv.ParseBool(DbReadSettings("hotkeyNotepadAltModifier"))
	key := SanitizeHotkey(DbReadSettings("hotkeyNotepadKey"))

	globals.NotepadHotkey = hotkeys.New(shiftModifier, ctrlModifier, altModifier, key)
}