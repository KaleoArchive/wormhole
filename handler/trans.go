package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleocheng/wormhole/api"
)

// NewTrans trans image from one registry to another
func NewTrans(c *gin.Context) {

	ts := &api.Trans{}
	c.Bind(ts)

	c.JSON(200, gin.H{
		"transId": "hello",
	})
}
