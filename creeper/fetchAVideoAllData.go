package creeper

import (
	"errors"
	"fmt"
	"time"
	"yx.com/videos/db"
)

func getAVideoAllData(card *Card)  error{
	fmt.Println("start get " + card.Mblog.Id)
	if card.Mblog.PageInfo.Type != "video" {
		return errors.New("is " + card.Mblog.PageInfo.Type + " not a video")
	}
	//video是否已经存在
	video, err := db.GetVideoNById(card.Mblog.Id)
	if err == nil && video != nil {
		return errors.New("video is existed, next")
	}

	var videoUrl string
	if card.Mblog.PageInfo.MediaInfo.StreamUrl != "" {
		videoUrl = card.Mblog.PageInfo.MediaInfo.StreamUrl
	}else if card.Mblog.PageInfo.MediaInfo.StreamUrlHd != "" {
		videoUrl = card.Mblog.PageInfo.MediaInfo.StreamUrlHd
	}else {
		return errors.New("video url is empty")
	}

	//获取video文件
	err = getAVideo(videoUrl)
	if err != nil {
		return errors.New("getAVideoAllData : " + err.Error())
	}

	//获取封面图
	err = getAVideoPic(card.Mblog.PageInfo.PagePic.Url)
	if err != nil {
		return errors.New("getAVideoAllData : " + err.Error())
	}

	/*
		获取评论
	*/
	wbId := card.Mblog.Id
	fmt.Println("start get " + wbId + " comments")
	cNum, err := GetVideoComments(wbId)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("get comments sucess %d / %d \n", cNum, card.Mblog.CommentsCount)
	}

	videoName := GetFileNameFromUrl(videoUrl)
	videoPicName := GetFileNameFromUrl(card.Mblog.PageInfo.PagePic.Url)

	//添加video
	newVideo := db.VideoN{
		ID:            card.Mblog.Id,
		CreateTime:    card.Mblog.CreatedAt,
		VideoTitle:    card.Mblog.PageInfo.Title,
		VideoFileName: videoName,
		LikeNum:       card.Mblog.AttitudesCount,
		CommentNum:    card.Mblog.CommentsCount,
		PicFileName:   videoPicName,
	}
	err = db.AddVideoN(&newVideo)
	if err != nil {
		return errors.New("getAVideoAllData : " + err.Error())
	}
	fmt.Println("get " + card.Mblog.Id + " sucess")

	time.Sleep(10 * time.Second)
	return nil
}


func getAVideoPic(url string)  error{
	if url == "" {
		return errors.New("videoPic url is empty")
	}

	videoPicName := GetFileNameFromUrl(url)
	savePath := VIDEO_PIC_DIR + "/" + videoPicName
	err := downLoadWeiBoImgAndSave(url, savePath)
	if err != nil {
		return errors.New("fetchVideoPicAndSave error : " + err.Error())
	}

	time.Sleep(5 * time.Second)

	return nil
}

func getAVideo(url string)  error{
	if url == "" {
		return errors.New("getAVideo ：video url is empty")
	}

	videoName := GetFileNameFromUrl(url)
	savePath := VIDEO_DIR + "/" + videoName
	err := downLoadWeiBoVideoAndSave(url, savePath)
	if err != nil {
		return errors.New("getAVideo : " + err.Error())
	}

	time.Sleep(5 * time.Second)

	return nil
}

