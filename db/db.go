package db

import (
	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func InitDB() error {
	var err error
	DB, err = bolt.Open("database.db", 0600, nil)
	if err != nil {
		return err
	}

	return DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("offers"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("oracles"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("messages"))
		if err != nil {
			return err
		}
		return nil
	})
}

func CloseDB() {
	DB.Close()
}