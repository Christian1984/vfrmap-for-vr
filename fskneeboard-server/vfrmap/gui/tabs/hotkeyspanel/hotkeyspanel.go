package hotkeyspanel

import (
	"strings"
	"vfrmap-for-vr/vfrmap/application/dbmanager"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"

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

func HotkeysPanel() *fyne.Container {
	logger.LogDebug("Initializing Hotkeys Panel...", false)

	// grid and centerContainer
	labelNotes := widget.NewLabel("IMPORTANT NOTES:\n" + 
		"- In order for the hotkeys to work, you have to manually open the FSKneeboard panel at least once per flight!\n" +
		"- Please completely close and re-open your FSKneeboard ingame panel after changing any hotkeys!\n" +
		"- Use your HOTAS software to map these hotkeys to your HOTAS!\n")

	// master switch
	masterLabel := widget.NewLabel("Toggle FSKneeboard Panel")

	masterShiftCb := widget.NewCheckWithData("Shift", masterShiftModifierBinding)
	masterShiftModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterShiftModifierBinding.Get()
		globals.MasterHotkey.ShiftKey = value
		dbmanager.StoreMasterHotkeyShiftModifier()
	}))

	masterCtrlCb := widget.NewCheckWithData("Ctrl", masterCtrlModifierBinding)
	masterCtrlModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterCtrlModifierBinding.Get()
		globals.MasterHotkey.CtrlKey = value
		dbmanager.StoreMasterHotkeyCtrlModifier()
	}))

	masterAltCb := widget.NewCheckWithData("Alt", masterAltModifierBinding)
	masterAltModifierBinding.AddListener(binding.NewDataListener(func() {
		value, _ := masterAltModifierBinding.Get()
		globals.MasterHotkey.AltKey = value
		dbmanager.StoreMasterHotkeyAltModifier()
	}))

	masterHotkey := widget.NewSelect(keyOptions, func(s string) {
		masterKeyBinding.Set(strings.ToLower(s))
	})
	masterKeyBinding.AddListener(binding.NewDataListener(func() {
		key, _ := masterKeyBinding.Get()

		globals.MasterHotkey.SetKey(dbmanager.SanitizeHotkey(key))
		dbmanager.StoreMasterHotkeyKey()

		if strings.ToUpper(key) != strings.ToUpper(masterHotkey.Selected) {
			logger.LogDebug("masterKeyBinding changed: [" + key + "]; updating ui select element...", false)
			if len(key) == 1 {
				masterHotkey.SetSelected(strings.ToUpper(key))
			} else {
				masterHotkey.SetSelected(keyOptions[0])
			}
		} else {
			logger.LogDebug("masterKeyBinding change listener: ui select element already up to date => [" + key + "]", false)
		}
	}))
	masterKeyBinding.Set(keyOptions[0])

	grid := container.NewGridWithColumns(
		3,
		masterLabel, container.NewHBox(masterShiftCb, masterCtrlCb, masterAltCb), masterHotkey,
		//msfsAutostartLabel, msfsAutostartCb, widget.NewLabel(""),
	)
	
	labelSpoiler := widget.NewLabel("(more hotkeys coming soon...)")

	vBox := container.NewVBox(
		labelNotes,
		widget.NewLabel(""),
		grid,
		widget.NewLabel(""),
		labelSpoiler,
	)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebug("Hotkeys Panel initialized", false)

	return centerContainer

}