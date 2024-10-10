package types

import "server/enums"

type ResponseMessage interface {
	interface{} | map[string]interface{} | Post
}

type DefaultResponseMessage[MessageType ResponseMessage] struct {
	Code    enums.RresponseMessage `json:"code"`
	Status  int                    `json:"status"`
	Message MessageType            `json:"message"`
}

func (r *DefaultResponseMessage[MessageType]) ResponseMessage() MessageType {
	return r.Message
}

func (r *DefaultResponseMessage[MessageType]) ResponseStatus() int {
	return r.Status
}

func (r *DefaultResponseMessage[MessageType]) ResponseCode() enums.RresponseMessage {
	return r.Code
}

func (r *DefaultResponseMessage[MessageType]) ErrorCodeStr() string {
	return string(r.Code)
}

type PostResponse struct {
	Code    enums.RresponseMessage `json:"code"`
	Status  int                    `json:"status"`
	Message Post                   `json:"message"`
}
