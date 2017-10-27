package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleocheng/wormhole/api"
)

// NewRegistry ...
func NewRegistry(c *gin.Context) {
	rs := &api.Registry{}
	c.Bind(rs)

	id, err := api.NewRegistry(rs)
	if err != nil {
		WriteErrResponse(c, err)
		return
	}
	WriteSuccessResponse(c, gin.H{
		"registryId": id,
	})
}
