package api

import (
	"errors"
	"net/http"
)

// ErrorResponse ...
type ErrorResponse struct {
	Code           string `json:"code"`
	Description    string `json:"description"`
	HTTPStatusCode int    `json:"httpStatusCode"`
}

var (
	errInvalidRegistry = errors.New("Invalid Registry")
)

var errorResponseMap = map[error]ErrorResponse{
	errInvalidRegistry: {
		Code:           "errInvalidRegistry",
		Description:    "Invalid Registry",
		HTTPStatusCode: http.StatusBadRequest,
	},
}

// GetErrorResponse ...
func GetErrorResponse(e error) ErrorResponse {
	//er := errorResponseMap[e]
	//if er != ErrorResponse{} {
	//	return er

	//}
	return ErrorResponse{
		Code:           "errNotDefine",
		Description:    "No Define Error",
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
