/*
  microservices cqrs pattern tin rabzelj
*/
package db

import (
	"context"

	"github.com/renegmed/microserv-cqrs-natss/query-service/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

// For one, PostgresRepository implements this interface
var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
