package kvdb

import (
	"os"

	"github.com/boltdb/bolt"
)

func InitDB() (*bolt.DB, error) {
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		os.Mkdir("db", 0700)
	}
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
