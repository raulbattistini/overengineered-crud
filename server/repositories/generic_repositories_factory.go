package repositories

import "gorm.io/gorm"

type RepositoryFactory[T any] struct {
	Database *gorm.DB
}

func NewRepositoryFactory[T any](db *gorm.DB) *RepositoryFactory[T] {
	return &RepositoryFactory[T]{Database: db}
}

func (f *RepositoryFactory[T]) CreateRepository() *GenericGormRepository[T] {
	return NewGenericGormRepository[T](f.Database)
}
