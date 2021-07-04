package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"yx.com/videos/config"
	"yx.com/videos/creeper"
	"yx.com/videos/db"
	"yx.com/videos/utils"
)


func GetLoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Main website",
	})
}


func Login(c *gin.Context)  {
	//获取临时code
	code := c.Query("code")
	NickName := c.Query("NickName")
	FaceImgUrl := c.Query("FaceImgUrl")
	//fmt.Println("code : " + code)

	result := LoginResult{}

	//去微信服务器上获取openid
	openId, err := getOpenid(code)
	if err != nil {
		result.Status = "error"
		result.OpenId = ""
		result.ErrMsg = "get openId err : " + err.Error()
		c.JSON(http.StatusOK, result)
		return
	}

	//查询数据库看用户名是否已经存在，若不存在则用传过来用户信息进行注册
	user, err := db.GetUserById(openId)
	if err == db.NotFindErr || user == nil {
		err := addWeiXinUser(openId, NickName, FaceImgUrl)
		if err != nil {
			result.Status = "error"
			result.OpenId = ""
			result.ErrMsg = "addUser error : " + err.Error()
		}else {
			result.Status = "ok"
			result.OpenId = openId
		}
	}else {
		result.Status = "ok"
		result.OpenId = openId
	}

	c.JSON(http.StatusOK, result)
}


func GetUserById(c *gin.Context)  {
	uid := c.Query("uid")

	result := RespStruct{}

	//查询数据库看用户名是否已经存在，若不存在则用传过来用户信息进行注册
	user, err := db.GetUserById(uid)
	if err == db.NotFindErr || user == nil {
		result.Status = "error"
		result.ErrorMsg = "not find this user"
	}else {
		result.Status = "ok"
		result.Data = user
	}

	c.JSON(http.StatusOK, result)
}




/*
func Register(c *gin.Context)  {
	OpenId := c.Query("OpenId")
	NickName := c.Query("NickName")
	FaceImgUrl := c.Query("FaceImgUrl")
	//查询数据库看用户名是否已经存在，若不存在则增加一个user
	user, err := db.GetUserById(OpenId)
	if err == db.NotFindErr {
		err := addUser(OpenId, NickName, FaceImgUrl)
		if err != nil {
			Logger.Error(err)
			c.String(http.StatusBadRequest, "error :" + err.Error())
		}
	}
	c.String(http.StatusOK, "login sucess")
}

 */

func addWeiBoUser(userId, userName, faceImageUrl string) error {
	//取得url ？之前的部分
	s1 := strings.Split(faceImageUrl, "?")[0]
	//取得文件名
	ss := strings.Split(s1, "/")
	faceImageName := ss[len(ss) - 1]
	//保存到服务器的path
	savePath := config.FACE_IMAGE_DIR + "/" + faceImageName
	//下载并保存faceImage
	err := utils.FetchFileAndSave(faceImageUrl, savePath)
	if err != nil {
		return errors.New("fetch faceImage err : " + err.Error())
	}

	//存入数据库,暂时把id存为userName
	user := db.User{
		ID: userId,
		UserName : userName,
		FaceImage : faceImageName,
	}

	err = db.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func addWeiXinUser(userId, userName, faceImageUrl string) error {
	//取得url ？之前的部分
	s1 := strings.Split(faceImageUrl, "?")[0]
	//取得文件名
	ss := strings.Split(s1, "/")
	faceImageName := ss[len(ss) - 1]
	//保存到服务器的path
	savePath := config.FACE_IMAGE_DIR + "/" + faceImageName
	//下载并保存faceImage
	//err := utils.FetchFileAndSave(faceImageUrl, savePath)
	err := creeper.DownloadFileAndSave(faceImageUrl, savePath, nil)
	if err != nil {
		return errors.New("fetch faceImage err : " + err.Error())
	}

	//存入数据库
	user := db.User{
		ID: userId,
		UserName : userName,
		FaceImage : faceImageName,
	}

	err = db.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func getOpenid(code string)  (openId string, err error){
	url := "https://api.weixin.qq.com/sns/jscode2session"

	client := http.DefaultClient

	urlParams := make(map[string]string)

	urlParams["appid"] = config.APP_ID
	urlParams["secret"] = config.SECRET
	urlParams["js_code"] = code
	urlParams["grant_type"] = "authorization_code"


	//组装url
	reqUrl := creeper.UrlAddParams(url, urlParams)

	req,err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}

	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("get openId error : " + err.Error())
	}else {
		defer func() {
			if e := resp.Body.Close(); e != nil {
				panic(e)
			}
		}()
	}

   // creeper.ShowResponse(resp)
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(respData))
	wxL := WxLogin{}
	err = json.Unmarshal(respData, &wxL)
	if err != nil {
		return "", err
	}

	if wxL.ErrMsg != "" {
		return "", errors.New(wxL.ErrMsg)
	}

	return wxL.OpenId, nil
}

type WxLogin struct {
	ErrCode int	`json:"errcode"`
	ErrMsg string	`json:"errmsg"`
	Session_key string	`json:"session_key"`
	OpenId  string	`json:"openid"`
}

type LoginResult struct {
	Status string	//"ok" or "register"
	OpenId string	//openid
	ErrMsg string
}

type WxRegister struct {
	OpenId string
	NickName string
	FaceImgUrl string
}

