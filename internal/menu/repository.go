package menu

import "context"

type Repository interface {
	Add(ctx context.Context, r *Menu) error
	// Replace(ctx context.Context, r *Menu) error
	Remove(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Menu, error)
	GetAll(ctx context.Context) ([]*Menu, error)
	GetAllByRestaurant(ctx context.Context, id string) ([]*Menu, error)
}