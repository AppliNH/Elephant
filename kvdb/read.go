package kvdb

import "github.com/boltdb/bolt"


func ReadData(tx *bolt.Tx,key string) string {
	b := tx.Bucket([]byte("stacks"))
	v := b.Get([]byte(key))
	return string(v)
}