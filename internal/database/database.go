package database

import (
	"github.com/hashicorp/go-memdb"
	"log"
	"sync"
)

var memDBInstance *memdb.MemDB
var once sync.Once

func CreateDB() {
	once.Do(func() {
		db, err := memdb.NewMemDB(DBSchema)
		if err != nil {
			log.Fatalf("failed to create database: %s\n", err)
		}
		memDBInstance = db
	})
}

func GetDB() *memdb.MemDB {
	return memDBInstance
}
