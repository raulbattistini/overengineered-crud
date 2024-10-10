package api

import (
	"encoding/json"
	"server/db"
	"server/enums"
	"server/repositories"
	"server/services"
	"server/types"
	"server/valiation"
)

func CreateAPost(body []byte) (types.DefaultResponseMessage[types.Post], error) {
	var post types.Post
	if err := json.Unmarshal(body, &post); err != nil {
		return types.DefaultResponseMessage[types.Post]{
			Code:    enums.BadRequestInvalidBody,
			Status:  enums.MapToStatusCode(enums.BadRequestInvalidBody),
			Message: types.Post{},
		}, err
	}

	postRepo, err := repositories.NewGormPostRepository(db.DB)
	if err != nil {
		code := enums.InternalServerError
		return types.DefaultResponseMessage[types.Post]{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	}

	postValidtor, err := valiation.NewPostValiddor(&post)
	if err != nil {
		code := enums.InternalServerError
		return types.DefaultResponseMessage[types.Post]{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	}

	postService := services.NewPostService(postRepo, postValidtor)
	maybePost, err := postService.CreateMewPost(&post)

	if err != nil {
		errCode := err.Error()
		switch errCode {
		case string(enums.ValidtorInvalidIdFormat):
			code := enums.BadRequestInvalidBody
			return types.DefaultResponseMessage[types.Post]{
				Code:    code,
				Status:  enums.MapToStatusCode(code),
				Message: types.Post{},
			}, err
		case string(enums.ValidtorNoPostRecived):
		case string(enums.ValidtorEmptyTitle):
		case string(enums.ValidtorEmptyContent):
			code := enums.BadRequestInvalidBody
			return types.DefaultResponseMessage[types.Post]{
				Code:    code,
				Status:  enums.MapToStatusCode(code),
				Message: types.Post{},
			}, err
		default:
			code := enums.InternalServerError
			return types.DefaultResponseMessage[types.Post]{
				Code:    code,
				Status:  enums.MapToStatusCode(code),
				Message: types.Post{},
			}, err
		}
	}
	code := enums.Created
	return types.DefaultResponseMessage[types.Post]{
		Code:    code,
		Status:  enums.MapToStatusCode(code),
		Message: *maybePost,
	}, nil
}
