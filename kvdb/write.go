package kvdb

import "github.com/boltdb/bolt"

func WriteData(tx *bolt.Tx,key string, value string) error {

	b := tx.Bucket([]byte("stacks"))
	return b.Put([]byte(key), []byte(value))

}