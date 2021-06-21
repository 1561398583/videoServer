package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yx.com/videos/db"
)

func AddLike(c *gin.Context)  {
	uid := c.Query("uid")
	videoId := c.Query("videoId")

	resp := RespStruct{}

	err := db.AddLike(uid, videoId)

	if err != nil {
		panic(err)
	}

	resp.Status = "ok"

	c.JSON(http.StatusOK, resp)
}

func DeleteLike(c *gin.Context)  {
	uid := c.Query("uid")
	videoId := c.Query("videoId")

	resp := RespStruct{}

	err := db.DeleteLike(uid, videoId)

	if err != nil {
		panic(err)
	}

	resp.Status = "ok"

	c.JSON(http.StatusOK, resp)
}


