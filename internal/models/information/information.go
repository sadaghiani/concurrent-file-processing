package information

import (
	"context"

	"github.com/sadaghiani/concurrent-file-processing/internal/repository"
)

type IInformation interface {
	Health(ctx context.Context) error
}

type Information struct {
	*repository.Repository
}

func NewInformation(repo *repository.Repository) *Information {
	return &Information{
		Repository: repo,
	}
}
