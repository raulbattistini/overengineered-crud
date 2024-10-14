package repositories

import "gorm.io/gorm"

type GenericRepository[T any] interface {
	FindById(id string) (*T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) error
	Remove(id string) error
	// Delete(entity *T) error // tbd if used
}

type GenericGormRepository[T any] struct {
	Database *gorm.DB
}

func NewGenericGormRepository[T any](db *gorm.DB) *GenericGormRepository[T] {
	return &GenericGormRepository[T]{Database: db}
}

func (r *GenericGormRepository[T]) FindById(id string) (*T, error) {
	var entity T
	if err := r.Database.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericGormRepository[T]) Create(entity *T) (*T, error) {
	if err := r.Database.Save(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *GenericGormRepository[T]) Update(entity *T) error {
	return r.Database.Save(entity).Error
}

func (r *GenericGormRepository[T]) Remove(id string) error {
	var entity T
	return r.Database.Where("id = ?", id).Delete(&entity).Error
}
