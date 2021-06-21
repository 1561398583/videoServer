package creeper

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"yx.com/videos/db"
)

//已存在的user id缓存,用以减少数据库查询次数。
var userMap = make(map[string]bool)

//查询数据库，看user是否存在，若存在直接返回；否则创建一个
func createUser(user *User) error {
	//user是否已经存在
	if b, ok := userMap[strconv.FormatInt(int64(user.Id), 10)]; ok {
		if b == true {
			return nil
		}
	}
	u, err := db.GetUserById(strconv.FormatInt(int64(user.Id), 10))
	if err == nil && u != nil{	//用户已经存在，不需要再创建
		userMap[u.ID] = true
		return nil
	}
	//创建user
	//下载头像
	faceUrl := user.Avatar_hd
	imgName := GetFileNameFromUrl(faceUrl)
	savePath := FACE_IMAGE_DIR + "/" + imgName
	err = downLoadWeiBoImgAndSave(faceUrl, savePath)
	if err == FaceImgErrGif {	//如果头像是gif，就给他一个默认头像
		imgName = "001.jpg"
	}

	dbUser := db.User{
		ID:        strconv.FormatInt(int64(user.Id), 10),
		UserName:  user.ScreenName,
		PassWord:  "",
		FaceImage: imgName,
	}

	err = db.CreateUser(&dbUser)
	if err != nil {
		return errors.New("create user error : " + err.Error())
	}
	userMap[strconv.FormatInt(int64(user.Id), 10)] = true
	return nil
}

func downLoadFaceImgAndSave(imgUrl, savePath string) error{
	headers := make(map[string]string)
	headers["Accept"] = "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8"
	headers["Accept-Encoding"] = "gzip, deflate, br"
	headers["Accept-Language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["Referer"] = "https://m.weibo.cn/"

	resp, err := GetResponse(imgUrl, nil, headers)
	if err != nil {
		return errors.New("downLoadFaceImgAndSave error : " + err.Error())
	}

	if resp.Header["Content-Type"][0] == "image/gif" {
		return FaceImgErrGif
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	n, err := f.Write(body)
	if err != nil {
		panic(err)
	}

	if n == 0 {
		return errors.New("faceImg write 0 bytes")
	}

	return nil
}



var FaceImgErrGif = errors.New("is a gif")
