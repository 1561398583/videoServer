package midWare

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(400, c.Errors[0])
	}
}
