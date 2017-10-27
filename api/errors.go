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
	errInvalidRegistry      = errors.New("Invalid Registry")
	errCreateRegistryFailed = errors.New("Create registry failed")
)

var errorResponseMap = map[error]ErrorResponse{
	errInvalidRegistry: {
		Code:           "errInvalidRegistry",
		Description:    "Invalid Registry",
		HTTPStatusCode: http.StatusBadRequest,
	},
	errCreateRegistryFailed: {
		Code:           "errCreateRegistryFailed",
		Description:    "Create registry failed, Check your parameter or try it later",
		HTTPStatusCode: http.StatusFailedDependency,
	},
}

// GetErrorResponse ...
func GetErrorResponse(e error) ErrorResponse {
	er := errorResponseMap[e]
	if er == (ErrorResponse{}) {
		return ErrorResponse{
			Code:           "errNotDefine",
			Description:    "No Define Error",
			HTTPStatusCode: http.StatusInternalServerError,
		}
	}
	return er
}
