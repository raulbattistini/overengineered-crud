package api

import (
	"server/db"
	"server/enums"
	"server/repositories"
	"server/services"
	"server/types"
	"server/valiation"
	"strconv"
)

func GetPostIdHandler(id string) (types.DefaultResponseMessage[types.Post], error) {
	postRepo, err := repositories.NewGormPostRepository(db.DB)
	if err != nil {
		code := enums.InternalServerError
		return types.DefaultResponseMessage[types.Post]{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	}

	intId, err := strconv.Atoi(id) // this is misplaced
	if err != nil {
		code := enums.InvalidIdFormat
		return types.DefaultResponseMessage[types.Post]{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	}

	postValidtor, err := valiation.NewPostValiddor(&types.Post{
		Id: intId,
	})
	postService := services.NewPostService(postRepo, postValidtor)
	post, err := postService.GetPostById(id)

	if err != nil {
		errCode := err.Error()

		switch errCode {
		case string(enums.ValidtorInvalidIdFormat):
			code := enums.InvalidIdFormat
			return types.DefaultResponseMessage[types.Post]{
				Code:    code,
				Status:  enums.MapToStatusCode(code),
				Message: types.Post{},
			}, err
		case string(enums.ValidtorNoIdRecived):
			code := enums.NoIdProvided
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
	code := enums.Success
	return types.DefaultResponseMessage[types.Post]{
		Code:    code,
		Status:  enums.MapToStatusCode(code),
		Message: *post,
	}, nil
}
