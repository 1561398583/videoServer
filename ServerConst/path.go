package ServerConst

import "runtime"

var (
	LOG_DIR string
	VIDEO_DIR string
	HTML_DIR string
	FACE_IMAGE_DIR string
	ASSETS_DIR string
	HOST = "http://localhost:8080"
)

func init()  {
	basePath := "/usr/videoProject/"
	goos := runtime.GOOS
	if goos == "windows" {
		basePath = "E:/videoProject/"
	}
	 LOG_DIR = basePath + "logs/"
	 VIDEO_DIR = basePath + "videos/"
	 HTML_DIR = basePath + "server/video/web/html/"
	 FACE_IMAGE_DIR = basePath + "faceImg/"
	 ASSETS_DIR = basePath + "assets"
}
