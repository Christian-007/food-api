package restaurant

import "context"

type RepositoryInmemory struct {
	restaurants map[string]*Restaurant
}

func NewRepositoryInMemory() *RepositoryInmemory {
	return &RepositoryInmemory{
		restaurants: make(map[string]*Restaurant),
	}
}

func (r *RepositoryInmemory) Add(ctx context.Context, entity *Restaurant) error {
	if r.isExist(entity.ID) {
		return ErrDuplicateIdentifier
	}

	r.restaurants[entity.ID] = entity
	return nil
}

func (r *RepositoryInmemory) Replace(ctx context.Context, entity *Restaurant) error {
  if !r.isExist(entity.ID) {
		return ErrNotFound
	}

	r.restaurants[entity.ID] = entity
	return nil
}

func (r *RepositoryInmemory) Remove(ctx context.Context, id string) error {
	if !r.isExist(id) {
		return ErrNotFound
	}

	delete(r.restaurants, id)
	return nil
}

func (r *RepositoryInmemory) Get(ctx context.Context, id string) (*Restaurant, error) {
	if !r.isExist(id) {
		return nil, ErrNotFound
	}

	return r.restaurants[id], nil
}

func (r *RepositoryInmemory) GetAll(ctx context.Context) ([]*Restaurant, error) {
	var restaurants []*Restaurant
	for _, restaurant := range r.restaurants {
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func (r *RepositoryInmemory) isExist(id string) bool {
	_,ok := r.restaurants[id]
	return ok
}
