package restaurant

import "context"

type Repository interface {
	Add(ctx context.Context, r *Restaurant) error
	Replace(ctx context.Context, r *Restaurant) error
	Remove(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Restaurant, error)
	GetAll(ctx context.Context) ([]*Restaurant, error)
}