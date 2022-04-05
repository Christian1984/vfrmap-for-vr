package dbmanager

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
)

type StorageData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Sender string `json:"sender,omitempty"`
}

type StorageDataSet struct {
	DataSets []StorageData `json:"data"`
	Sender string `json:"sender,omitempty"`
}

type StorageDataKeysArray struct {
	Keys []string
}

func DataController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		logger.LogError("Method "+r.Method+" not allowed!", false)
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var res StorageData

	switch r.Method {
	case http.MethodGet:
		key := r.URL.Query().Get("key")

		if len(strings.TrimSpace(key)) == 0 {
			logger.LogError("Property \"key\" must NOT be empty!", false)
			http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
			return
		}

		out := DbReadData(strings.TrimSpace(key))
		res = StorageData{key, strings.TrimSpace(out), ""}
		break

	case http.MethodPost:
		var storageData StorageData
		sdErr := json.NewDecoder(r.Body).Decode(&storageData)
		if sdErr != nil {
			logger.LogError("Error in dataController POST method: "+sdErr.Error(), true)
			http.Error(w, sdErr.Error(), http.StatusBadRequest)
			return
		}

		logger.LogDebug("Received StorageData: key=["+strings.TrimSpace(storageData.Key)+"], value=["+strings.TrimSpace(storageData.Value)+"]", false)

		if len(strings.TrimSpace(storageData.Key)) == 0 {
			logger.LogError("Property \"key\" must NOT be empty!", false)
			http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
			return
		}

		DbWriteData(strings.TrimSpace(storageData.Key), strings.TrimSpace(storageData.Value))
		res = storageData

		globals.Notepad.BroadcastIfNote(storageData.Sender, storageData.Key)

		break
	}

	responseJson, jsonErr := json.Marshal(res)

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

func DataSetController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		logger.LogError("Method "+r.Method+" not allowed!", false)
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var res StorageDataSet

	switch r.Method {
	case http.MethodGet:
		keysString := r.URL.Query().Get("keys")

		logger.LogDebug("Received Keys for data retrieval (raw): "+keysString, false)

		keys := StorageDataKeysArray{}
		jsonErr := json.Unmarshal([]byte(keysString), &keys.Keys)

		if jsonErr != nil {
			logger.LogError("Error in dataSetController GET method: "+jsonErr.Error(), true)
			http.Error(w, jsonErr.Error(), http.StatusBadRequest)
			return
		}

		logger.LogDebug("Extracted "+strconv.Itoa(len(keys.Keys))+" keys for data retrieval:", false)
		for _, key := range keys.Keys {
			logger.LogDebug("  key=["+strings.TrimSpace(key)+"]", false)
		}

		for _, key := range keys.Keys {
			if len(strings.TrimSpace(key)) == 0 {
				logger.LogError("Property \"key\" must NOT be empty!", false)
				http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
				return
			}
		}

		for _, key := range keys.Keys {
			value := DbReadData(strings.TrimSpace(key))
			sd := StorageData{strings.TrimSpace(key), value, ""}
			res.DataSets = append(res.DataSets, sd)
		}

		break

	case http.MethodPost:
		var storageDataSet StorageDataSet
		sdErr := json.NewDecoder(r.Body).Decode(&storageDataSet)
		if sdErr != nil {
			logger.LogError("Error in dataSetController POST method: "+sdErr.Error(), false)
			http.Error(w, sdErr.Error(), http.StatusBadRequest)
			return
		}

		logger.LogDebug("Received "+strconv.Itoa(len(storageDataSet.DataSets))+" StorageDataSet for storage:", false)
		for _, ds := range storageDataSet.DataSets {
			logger.LogDebug("StorageData: key=["+strings.TrimSpace(ds.Key)+"], value=["+strings.TrimSpace(ds.Value)+"]", false)
		}

		for _, ds := range storageDataSet.DataSets {
			if len(strings.TrimSpace(ds.Key)) == 0 {
				logger.LogError("Property \"key\" must NOT be empty!", false)
				http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
				return
			}
		}

		var keys []string

		for _, ds := range storageDataSet.DataSets {
			DbWriteData(strings.TrimSpace(ds.Key), strings.TrimSpace(ds.Value))
			keys = append(keys, ds.Key)
		}

		globals.Notepad.BroadcastIfContainsNote(storageDataSet.Sender, keys)

		res = storageDataSet
		break
	}

	responseJson, jsonErr := json.Marshal(res)

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