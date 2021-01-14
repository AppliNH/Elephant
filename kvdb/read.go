package kvdb

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

func ReadData(db *bolt.DB, key string) (string, error) {

	var res []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("stacks"))

		res = bucket.Get([]byte(key))

		if len(res) == 0 {
			return errors.New("No data found for " + key)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return string(res), nil
}
func ReadAll(db *bolt.DB) (map[string]string, error) {

	data := make(map[string]string)

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("stacks"))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			val := fmt.Sprintf("%s", v)
			key := fmt.Sprintf("%s", k)
			data[key] = val
		}
		return nil
	})

	return data, err
}
