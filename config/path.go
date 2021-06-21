package config

import (
	"runtime"
)

var (
	LOG_DIR string
	HTML_DIR string
	ASSETS_DIR string
	VIDEO_DIR string
	FACE_IMAGE_DIR string
	VIDEO_PIC_DIR string    //video封面图
	REQUEST_MAX_NUM int = 1000
	HOST = "https://www.xiaoyxiao.cn"
)

func InitConfig()  {
	basePath := "/usr/videoProject"
	goos := runtime.GOOS
	if goos == "windows" {
		basePath = "E:/videoProject/server"
		HOST = "http://localhost:8080"
	}
	LOG_DIR = basePath + "/logs"
	HTML_DIR = basePath + "/html"
	ASSETS_DIR = basePath + "/assets"
	VIDEO_PIC_DIR = ASSETS_DIR + "/videoPic"
	VIDEO_DIR = ASSETS_DIR + "/videos"
	FACE_IMAGE_DIR = ASSETS_DIR + "/faceImg"

}
