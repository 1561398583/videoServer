package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yx.com/videos/ServerConst"
	"yx.com/videos/db"
	"yx.com/videos/utils"
)


func GetLoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Main website",
	})
}

func Login(c *gin.Context)  {
	//获取用户名
	userName := c.PostForm("userName")
	faceImageUrl := c.PostForm("faceImageUrl")

	//查询数据库看用户名是否已经存在，若不存在则增加一个user
	user, _ := db.GetUserByUserName(userName)
	if user == nil {
		err := addUser(userName, faceImageUrl)
		if err != nil {
			Logger.Error(err)
			c.String(http.StatusBadRequest, "error :" + err.Error())
		}
	}
	c.String(http.StatusOK, "login sucess")
}

func addUser(userName, faceImageUrl string) error {
	//取得url ？之前的部分
	s1 := strings.Split(faceImageUrl, "?")[0]
	//取得文件名
	ss := strings.Split(s1, "/")
	faceImageName := ss[len(ss) - 1]
	//保存到服务器的path
	savePath := ServerConst.FACE_IMAGE_DIR + faceImageName
	//下载并保存faceImage
	err := utils.FetchFileAndSave(faceImageUrl, savePath)
	if err != nil {
		return errors.New("can not get face image")
	}

	//存入数据库,暂时把id存为userName
	user := db.User{
		UserName:userName,
		FaceImage:faceImageName,
	}

	err = db.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}
