package db

import (
	"fmt"
	"testing"
)

func TestGetVideosBySinceId(t *testing.T) {
	sinceId := "4646963578931551"
	videos, err := GetVideoNsBySinceId(sinceId, 10)
	if err != nil {
		t.Error(err)
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetFirstVideo(t *testing.T) {
	video := GetFirstVideo()
	fmt.Printf("%#v", *video)
}

func TestGetLastVideoVideo(t *testing.T) {
	video := GetLastVideo()
	fmt.Printf("%#v", *video)
}

func TestGetPreVideo(t *testing.T) {
	video := GetPreVideo("4547024495837368")
	fmt.Printf("%#v", *video)
}

func TestGetNextVideo(t *testing.T) {
	video := GetNextVideo("4547024495837368")
	fmt.Printf("%#v", *video)
}

//没有封面图片的加上默认的图片
func Test_addPic(t *testing.T)  {
	var videoNum int
	DB.Raw("SELECT count(*) FROM video_ns").Scan(&videoNum)

	offset := 0
	var videos []*VideoN
	for offset < videoNum {
		// SELECT * FROM video_ns OFFSET offset LIMIT 10;
		DB.Limit(10).Offset(offset).Find(&videos)

		addPic(videos)

		offset += len(videos)
		videos = nil
	}
}

func addPic(videos []*VideoN)  {
	noPicVideos := make([]*VideoN, 0)
	for _, video := range videos {
		if video.VideoTitle == "" {
			video.VideoTitle = "搞笑视频"
		}
		if video.PicFileName == ""{
			video.PicFileName = "001.jpg"
		}
		noPicVideos = append(noPicVideos, video)
	}

	for _, nvideo := range noPicVideos {
		DB.Save(nvideo)
	}
}

