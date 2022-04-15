package hotkeys

import (
	"strings"
)

type Hotkey struct {
	AltKey   bool   `json:"altkey"`
	CtrlKey  bool   `json:"ctrlkey"`
	ShiftKey bool   `json:"shiftkey"`
	KeyCode  int    `json:"keycode"`
	Key      string `json:"key"`
}

func sanitizedKeycode(key string) int {
	keyCode := keycode(key)

	if keyCode == 0 {
		keyCode = -1
	}

	return keyCode
}

func (hc *Hotkey) SetKey(key string) {
	upperKey := strings.ToUpper(key)
	hc.Key = upperKey
	hc.KeyCode = sanitizedKeycode(upperKey)
}

func New(shiftModifier bool, ctrlModifier bool, altModifier bool, key string) Hotkey {
	upperKey := strings.ToUpper(key)
	keyCode := sanitizedKeycode(upperKey)
	return Hotkey{ShiftKey: shiftModifier, CtrlKey: ctrlModifier, AltKey: altModifier, KeyCode: keyCode, Key: upperKey}
}