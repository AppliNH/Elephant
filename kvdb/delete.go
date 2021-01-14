package kvdb

import "github.com/boltdb/bolt"

// DeleteKey deletes a key in the kvdb
func DeleteKey(db *bolt.DB, key string) error {

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("stacks"))

		return bucket.Delete([]byte(key))
	})

}
