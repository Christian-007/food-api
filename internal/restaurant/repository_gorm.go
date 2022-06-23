package restaurant

import (
	"context"
	"food/internal/common/dbutils"

	"gorm.io/gorm"
)

type RepositoryGorm struct {
	db *gorm.DB
}

func NewRepositoryGorm(db *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{
		db: db,
	}
}

func (r *RepositoryGorm) Add(ctx context.Context, entity *Restaurant) error {
  err := r.db.Create(entity).Error

	if err != nil {
		if dbutils.IsUniqueViolation(err) {
			return ErrDuplicateIdentifier
		}
	}

	return err
}

func (r *RepositoryGorm) Get(ctx context.Context, id string) (*Restaurant, error) {
	var restaurant Restaurant
	err := r.db.Where("id = ?", id).First(&restaurant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &restaurant, nil
}

func (r *RepositoryGorm) Replace(ctx context.Context, restaurant *Restaurant) error {
	err := r.db.Model(&Restaurant{}).Where("id = ?", restaurant.ID).Updates(restaurant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
	}

	return err
}

func (r *RepositoryGorm) Remove(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).First(&Restaurant{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	err = r.db.Delete(&Restaurant{}, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
	}

	return err
}

func (r *RepositoryGorm) GetAll(ctx context.Context) ([]*Restaurant, error) {
	var restaurants []*Restaurant
	err := r.db.Find(&restaurants).Error
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}
