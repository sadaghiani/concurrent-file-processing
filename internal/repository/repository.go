package repository

import (
	"github.com/sadaghiani/concurrent-file-processing/pkg/database"
)

type Repository struct {
	IRepositoryMongo IRepositoryMongo
}

func NewRepository(mongoDataStore database.IMongoDataStore) *Repository {
	return &Repository{
		newRepositoryMongo(mongoDataStore),
	}
}
