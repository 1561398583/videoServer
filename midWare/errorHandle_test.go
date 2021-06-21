package midWare

import (
	"errors"
	"github.com/gin-gonic/gin"
	"testing"
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

func TestErrHandler(t *testing.T)  {
	r := gin.Default()
	r.Use(ErrorHandler)
	r.GET("/error", func(c *gin.Context) {
		err1 := errors.New("get  error1")
		err2 := errors.New("get  error2")
		err3 := errors.New("get  error3")
		c.Error(err1)
		c.Error(err2)
		c.Error(err3)
		return
	})

	r.Run(":8080")
}
