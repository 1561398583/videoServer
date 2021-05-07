package midWare

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func InitErrorTest(r *gin.Engine)  {
	//test panic
	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})
	//test error
	r.GET("/error", func(c *gin.Context) {
		err := errors.New("get a error")
		c.Error(err)
	})
}
