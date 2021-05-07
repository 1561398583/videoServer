package streamServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yx.com/videos/ServerConst"
)

func UploadVideo(c *gin.Context)  {
	/*
		name := c.PostForm("name")
		fmt.Println("name : " + name)
		description := c.PostForm("description")
		fmt.Println("description : " + description)
	*/

	form, err := c.MultipartForm()
	if err != nil{
		c.String(http.StatusBadRequest, "err : s%", err.Error())
		return
	}
	name := form.Value["name"][0]
	fmt.Println("name : " + name)
	description := form.Value["description"][0]
	fmt.Println("description : " + description)

	files := form.File["files"]
	for _, file := range files {
		dest := ServerConst.VIDEO_DIR + file.Filename
		if err := c.SaveUploadedFile(file, dest); err != nil{
			c.String(http.StatusBadRequest, "err : s%", err.Error())
		}
	}

	c.String(http.StatusOK, "upload sucess")
	return
}

