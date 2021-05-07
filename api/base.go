package api

import (
	"github.com/gin-gonic/gin"
	"yx.com/videos/utils"
)

var Logger *utils.PdLog

func RegistApi(r *gin.Engine)  {
	//user api
	r.GET("/login", GetLoginPage)
	r.POST("/login", Login)

	//video api
	r.GET("/getVideos", GetVideos)
	r.GET("/getPreVideo", GetPreVideo)
	r.GET("/getNextVideo", GetNextVideo)
	r.GET("/getStartVideos", GetStartVideos)
}

type ResponseStruct struct {
	OK int //1:ok; 0:error
	Msg string
	Data interface{}
}
