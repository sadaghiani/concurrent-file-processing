package information

import (
	"context"
)

func (i *Information) Health(ctx context.Context) error {

	return i.Repository.IRepositoryMongo.Ping(ctx)
}
