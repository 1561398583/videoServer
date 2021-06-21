package midWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yx.com/videos/api"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		resp := api.RespStruct{
			Status: "error",
		}
		errMsg := ""
		for _, err := range c.Errors {
			errMsg = errMsg + err.Err.Error() + " ; "
		}
		resp.ErrorMsg = errMsg
		c.JSON(http.StatusOK, resp)
	}
}
