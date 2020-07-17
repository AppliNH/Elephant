package kvdb

import (
	"github.com/boltdb/bolt"
)

func InitDB() (*bolt.DB, error) {

	db, _ := bolt.Open("db/elephant.db", 0644, nil)

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("stacks"))
		if err != nil {
			return err
		}
		return nil
	})

	return db, err

}
