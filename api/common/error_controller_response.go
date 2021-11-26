package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest errorControllerResponseCode = "bad_request"
	ErrForbidden  errorControllerResponseCode = "forbidden"
)

//ControllerResponse default payload response
type ControllerResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//NewBadRequestResponse bad request format response
func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		"Bad request",
		map[string]interface{}{},
	}
}

//NewBadRequestResponse bad request format response with message
func NewBadRequestResponseWithMessage(msg string) (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		msg,
		map[string]interface{}{},
	}
}

//NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{
		"Forbidden",
		map[string]interface{}{},
	}
}
