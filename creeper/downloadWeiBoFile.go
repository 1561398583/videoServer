package creeper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func downLoadWeiBoImgAndSave(imgUrl, savePath string) error{
	headers := make(map[string]string)
	headers["Accept"] = "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8"
	headers["Accept-Encoding"] = "gzip, deflate, br"
	headers["Accept-Language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["Referer"] = "https://m.weibo.cn/"

	resp, err := GetResponse(imgUrl, nil, headers)
	if err != nil {
		return errors.New("downLoadFaceImgAndSave error : " + err.Error())
	}else {
		defer func() {
			if err := resp.Body.Close();err != nil {
				panic(err)
			}
		}()
	}

	if resp.Header["Content-Type"][0] == "image/gif" {
		return FaceImgErrGif
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(savePath, body, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("write %d bytes to %s \n", len(body), savePath)
	return nil
}

func downLoadWeiBoVideoAndSave(videoUrl, savePath string) error{
	headers := make(map[string]string)
	headers["Accept"] = "*/*"
	headers["Accept-Encoding"] = "identity;q=1, *;q=0"
	headers["Accept-Language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["Connection"] = "keep-alive"
	headers["Host"] = "f.video.weibocdn.com"
	headers["Range"] = "bytes=0-"
	headers["Referer"] = "https://m.weibo.cn/"

	resp,err := GetResponse(videoUrl, nil, headers)
	if err != nil {
		return errors.New("downloadOneVideoAndSave : " + err.Error())
	}else {
		defer func() {
			if err := resp.Body.Close();err != nil {
				panic(err)
			}
		}()
	}

	if resp.Header["Content-Type"][0] != "video/mp4" {
		fmt.Println("want get a video , but not ")
		ShowResponse(resp)
		return errors.New("downLoadVideoAndSave : is not a mp4")
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

	fmt.Printf("write %d bytes to %s \n", n, savePath)

	return nil
}

