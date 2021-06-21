package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
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
		err1 := errors.New("parse multipart form data error : " + err.Error())
		c.Error(err1)
		return
	}

	dest := form.Value["dest"][0]


	file := form.File["file"][0]

	if err = c.SaveUploadedFile(file, dest); err != nil{
		err1 := errors.New("save file err : " + err.Error())
		c.Error(err1)
		return
	}


	c.String(http.StatusOK, "upload sucess")

	return
}

//文件是否已经存在
func IsFileExcist(c *gin.Context)  {
	body := c.Request.Body
	reqData, err := ioutil.ReadAll(body)
	if err != nil {
		c.Error(err)
		return
	}
	fp := fp{}
	err = json.Unmarshal(reqData, &fp)
	if err != nil {
		c.Error(err)
		return
	}

	fmt.Println(fp.FilePath)

	resp := RespJ{}

	_, err = os.Stat(fp.FilePath)
	if err != nil {
		resp.Status = "ok"
		resp.Data = "false"
	}else {
		resp.Status = "ok"
		resp.Data = "true"
	}

	c.JSON(http.StatusOK, resp)
}

type fp struct {
	FilePath string
}

type RespJ struct {
	Status string //ok; error
	ErrorMsg string
	Data string
}

