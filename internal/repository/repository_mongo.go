package repository

import (
	"context"

	"github.com/sadaghiani/concurrent-file-processing/internal/entities"
	"github.com/sadaghiani/concurrent-file-processing/internal/utils"
	"github.com/sadaghiani/concurrent-file-processing/pkg/config"
	"github.com/sadaghiani/concurrent-file-processing/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepositoryMongo interface {
	Ping(ctx context.Context) error
	InsertBatch(ctx context.Context, input []interface{}) error
	Find(ctx context.Context, page, limit int) ([]entities.Data, error)
}

type repositoryMongo struct {
	database.IMongoDataStore
}

func newRepositoryMongo(m database.IMongoDataStore) IRepositoryMongo {
	return &repositoryMongo{
		m,
	}
}

func (r *repositoryMongo) Ping(ctx context.Context) error {
	client, err := r.GetClient()
	if err != nil {
		return err
	}
	return client.Ping(ctx, nil)
}

func (r *repositoryMongo) InsertBatch(ctx context.Context, input []interface{}) error {
	_, err := r.Coll(config.Config.GetString(utils.MustBindEnvToString(utils.DATABASE_COLLECTION))).InsertMany(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryMongo) Find(ctx context.Context, page, limit int) ([]entities.Data, error) {
	filter := bson.D{{}}
	options := new(options.FindOptions)
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		options.SetSkip(int64((page - 1) * limit))
		options.SetLimit(int64(limit))
	}
	cursor, err := r.Coll(config.Config.GetString(utils.MustBindEnvToString(utils.DATABASE_COLLECTION))).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var data []entities.Data
	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}
