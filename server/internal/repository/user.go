package repository

import (
	"errors"
	"fmt"
	"server/internal/models/user"
)

func (db *InMemoryDatabase) UpsertUser(user *user.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(true)
	err := txn.Insert("user", user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %s", err)
	}
	txn.Commit()
	return nil
}

func (db *InMemoryDatabase) GetUserOrFailByAPIToken(apiToken string) (*user.User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "api_token", apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %s", err)
	}
	if raw == nil {
		return nil, errors.New("user doesn't exist")
	}

	return raw.(*user.User), nil
}
