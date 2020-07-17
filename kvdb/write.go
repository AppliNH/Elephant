package kvdb

import (
	"github.com/boltdb/bolt"
)

func WriteData(db *bolt.DB, key string, value string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("stacks"))

		bucket.Put([]byte(key), []byte(value))

		return nil
	})

	return err

}
