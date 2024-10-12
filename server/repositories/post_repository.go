package repositories

import (
	"log"
	"server/types"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindById(id string) (*types.Post, error)
	CreatePost(post *types.Post) (*types.Post, error)
	UpdatePost(post *types.Post) error
	RemovePost(id string) error
}

type GormPostRepository struct {
	Database *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) *GormPostRepository {
	if db == nil {
		return nil
	}
	return &GormPostRepository{Database: db.Table("posts")}
}

func (r *GormPostRepository) FindById(id string) (*types.Post, error) {
	var post types.Post
	if err := r.Database.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepository) UpdatePost(post *types.Post) error {
	return r.Database.Save(&post).Error
}

func (r *GormPostRepository) CreatePost(post *types.Post) (*types.Post, error) {
	if err := r.Database.Save(&post).Error; err != nil {
		log.Println("error\n", err)
		return nil, err
	}
	return post, nil
}

func (r *GormPostRepository) RemovePost(id string) error {
	var post types.Post
	return r.Database.Where("id = ?", id).Delete(&post).Error
}
