package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type StorageData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Sender string `json:"sender,omitempty`
}

type StorageDataSet struct {
	DataSets []StorageData `json:"data"`
	Sender string `json:"sender,omitempty`
}

type StorageDataKeysArray struct {
	Keys []string
}

const boltBucketName = "fskneeboard"
const boltFileName = "fskneeboard.db"

var db *bolt.DB

func dbConnect() error {
	boldDb, db_err := bolt.Open(boltFileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if db_err != nil {
		return db_err
	}

	db = boldDb
	return nil
}

func dbInit() {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltBucketName))
		if err != nil {
			return fmt.Errorf("Cannot create bucket: %s", err)
		}
		return nil
	})
}

func dbWrite(key string, value string) {
	if verbose {
		fmt.Printf("Storing data: %s = %s\n", key, value)
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltBucketName))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func dbRead(key string) string {
	/*
	if verbose {
		fmt.Printf("Reading data for key=%s\n", key)
	}
	*/

	var out *string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltBucketName))
		v := b.Get([]byte(key))

		outs := string(v[:])
		out = &outs

		if verbose {
			fmt.Printf("%s is: %s\n", key, *out)
		}

		return nil
	})

	return *out
}

// controller methods
func dataController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var res StorageData

	switch r.Method {
	case http.MethodGet:
		key := r.URL.Query().Get("key")

		if len(strings.TrimSpace(key)) == 0 {
			http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
			return
		}

		out := dbRead(strings.TrimSpace(key))
		res = StorageData{key, strings.TrimSpace(out), ""}
		break

	case http.MethodPost:
		var storageData StorageData
		sdErr := json.NewDecoder(r.Body).Decode(&storageData)
		if sdErr != nil {
			fmt.Println("Error in handleData: " + sdErr.Error())
			http.Error(w, sdErr.Error(), http.StatusBadRequest)
			return
		}

		/*
		if verbose {
			fmt.Println("Received StorageData: key=" + strings.TrimSpace(storageData.Key) + ", value=" + strings.TrimSpace(storageData.Value))
		}
		*/

		if len(strings.TrimSpace(storageData.Key)) == 0 {
			http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
			return
		}

		dbWrite(strings.TrimSpace(storageData.Key), strings.TrimSpace(storageData.Value))
		res = storageData
		
		np.BroadcastIfNote(storageData.Sender, storageData.Key)

		break
	}

	responseJson, jsonErr := json.Marshal(res)

	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
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

func dataSetController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var res StorageDataSet

	switch r.Method {
	case http.MethodGet:
		keysString := r.URL.Query().Get("keys")

		if verbose {
			fmt.Printf("Received Keys for data retrieval (raw): %s\n", keysString)
		}

		keys := StorageDataKeysArray{}
		jsonErr := json.Unmarshal([]byte(keysString), &keys.Keys)

		if jsonErr != nil {
			fmt.Println("Error in handleDataSet: " + jsonErr.Error())
			http.Error(w, jsonErr.Error(), http.StatusBadRequest)
			return
		}

		if verbose {
			fmt.Printf("Extracted %d Keys for data retrieval:\n", len(keys.Keys))
			for _, key := range keys.Keys {
				fmt.Println("key=" + strings.TrimSpace(key))
			}
		}

		for _, key := range keys.Keys {
			if len(strings.TrimSpace(key)) == 0 {
				http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
				return
			}
		}

		for _, key := range keys.Keys {
			value := dbRead(strings.TrimSpace(key))
			sd := StorageData{strings.TrimSpace(key), value, ""}
			res.DataSets = append(res.DataSets, sd)
		}

		break

	case http.MethodPost:
		var storageDataSet StorageDataSet
		sdErr := json.NewDecoder(r.Body).Decode(&storageDataSet)
		if sdErr != nil {
			fmt.Println("Error in handleData: " + sdErr.Error())
			http.Error(w, sdErr.Error(), http.StatusBadRequest)
			return
		}

		if verbose {
			fmt.Printf("Received %d StorageDataSet for storage:\n", len(storageDataSet.DataSets))
			/*
			for _, ds := range storageDataSet.DataSets {
				fmt.Println("StorageData: key=" + strings.TrimSpace(ds.Key) + ", value=" + strings.TrimSpace(ds.Value))
			}
			*/
		}

		for _, ds := range storageDataSet.DataSets {
			if len(strings.TrimSpace(ds.Key)) == 0 {
				http.Error(w, "Property \"key\" must NOT be empty!", http.StatusBadRequest)
				return
			}
		}

		var keys []string

		for _, ds := range storageDataSet.DataSets {
			dbWrite(strings.TrimSpace(ds.Key), strings.TrimSpace(ds.Value))
			keys = append(keys, ds.Key)
		}

		np.BroadcastIfContainsNote(storageDataSet.Sender, keys)

		res = storageDataSet
		break
	}

	responseJson, jsonErr := json.Marshal(res)

	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
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
