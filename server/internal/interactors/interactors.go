package interactors

import "server/internal/repository"

type ServerInteractor struct {
	Database *repository.InMemoryDatabase
}

func NewServerInteractor() (*ServerInteractor, error) {
	inMemoryDB, err := repository.NewInMemoryDatabase()
	if err != nil {
		return nil, err
	}

	return &ServerInteractor{
		Database: inMemoryDB,
	}, nil
}
