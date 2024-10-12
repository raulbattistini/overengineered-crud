package services

import (
	"fmt"
	"server/hepers"
	"server/repositories"
	"server/types"
	"server/valiation"
)

type PostService struct {
	PostRepository repositories.PostRepository
	PostValidtor   valiation.PostValidtor
	PostLogger     hepers.LoggerInterface
}

func NewPostService(pr repositories.PostRepository, pv valiation.PostValidtor, pl hepers.LoggerInterface) *PostService {
	return &PostService{
		PostRepository: pr,
		PostValidtor:   pv,
		PostLogger:     pl,
	}
}

// missing valiation of input and appropriate handling
func (ps *PostService) GetPostById(id string) (*types.Post, error) {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateGet(id); err != nil {
		msg := fmt.Sprintf("error valiating get of id %s", id)
		ps.PostLogger.LogError(&msg, err)
		return nil, err
	}
	return ps.PostRepository.FindById(id)
}

func (ps *PostService) CreateMewPost(pst *types.Post) (*types.Post, error) {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateMewPost(pst); err != nil {
		msg := fmt.Sprintf("error valiating creatin of post %v", pst)
		ps.PostLogger.LogError(&msg, err)
		return nil, err
	}
	return ps.PostRepository.CreatePost(pst)
}

func (ps *PostService) UpdatePostById(pst *types.Post) error {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateUpdate(pst); err != nil {
		msg := fmt.Sprintf("error valiating updating the post with body %v", pst)
		ps.PostLogger.LogError(&msg, err)
		return err
	}
	return ps.PostRepository.UpdatePost(pst)
}

func (ps *PostService) RemovePostById(id string) error {
	// missing valiation of input (PostValiator injected here will do) and appropriate handling
	if err := ps.PostValidtor.ValiateRemovePost(id); err != nil {
		msg := fmt.Sprintf("error valiating removal off post with id %s", id)
		ps.PostLogger.LogError(&msg, err)
		return err
	}
	return ps.PostRepository.RemovePost(id)
}
