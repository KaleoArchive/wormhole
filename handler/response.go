package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleocheng/wormhole/api"
)

// WriteErrResponse ...
func WriteErrResponse(c *gin.Context, e error) {
	errResponse := api.GetErrorResponse(e)
	c.JSON(errResponse.HTTPStatusCode, errResponse)
}
