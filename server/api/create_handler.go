package api

import (
	"encoding/json"
	"errors"
	"log"
	"server/types"
)

func CreatePost(body []byte) (types.DefaultResponseMessage[interface{}], error) {
	var err error
	var responseMsg types.RresponseMessage
	var status int
	var response types.DefaultResponseMessage[interface{}]
	var post types.Post

	if err = json.Unmarshal(body, &post); err != nil {
		responseMsg = types.BadRequestInvalidBody
		status = types.MapToStatusCode(responseMsg)
		message := "Invalid body"
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	if post.Content == "" {
		message := "No content provided"
		err = errors.New(string(types.BadRequestInvalidBody))
		responseMsg = types.BadRequestInvalidBody
		status = types.MapToStatusCode(responseMsg)
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	log.Printf("body received to create post\n%s", body)
	if post.Title == nil {
		post.Title = "No title"
	}

	if post.Title == "abc" { // should check against the database
		message := "Post already exists"
		err = errors.New(string(types.AlreadyExists))
		responseMsg = types.AlreadyExists
		status = types.MapToStatusCode(responseMsg)
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	responseMsg = types.Created
	status = types.MapToStatusCode(responseMsg)
	message := post
	response.Message = message
	response.Status = status
	response.Code = responseMsg

	return response, nil
}
