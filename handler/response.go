package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleocheng/wormhole/api"
)

// WriteErrResponse ...
func WriteErrResponse(c *gin.Context, e error) {
	errResponse := api.GetErrorResponse(e)
	c.JSON(errResponse.HTTPStatusCode, errResponse)
}

// WriteSuccessResponse ...
func WriteSuccessResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}
