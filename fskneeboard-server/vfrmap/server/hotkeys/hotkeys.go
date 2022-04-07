package hotkeys

import (
	"encoding/json"
	"net/http"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
)

func ServeMasterHotkey(w http.ResponseWriter, r *http.Request) {
	logger.LogDebug("ServeMasterHotkey called!", false)

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