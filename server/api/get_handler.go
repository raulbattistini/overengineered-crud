package api

import (
	"errors"
	"server/context"
	"server/enums"
	"server/hepers"
	"server/repositories"
	"server/services"
	"server/types"
	"strconv"
	"strings"
)

func GetPostByIdHandler(id string) (types.DefaultResponseMessage[interface{}], error) {
	id = strings.Trim(id, "")
	var err error
	var responseMsg types.RresponseMessage
	var status int
	var response types.DefaultResponseMessage[interface{}]

	if id == "" {
		responseMsg = types.NoIdProvided
		err = errors.New(string(types.NoIdProvided))
		status = enums.MapToStatusCode(responseMsg)
		message := "No id was provided "
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	fmtId, err := strconv.Atoi(id)
	if err != nil {
		responseMsg = types.InvalidIdFormat
		status = enums.MapToStatusCode(responseMsg)
		message := "Invalid id format "
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	// mock
	if id == "7" {
		responseMsg = types.NotFound
		status = enums.MapToStatusCode(responseMsg)
		message := "Post pot found" // differentiated|special message for this case
		errors.New(message)
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	post := &types.Post{
		Id:      fmtId,
		Title:   hepers.GenRandomPostTitle(10),
		Content: hepers.GenRandomPostContent(150),
	}

	if err != nil {
		responseMsg = types.InternalServerError
		status = enums.MapToStatusCode(responseMsg)
		message := types.InternalServerError
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}
	responseMsg = types.Success
	status = enums.MapToStatusCode(responseMsg)

	postResponse := &post
	response.Message = postResponse
	response.Status = status
	response.Code = responseMsg
	return response, nil
}

func GetPostIdHandler(id string) (types.DefaultResponseMessage[types.Post], error) {
	postRepo := repositories.NewGormPostRepository(context.AppContext.Database)
	postValidtor := valiators.NewValidtorPostType(&types.Post{
		Id: id,
	})
	postService := services.NewPostService(postRepo, postValidtor)
	post, err := postService.GetPostById(id)

	errCode := err.Error()
	switch errCode {
	case nil:
		code := enums.Success
		return types.DefaultResponseMessage{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: post,
		}, nil
	case enums.ValidtorInvalidIdFormat:
		code := enums.InvalidIdFormat
		return types.DefaultResponseMessage{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	case enums.NoIdRecived:
		code := enums.NoIdProvided
		return types.DefaultResponseMessage{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	default:
		code := enums.InternalServerError
		return types.DefaultResponseMessage{
			Code:    code,
			Status:  enums.MapToStatusCode(code),
			Message: types.Post{},
		}, err
	}
}
