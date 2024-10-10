package api

import (
	"encoding/json"
	"server/enums"
	"server/types"
	"strconv"
)

func AtualizePostById(body []byte, id string) (types.DefaultResponseMessage[map[string]interface{}], error) {
	var err error
	var responseMsg enums.RresponseMessage
	var status int
	var response types.DefaultResponseMessage[map[string]interface{}]
	var post types.Post

	if err = json.Unmarshal(body, &post); err != nil {
		responseMsg = enums.BadRequestInvalidBody
		status = enums.MapToStatusCode(responseMsg)
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
		responseMsg = enums.InvalidIdFormat
		status = enums.MapToStatusCode(responseMsg)
		message := map[string]interface{}{
			"content": "Invalid id format ",
		}
		response.Message = message
		response.Status = status
		response.Code = responseMsg
		return response, err
	}
	_ = fmtId

	resp, err := GetPostIdHandler(id)
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
	case enums.NotFound:
	default:
	}

	message := map[string]interface{}{
		"previous_post_info": "there shuld be a get here before",
		"message":            "Post updated",
	}
	status = enums.MapToStatusCode(responseMsg)
	responseMsg = enums.NonAuthoritativeUpdated

	response.Message = message
	response.Status = status
	response.Code = responseMsg

	return response, nil
}
