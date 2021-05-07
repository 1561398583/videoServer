package midWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
limit connect number
 */
type RequestLimiter struct {
	maxinum int
	token chan int
}

func RequestLimit(c *gin.Context){
	if limiter == nil {
		InitializeRequestLimiter(1000)
	}

	token := limiter.getToken()
	if !token {
		c.String(http.StatusBadRequest, "server busy")
		//fmt.Println("server busy")
		c.Abort()
		return
	}else {
		//getToken()和releaseToken()之间调用了其他func，如果其他func产生panic会导致releaseToken()得不到执行，导致资源泄露。
		//这里defer之后，就算其他func 产生panic，releaseToken()也能得到执行。
		defer limiter.releaseToken()
	}
	c.Next()
}

func InitializeRequestLimiter(limitNum int) {
	limiter = &RequestLimiter{
		maxinum: limitNum,
		token: make(chan int, limitNum),
	}
}

var limiter *RequestLimiter

func (cl *RequestLimiter) getToken() bool {
	if len(cl.token) >= cl.maxinum {
		return false
	}
	cl.token <- 1
	return true
}

func (cl *RequestLimiter) releaseToken() {
	<-cl.token
}

 