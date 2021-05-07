package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SAVEDIR = "/home/yx/Downloads/"
	HTMLDIR = "/home/yx/Videos/web/html/"
)

func UploadFile(c *gin.Context)  {
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
		dest := SAVEDIR + file.Filename
		if err := c.SaveUploadedFile(file, dest); err != nil{
			c.String(http.StatusBadRequest, "err : s%", err.Error())
		}
	}

	c.String(http.StatusOK, "upload sucess")
	return
}
