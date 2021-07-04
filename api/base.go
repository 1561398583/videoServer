package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yx.com/videos/utils"
)

var Logger *utils.PdLog

func RegistApi(r *gin.Engine)  {
	//user api
	r.GET("/login", Login)
	r.GET("/getUserById", GetUserById)

	//video api
	r.GET("/getVideos", GetVideos)
	r.GET("/getMyLikeVideos", GetMyLikeVideos)
	r.GET("/getPreVideo", GetPreVideo)
	r.GET("/getNextVideo", GetNextVideo)
	r.GET("/getStartVideos", GetStartVideos)
	r.GET("/addLike", AddLike)
	r.GET("/deleteLike", DeleteLike)
	r.GET("/getLevel1Comments", GetLevel1Comments)
	r.GET("/getLevel2Comments", GetLevel2Comments)
	r.POST("/addL1Comment", AddL1Comment)
	r.GET("/addUserLikeComment", AddUserLikeComment)
	r.GET("/deleteUserLikeComment", DeleteUserLikeComment)


	r.MaxMultipartMemory = 8 << 24  //128M
	r.POST("/uploadFile", UploadFile)
	r.POST("/isFileExist", IsFileExcist)
	r.GET("/uploadFilePage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "uploadFilePage.tmpl", gin.H{
			"title": "Main website",
		})
	})
}

type RespStruct struct {
	Status string  //"ok" or "error"
	ErrorMsg string
	Data interface{}
}

