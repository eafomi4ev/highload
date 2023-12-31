package rest

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

//func (e ErrorResponse) StatusCode() int {
//	return e.Status
//}

func BadRequest(msg string, err error) ErrorResponse {
	return buildErrorResponse(http.StatusBadRequest, msg, err)
}

func NotFound(msg string, err error) ErrorResponse {
	return buildErrorResponse(http.StatusNotFound, msg, err)
}

func InternalServerError(msg string, err error) ErrorResponse {
	return buildErrorResponse(http.StatusInternalServerError, msg, err)
}

func buildErrorResponse(status int, msg string, err error) ErrorResponse {
	if msg == "" {
		msg = http.StatusText(status)
	}

	er := ErrorResponse{
		Status:  status,
		Message: fmt.Errorf("%s: %w", msg, err).Error(),
	}

	return er
}
