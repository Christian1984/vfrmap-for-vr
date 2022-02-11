package notepad

import (
	"vfrmap-for-vr/vfrmap/websockets"
)


type Notepad struct {
}

func New(ws *websockets.Websocket, verbose bool) Notepad {
	np := Notepad{}
	return np
}

func (np *Notepad) BroadcastNote() {
}
