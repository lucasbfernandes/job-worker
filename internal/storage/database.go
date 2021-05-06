package storage

import (
	"github.com/hashicorp/go-memdb"

	"log"
	"sync"
)

var memDBInstance *memdb.MemDB
var once sync.Once
var createDBErr error

func CreateDB() error {
	once.Do(func() {
		db, err := memdb.NewMemDB(DBSchema)
		if err != nil {
			log.Printf("failed to create storage: %s\n", err)
			createDBErr = err
		}
		memDBInstance = db
	})
	return createDBErr
}

func GetDB() *memdb.MemDB {
	return memDBInstance
}
