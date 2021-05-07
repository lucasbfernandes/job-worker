package repository

import (
	"github.com/hashicorp/go-memdb"

	"fmt"
	userEntity "server/internal/models/user"
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

// TODO this is mocked because login and users CRUD are not implemented. This will be removed in the next releases
func (db *InMemoryDatabase) SeedUsers() error {
	admin, err := userEntity.NewUser("admin", "qTMaYIfw8q3esZ6Dv2rQ", "ADMIN")
	if err != nil {
		return err
	}

	user, err := userEntity.NewUser("user", "9EzGJOTcMHFMXphfvAuM", "USER")
	if err != nil {
		return err
	}

	err = db.UpsertUser(admin)
	if err != nil {
		return err
	}

	err = db.UpsertUser(user)
	if err != nil {
		return err
	}

	return nil
}
