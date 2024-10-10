package types

import "server/enums"

type DefaultError struct {
	Code    enums.RresponseMessage `json:"code"`
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
}

func (e *DefaultError) ErrorMessage() string {
	return e.Message
}

func (e *DefaultError) ErrorStatus() int {
	return e.Status
}

func (e *DefaultError) ErrorCode() enums.RresponseMessage {
	return e.Code
}

func (e *DefaultError) ErrorCodeStr() string {
	return string(e.Code)
}
