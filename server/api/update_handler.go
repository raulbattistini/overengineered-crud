package api

import (
	"encoding/json"
	"server/types"
	"strconv"
)

func AtualizePostById(body []byte, id string) (types.DefaultResponseMessage[map[string]interface{}], error) {
	var err error
	var responseMsg types.RresponseMessage
	var status int
	var response types.DefaultResponseMessage[map[string]interface{}]
	var post types.Post

	if err = json.Unmarshal(body, &post); err != nil {
		responseMsg = types.BadRequestInvalidBody
		status = types.MapToStatusCode(responseMsg)
		message := map[string]interface{}{
			"body_sent": string(body), // ideally should be cleaning and sanitizing input before using it
			"contente":  "Invalid body provided",
		}
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}

	fmtId, err := strconv.Atoi(id)
	if err != nil {
		responseMsg = types.InvalidIdFormat
		status = types.MapToStatusCode(responseMsg)
		message := map[string]interface{}{
			"content": "Invalid id format ",
		}
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}
	_ = fmtId

	resp, err := GetPostByIdHandler(id)
	respCode := resp.ResponseCode()
	respStatus := resp.ResponseStatus()
	respMessage := resp.ResponseMessage()

	if err != nil {
		response.Message = map[string]interface{}{
			"response": respMessage,
		}
		response.Status = respStatus
		response.Code = respCode
		return response, err
	}
	switch respCode {
	case types.NotFound:
	default:
	}

	message := map[string]interface{}{
		"previous_post_info": "there shuld be a get here before",
		"message":            "Post updated",
	}
	status = types.MapToStatusCode(responseMsg)
	responseMsg = types.NonAuthoritativeUpdated

	response.Message = message
	response.Status = status
	response.Code = responseMsg

	return response, nil
}
