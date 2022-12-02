package dbmanager

import (
	"fmt"
	"time"
	"vfrmap-for-vr/vfrmap/logger"

	"github.com/boltdb/bolt"
)

const clientDataBucket = "fskneeboard"
const serverSettingsBucket = "serversettings"

const boltFileName = "fskneeboard.db"

var db *bolt.DB

func DbConnect() error {
	boltDb, db_err := bolt.Open(boltFileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if db_err != nil {
		return db_err
	}

	db = boltDb
	return nil
}

func initBucket(name string, tx *bolt.Tx) error {
	_, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		logger.LogError("Cannot create bucket " + name + " in db " + boltFileName + ", details: " + err.Error())
		return fmt.Errorf("Cannot create bucket: %s", err)
	}
	return nil
}

func DbInit() {
	db.Update(func(tx *bolt.Tx) error {
		dtbcktErr := initBucket(clientDataBucket, tx)
		if dtbcktErr != nil {
			return dtbcktErr
		}

		srvSttErr := initBucket(serverSettingsBucket, tx)
		if srvSttErr != nil {
			return srvSttErr
		}

		return nil
	})
}

func DbClose() {
	db.Close()
}

func dbWrite(bucket string, key string, value string) {
	logger.LogDebug("Storing data: [" + key + "]=[" + value + "] in bucket [" + bucket + "]")

	if db == nil {
		logger.LogDebug("DB not initialized! Cannot store value for [" + key + "]")
	} else {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			err := b.Put([]byte(key), []byte(value))
			return err
		})
	}
}

func dbRead(bucket string, key string) string {
	logger.LogDebug("Reading data for key [" + key + "] from bucket [" + bucket + "]")

	var out *string

	if db == nil {
		logger.LogDebug("DB not initialized! Cannot read value for [" + key + "]")
	} else {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			v := b.Get([]byte(key))

			outs := string(v[:])
			out = &outs

			logger.LogDebug("[" + key + "] is: [" + *out + "]")

			return nil
		})
	}

	return *out
}

func DbWriteData(key string, value string) {
	dbWrite(clientDataBucket, key, value)
}

func DbReadData(key string) string {
	return dbRead(clientDataBucket, key)
}

func DbWriteSettings(key string, value string) {
	dbWrite(serverSettingsBucket, key, value)
}

func DbReadSettings(key string) string {
	return dbRead(serverSettingsBucket, key)
}
