package midWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitAbortTest(engine *gin.Engine)  {
	engine.Use(Abort1)
	engine.Use(Abort2)
	engine.Use(Abort3)
	engine.GET("/abortTest", AbortTest)
}

func Abort1(c *gin.Context) {
	fmt.Println("start abort1")
	c.Next()
	fmt.Println("end abort1")
}

func Abort2(c *gin.Context) {
	fmt.Println("start abort2")
	c.Next()
	fmt.Println("end abort2")
}

func Abort3(c *gin.Context) {
	fmt.Println("start abort3")
	//c.Abort()
	c.Next()
	fmt.Println("end abort3")
}

func AbortTest(c *gin.Context) {
	fmt.Println("start AbortTest")
	fmt.Println("end AbortTest")
	return
}
