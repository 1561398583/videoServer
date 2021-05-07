package streamServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	fmt.Println("login")
	userName := c.Query("userName")
	fmt.Println(userName)
}
