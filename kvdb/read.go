package kvdb

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ReadData(db *bolt.DB, key string) string {

	var res []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("stacks"))

		res = bucket.Get([]byte(key))

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return string(res)
}
