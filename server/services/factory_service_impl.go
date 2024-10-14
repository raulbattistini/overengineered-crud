package services

import (
	"server/repositories"
	"server/types"
)

type GenericPostService struct {
	Repo repositories.GenericRepository[types.Post]
}

func NewGenericPostService(factory *repositories.RepositoryFactory[types.Post]) *GenericPostService {
	return &GenericPostService{
		Repo: factory.CreateRepository(),
	}
}

func (s *GenericPostService) GetPostById(id string) (*types.Post, error) {
	return s.Repo.FindById(id)
}

func (s *GenericPostService) CreatePost(post *types.Post) (*types.Post, error) {
	return s.Repo.Create(post)
}

func (s *GenericPostService) UpdatePost(post *types.Post) error {
	return s.Repo.Update(post)
}

func (s *GenericPostService) RemovePost(id string) error {
	return s.Repo.Remove(id)
}
