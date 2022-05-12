package hotkeys

import (
	"encoding/json"
	"net/http"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
	"vfrmap-for-vr/vfrmap/websockets"
)

var Ws *websockets.Websocket

func NotifyHotkeysUpdated() {
	if Ws != nil {
		logger.LogDebugVerboseOverride("Broadcasting hotkeys update!", false)

		Ws.Broadcast(map[string]interface{}{
			"msg": "hotkeys updated",
		})
	}
}

func ServeMasterHotkey(w http.ResponseWriter, r *http.Request) {
	logger.LogDebugVerboseOverride("ServeMasterHotkey called!", false)

	responseJson, jsonErr := json.Marshal(globals.MasterHotkey)

	if jsonErr != nil {
		utils.Println(jsonErr.Error())
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responseJson))
}
