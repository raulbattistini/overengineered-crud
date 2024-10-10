package services

import (
	"server/repositories"
	"server/types"
	"server/valiation"
)

type PostService struct {
	PostRepository repositories.PostRepository
	PostValidtor   valiation.PostValidtor
}

func NewPostService(pr repositories.PostRepository, pv valiation.PostValidtor) *PostService {
	return &PostService{
		PostRepository: pr,
		PostValidtor:   pv,
	}
}

// missing valiation of input and appropriate handling
func (ps *PostService) GetPostById(id string) (*types.Post, error) {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateGet(id); err != nil {
		return nil, err
	}
	return ps.PostRepository.FindById(id)
}

func (ps *PostService) CreateMewPost(pst *types.Post) (*types.Post, error) {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateMewPost(pst); err != nil {
		return nil, err
	}
	return ps.PostRepository.CreatePost(pst)
}

func (ps *PostService) UpdatePostById(pst *types.Post) error {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateUpdate(pst); err != nil {
		return err
	}
	return ps.PostRepository.UpdatePost(pst)
}

func (ps *PostService) RemovePostById(id string) error {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateRemovePost(id); err != nil {
		return err
	}
	return ps.PostRepository.RemovePost(id)
}
