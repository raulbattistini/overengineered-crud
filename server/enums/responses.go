package enums

import "net/http"

type RresponseMessage string

const (
	NotFound                RresponseMessage = "NOT_FOUND"
	NoIdProvided            RresponseMessage = "NO_ID_FOUNDED"
	InvalidIdFormat         RresponseMessage = "INVALID_ID_FORM"
	InternalServerError     RresponseMessage = "INTERNAL_ERROR"
	BadRequestInvalidBody   RresponseMessage = "BAD_REQUEST_INVALID_BODY"
	Success                 RresponseMessage = "OK_FOUNDED"
	Created                 RresponseMessage = "OK_CREATED"
	AlreadyExists           RresponseMessage = "ALREADY_EXISTS"
	NoContent               RresponseMessage = "NO_CONTENT"
	NonAuthoritativeUpdated RresponseMessage = "NON_AU_THORITATIVE_UDATE"
)

func MapToStatusCode(msg RresponseMessage) int {
	switch msg {
	case Created:
		return http.StatusCreated // 201
	case Success:
		return http.StatusOK // 200
	case NonAuthoritativeUpdated:
		return http.StatusNonAuthoritativeInfo // 203
	case NoContent:
		return http.StatusNoContent // 204
	case NotFound:
		return http.StatusNotFound // 404
	case NoIdProvided, InvalidIdFormat, BadRequestInvalidBody:
		return http.StatusBadRequest // 400
	case AlreadyExists:
		return http.StatusConflict // 409
	case InternalServerError:
		return http.StatusInternalServerError // 500
	default:
		return http.StatusInternalServerError // Default to 500 for undefined messages
	}
}
