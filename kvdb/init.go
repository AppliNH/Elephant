package kvdb

import "github.com/boltdb/bolt"

type Transac *bolt.Tx


func InitDB() (*bolt.Tx, error)  {

	db, _ := bolt.Open("db/elephant.db", 0600, nil)
	
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("stacks"))
		if err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		t, erro :=db.Begin(true)
		return t, erro
		
	}
	return nil, err

}