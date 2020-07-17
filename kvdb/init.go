package kvdb

import (
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

func InitDB() (*bolt.DB, error) {

	ex, _ := os.Executable()

	exPath := filepath.Dir(ex)

	if _, err := os.Stat(exPath + "/db"); os.IsNotExist(err) {
		os.Mkdir(exPath+"/db", 0700)
	}

	db, _ := bolt.Open(exPath+"/db/elephant.db", 0644, nil)

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("stacks"))
		if err != nil {
			return err
		}
		return nil
	})

	return db, err

}
