package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.String(http.StatusBadRequest, "parse multipart form data error : %s", err.Error())
		return
	}

	dest := form.Value["dest"][0]


	file := form.File["file"][0]

	if err := c.SaveUploadedFile(file, dest); err != nil{
		c.String(http.StatusBadRequest, "err : s%", err.Error())
	}


	c.String(http.StatusOK, "upload sucess")

	return
}

