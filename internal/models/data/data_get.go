package data

import (
	"context"

	"github.com/sadaghiani/concurrent-file-processing/internal/entities"
)

func (d *Data) Get(ctx context.Context, page, limit int) ([]entities.Data, error) {

	data, err := d.Repository.IRepositoryMongo.Find(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}
