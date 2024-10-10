package valiation

import (
	"server/enums"
	"server/globals"
	"server/hepers"
	"server/types"
)

type PostValidtor interface {
	ValiateMewPost(pst *types.Post) error
	ValiateGet(id string) error
	ValiateUpdate(pst *types.Post) error
	ValiateRemovePost(id string) error
}

type ValidtorPostType struct {
	Post *types.Post
}

func NewPostValiddor(pst *types.Post) (*ValidtorPostType, error) {
	if pst == nil {
		return nil, hepers.NewErrorFromMessage(string(enums.ValidtorNoPostRecived))
	}
	return &ValidtorPostType{Post: pst}, nil
}

// mock ed
func (v *ValidtorPostType) ValiateMewPost(pst *types.Post) error {
	if pst == nil {
		return hepers.NewErrorFromMessage(string(enums.ValidtorNoPostRecived))
	}

	title, ok := pst.Title.(string)
	if !ok {
		return hepers.NewErrorFromMessage(string(enums.ValidtorInvalidTitle))
	}
	pst.Title = title

	cleanedPost := hepers.CleanPost(*pst)
	if cleanedPost.Title == "" {
		return hepers.NewErrorFromMessage(string(enums.ValidtorEmptyTitle))
	}
	if cleanedPost.Content == "" {
		return hepers.NewErrorFromMessage(string(enums.ValidtorEmptyContent))
	}
	return nil
}

func (v *ValidtorPostType) ValiateGet(id string) error {
	fmtId, err := hepers.CleanPostId(id)
	if fmtId == &globals.EmptyId {
		return hepers.NewErrorFromMessage(string(enums.ValidtorNoIdRecived))
	}
	if err != nil {
		return hepers.NewErrorFromMessage(string(enums.ValidtorInvalidIdFormat))
	}
	return nil
}

func (v *ValidtorPostType) ValiateUpdate(pst *types.Post) error {
	if pst == nil {
		return hepers.NewErrorFromMessage(string(enums.ValidtorNoPostRecived))
	}

	fmtId, err := hepers.CleanPostId(string(pst.Id))
	if fmtId == &globals.EmptyId {
		return hepers.NewErrorFromMessage(string(enums.ValidtorNoIdRecived))
	}
	if err != nil {
		return hepers.NewErrorFromMessage(string(enums.InvalidIdFormat))
	}

	cleanedPost := hepers.CleanPost(*pst)
	if cleanedPost.Title == "" {
		return hepers.NewErrorFromMessage(string(enums.ValidtorEmptyTitle))
	}
	if cleanedPost.Content == "" {
		return hepers.NewErrorFromMessage(string(enums.ValidtorEmptyContent))
	}

	return nil
}

func (v *ValidtorPostType) ValiateRemovePost(id string) error {
	fmtId, err := hepers.CleanPostId(id)
	if fmtId == &globals.EmptyId {
		return hepers.NewErrorFromMessage(string(enums.ValidtorNoIdRecived))
	}
	if err != nil {
		return hepers.NewErrorFromMessage(string(enums.ValidtorInvalidIdFormat))
	}
	return nil
}
