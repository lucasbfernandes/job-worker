package repository

import (
	"github.com/hashicorp/go-memdb"

	"fmt"
	"sync"
)

type InMemoryDatabase struct {
	instance *memdb.MemDB

	mutex sync.Mutex
}

func NewInMemoryDatabase() (*InMemoryDatabase, error) {
	dbInstance, err := memdb.NewMemDB(DBSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage: %s", err)
	}

	return &InMemoryDatabase{
		instance: dbInstance,
	}, nil
}
