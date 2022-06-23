package menu

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

func (r *RepositoryGorm) Add(ctx context.Context, entity *Menu) error {
  err := r.db.Create(entity).Error

	if err != nil {
		if dbutils.IsUniqueViolation(err) {
			return ErrDuplicateIdentifier
		}
	}

	return err
}

func (r *RepositoryGorm) Get(ctx context.Context, id string) (*Menu, error) {
	var menu Menu

	err := r.db.Where("id = ?", id).First(&menu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &menu, nil
}

func (r *RepositoryGorm) GetAll(ctx context.Context) ([]*Menu, error) {
	var menus []*Menu
	err := r.db.Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *RepositoryGorm) GetAllByRestaurant(ctx context.Context, id string) ([]*Menu, error) {
	var menus []*Menu
	err := r.db.Where("restaurant_id = ?", id).Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *RepositoryGorm) Remove(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).First(&Menu{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	err = r.db.Delete(&Menu{}, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
	}

	return err
}
