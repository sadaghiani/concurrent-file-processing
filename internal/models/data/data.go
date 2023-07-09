package data

import (
	"context"

	"github.com/sadaghiani/concurrent-file-processing/internal/entities"
	"github.com/sadaghiani/concurrent-file-processing/internal/repository"
)

type IDeta interface {
	Get(ctx context.Context, page, limit int) ([]entities.Data, error)
}

type Data struct {
	*repository.Repository
}

func NewData(repo *repository.Repository) IDeta {
	return &Data{
		repo,
	}
}
