package hotkeyspanel

import (
	"strings"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/server/hotkeys"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var keyOptions = []string{
	"[off]", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var masterShiftModifierBinding = binding.NewBool()
var masterCtrlModifierBinding = binding.NewBool()
var masterAltModifierBinding = binding.NewBool()
var masterKeyBinding = binding.NewString()

var mapsShiftModifierBinding = binding.NewBool()
var mapsCtrlModifierBinding = binding.NewBool()
var mapsAltModifierBinding = binding.NewBool()
var mapsKeyBinding = binding.NewString()

var chartsShiftModifierBinding = binding.NewBool()
var chartsCtrlModifierBinding = binding.NewBool()
var chartsAltModifierBinding = binding.NewBool()
var chartsKeyBinding = binding.NewString()

var notepadShiftModifierBinding = binding.NewBool()
var notepadCtrlModifierBinding = binding.NewBool()
var notepadAltModifierBinding = binding.NewBool()
var notepadKeyBinding = binding.NewString()

func updateHotkeyStatus(shiftModifier bool, shiftBinding *binding.Bool,
	ctrlModifier bool, ctrlBinding *binding.Bool,
	altModifier bool, altBinding *binding.Bool,
	key string, keyBinding *binding.String) {
	if len(key) != 1 {
		(*keyBinding).Set(keyOptions[0])
	} else {
		(*keyBinding).Set(strings.ToUpper(key))
	}

	(*shiftBinding).Set(shiftModifier)
	(*ctrlBinding).Set(ctrlModifier)
	(*altBinding).Set(altModifier)
}

func UpdateMasterHotkeyStatus(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	updateHotkeyStatus(
		shiftModifier, &masterShiftModifierBinding,
		ctrlModifier, &masterCtrlModifierBinding,
		altModifier, &masterAltModifierBinding,
		key, &masterKeyBinding,
	)
}

func UpdateMapsHotkeyStatus(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	updateHotkeyStatus(
		shiftModifier, &mapsShiftModifierBinding,
		ctrlModifier, &mapsCtrlModifierBinding,
		altModifier, &mapsAltModifierBinding,
		key, &mapsKeyBinding,
	)
}

func UpdateChartsHotkeyStatus(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	updateHotkeyStatus(
		shiftModifier, &chartsShiftModifierBinding,
		ctrlModifier, &chartsCtrlModifierBinding,
		altModifier, &chartsAltModifierBinding,
		key, &chartsKeyBinding,
	)
}

func UpdateNotepadHotkeyStatus(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) {
	updateHotkeyStatus(
		shiftModifier, &notepadShiftModifierBinding,
		ctrlModifier, &notepadCtrlModifierBinding,
		altModifier, &notepadAltModifierBinding,
		key, &notepadKeyBinding,
	)
}

func HotkeysPanel() *fyne.Container {
	logger.LogDebug("Initializing Hotkeys Panel...")

	// grid and centerContainer
	labelNotes := widget.NewLabel("IMPORTANT NOTES:\n" +
		"- In order for the hotkeys to work, you have to manually open the FSKneeboard panel at least once per flight!\n" +
		"- Use your HOTAS software to map these hotkeys to your HOTAS!\n")

	// master switch
	masterLabel := widget.NewLabel("Toggle FSKneeboard Panel")

	masterShiftCb := widget.NewCheckWithData("Shift", masterShiftModifierBinding)
	masterShiftModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterShiftModifierBinding.Get()
		globals.MasterHotkey.ShiftKey = value
		dbmanager.StoreMasterHotkeyShiftModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	masterCtrlCb := widget.NewCheckWithData("Ctrl", masterCtrlModifierBinding)
	masterCtrlModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterCtrlModifierBinding.Get()
		globals.MasterHotkey.CtrlKey = value
		dbmanager.StoreMasterHotkeyCtrlModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	masterAltCb := widget.NewCheckWithData("Alt", masterAltModifierBinding)
	masterAltModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterAltModifierBinding.Get()
		globals.MasterHotkey.AltKey = value
		dbmanager.StoreMasterHotkeyAltModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	masterHotkey := widget.NewSelect(keyOptions, func(s string) {
		masterKeyBinding.Set(strings.ToLower(s))
	})
	masterKeyBinding.AddListener(binding.NewDataListener(func() {
		key, _ := masterKeyBinding.Get()

		globals.MasterHotkey.SetKey(dbmanager.SanitizeHotkey(key))
		dbmanager.StoreMasterHotkeyKey()
		hotkeys.NotifyHotkeysUpdated()

		if strings.ToUpper(key) != strings.ToUpper(masterHotkey.Selected) {
			logger.LogDebugVerboseOverride("masterKeyBinding changed: ["+key+"]; updating ui select element...", false)
			if len(key) == 1 {
				masterHotkey.SetSelected(strings.ToUpper(key))
			} else {
				masterHotkey.SetSelected(keyOptions[0])
			}
		} else {
			logger.LogDebugVerboseOverride("masterKeyBinding change listener: ui select element already up to date => ["+key+"]", false)
		}
	}))
	masterKeyBinding.Set(keyOptions[0])

	// maps switch
	mapsLabel := widget.NewLabel("Go to Maps")

	mapsShiftCb := widget.NewCheckWithData("Shift", mapsShiftModifierBinding)
	mapsShiftModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := mapsShiftModifierBinding.Get()
		globals.MapsHotkey.ShiftKey = value
		dbmanager.StoreMapsHotkeyShiftModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	mapsCtrlCb := widget.NewCheckWithData("Ctrl", mapsCtrlModifierBinding)
	mapsCtrlModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := mapsCtrlModifierBinding.Get()
		globals.MapsHotkey.CtrlKey = value
		dbmanager.StoreMapsHotkeyCtrlModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	mapsAltCb := widget.NewCheckWithData("Alt", mapsAltModifierBinding)
	mapsAltModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := mapsAltModifierBinding.Get()
		globals.MapsHotkey.AltKey = value
		dbmanager.StoreMapsHotkeyAltModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	mapsHotkey := widget.NewSelect(keyOptions, func(s string) {
		mapsKeyBinding.Set(strings.ToLower(s))
	})
	mapsKeyBinding.AddListener(binding.NewDataListener(func() {
		key, _ := mapsKeyBinding.Get()

		globals.MapsHotkey.SetKey(dbmanager.SanitizeHotkey(key))
		dbmanager.StoreMapsHotkeyKey()
		hotkeys.NotifyHotkeysUpdated()

		if strings.ToUpper(key) != strings.ToUpper(mapsHotkey.Selected) {
			logger.LogDebugVerboseOverride("mapsKeyBinding changed: ["+key+"]; updating ui select element...", false)
			if len(key) == 1 {
				mapsHotkey.SetSelected(strings.ToUpper(key))
			} else {
				mapsHotkey.SetSelected(keyOptions[0])
			}
		} else {
			logger.LogDebugVerboseOverride("mapsKeyBinding change listener: ui select element already up to date => ["+key+"]", false)
		}
	}))
	mapsKeyBinding.Set(keyOptions[0])

	// charts switch
	chartsLabel := widget.NewLabel("Go to Charts")

	chartsShiftCb := widget.NewCheckWithData("Shift", chartsShiftModifierBinding)
	chartsShiftModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := chartsShiftModifierBinding.Get()
		globals.ChartsHotkey.ShiftKey = value
		dbmanager.StoreChartsHotkeyShiftModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	chartsCtrlCb := widget.NewCheckWithData("Ctrl", chartsCtrlModifierBinding)
	chartsCtrlModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := chartsCtrlModifierBinding.Get()
		globals.ChartsHotkey.CtrlKey = value
		dbmanager.StoreChartsHotkeyCtrlModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	chartsAltCb := widget.NewCheckWithData("Alt", chartsAltModifierBinding)
	chartsAltModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := chartsAltModifierBinding.Get()
		globals.ChartsHotkey.AltKey = value
		dbmanager.StoreChartsHotkeyAltModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	chartsHotkey := widget.NewSelect(keyOptions, func(s string) {
		chartsKeyBinding.Set(strings.ToLower(s))
	})
	chartsKeyBinding.AddListener(binding.NewDataListener(func() {
		key, _ := chartsKeyBinding.Get()

		globals.ChartsHotkey.SetKey(dbmanager.SanitizeHotkey(key))
		dbmanager.StoreChartsHotkeyKey()
		hotkeys.NotifyHotkeysUpdated()

		if strings.ToUpper(key) != strings.ToUpper(chartsHotkey.Selected) {
			logger.LogDebugVerboseOverride("chartsKeyBinding changed: ["+key+"]; updating ui select element...", false)
			if len(key) == 1 {
				chartsHotkey.SetSelected(strings.ToUpper(key))
			} else {
				chartsHotkey.SetSelected(keyOptions[0])
			}
		} else {
			logger.LogDebugVerboseOverride("chartsKeyBinding change listener: ui select element already up to date => ["+key+"]", false)
		}
	}))
	chartsKeyBinding.Set(keyOptions[0])

	// notepad switch
	notepadLabel := widget.NewLabel("Go to Notepad")

	notepadShiftCb := widget.NewCheckWithData("Shift", notepadShiftModifierBinding)
	notepadShiftModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := notepadShiftModifierBinding.Get()
		globals.NotepadHotkey.ShiftKey = value
		dbmanager.StoreNotepadHotkeyShiftModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	notepadCtrlCb := widget.NewCheckWithData("Ctrl", notepadCtrlModifierBinding)
	notepadCtrlModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := notepadCtrlModifierBinding.Get()
		globals.NotepadHotkey.CtrlKey = value
		dbmanager.StoreNotepadHotkeyCtrlModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	notepadAltCb := widget.NewCheckWithData("Alt", notepadAltModifierBinding)
	notepadAltModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := notepadAltModifierBinding.Get()
		globals.NotepadHotkey.AltKey = value
		dbmanager.StoreNotepadHotkeyAltModifier()
		hotkeys.NotifyHotkeysUpdated()
	}))

	notepadHotkey := widget.NewSelect(keyOptions, func(s string) {
		notepadKeyBinding.Set(strings.ToLower(s))
	})
	notepadKeyBinding.AddListener(binding.NewDataListener(func() {
		key, _ := notepadKeyBinding.Get()

		globals.NotepadHotkey.SetKey(dbmanager.SanitizeHotkey(key))
		dbmanager.StoreNotepadHotkeyKey()
		hotkeys.NotifyHotkeysUpdated()

		if strings.ToUpper(key) != strings.ToUpper(notepadHotkey.Selected) {
			logger.LogDebugVerboseOverride("notepadKeyBinding changed: ["+key+"]; updating ui select element...", false)
			if len(key) == 1 {
				notepadHotkey.SetSelected(strings.ToUpper(key))
			} else {
				notepadHotkey.SetSelected(keyOptions[0])
			}
		} else {
			logger.LogDebugVerboseOverride("notepadKeyBinding change listener: ui select element already up to date => ["+key+"]", false)
		}
	}))
	notepadKeyBinding.Set(keyOptions[0])

	grid := container.NewGridWithColumns(
		3,
		masterLabel, container.NewHBox(masterShiftCb, masterCtrlCb, masterAltCb), masterHotkey,
		mapsLabel, container.NewHBox(mapsShiftCb, mapsCtrlCb, mapsAltCb), mapsHotkey,
		chartsLabel, container.NewHBox(chartsShiftCb, chartsCtrlCb, chartsAltCb), chartsHotkey,
		notepadLabel, container.NewHBox(notepadShiftCb, notepadCtrlCb, notepadAltCb), notepadHotkey,
		//msfsAutostartLabel, msfsAutostartCb, widget.NewLabel(""),
	)

	vBox := container.NewVBox(
		labelNotes,
		widget.NewLabel(""),
		grid,
	)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebugVerboseOverride("Hotkeys Panel initialized", false)

	return centerContainer

}
