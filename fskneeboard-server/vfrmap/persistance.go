package main

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const boltBucketName = "fskneeboard"
const boltFileName = "fskneeboard.db"

var db *bolt.DB

func dbConnect() error {
	boldDb, db_err := bolt.Open(boltFileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if db_err != nil {
		return db_err
	}

	db = boldDb
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

	if verbose {
		fmt.Printf("Reading data for key=%s\n", key)
	}

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
