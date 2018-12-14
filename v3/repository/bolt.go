package repository

import (
	"log"

	"github.com/boltdb/bolt"
)

// BoltClient will be used to connect to BoltDB
type BoltClient struct {
	boltdb *bolt.DB
}

// Open - Opens up the database conneciton
func (bc *BoltClient) Open() {
	var err error
	bc.boltdb, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}
