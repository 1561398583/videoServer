package config

import "runtime"

var (
	LOG_DIR string
	VIDEO_DIR string
	HTML_DIR string
	FACE_IMAGE_DIR string
	ASSETS_DIR string
	HOST = "http://121.5.72.78:8080"
)

func InitConfig()  {
	basePath := "/usr/videoProject/"
	goos := runtime.GOOS
	if goos == "windows" {
		basePath = "E:/videoProject/server/"
		HOST = "http://localhost:8080"
	}
	 LOG_DIR = basePath + "logs/"
	 VIDEO_DIR = basePath + "videos/"
	 HTML_DIR = basePath + "videoServer/html/"
	 FACE_IMAGE_DIR = basePath + "faceImg/"
	 ASSETS_DIR = basePath + "assets"
}
