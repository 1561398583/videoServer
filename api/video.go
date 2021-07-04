package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"yx.com/videos/config"
	"yx.com/videos/db"
)

func GetVideos(c *gin.Context){
	sinceVideoId := c.Query("sinceVideoId")
	userId := c.Query("uid")
	numStr := c.Query("num")
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		if e := c.Error(errors.New("num param is error")); e != nil {
			panic(e)
		}
		return
	}
	videos, err := db.GetVideoNsBySinceId(sinceVideoId, int(num))

	if err != nil {
		if e := c.Error(err); e != nil {
			panic(e)
		}
		return
	}

	//转换struct
	clientVideos := make([]*VideoInfo2Client, len(videos))
	for i := 0; i < len(videos); i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = config.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.PicUrl = config.HOST + "/" + "assets/videoPic/" + videos[i].PicFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		islike := db.IsUserLikeVideo(userId, videos[i].ID)
		v2c.IsLike = islike
		clientVideos[i] = v2c
	}
	c.JSON(200, clientVideos)
}

func GetMyLikeVideos(c *gin.Context){
	userId := c.Query("uid")
	offsetStr := c.Query("offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		if e := c.Error(errors.New("num param is error")); e != nil {
			panic(e)
		}
		return
	}
	ulvs, err := db.GetUserLikeVideosInfo(userId, int(offset))

	if err != nil {
		if e := c.Error(err); e != nil {
			panic(e)
		}
		return
	}

	//获取videos
	videoIds := make([]string, len(ulvs))
	for i := 0; i < len(ulvs); i++ {
		videoIds[i] = ulvs[i].VideoId
	}

	videos, err := db.GetVideoNByIds(videoIds)
	if err != nil {
		if e := c.Error(err); e != nil {
			panic(e)
		}
		return
	}

	//转换struct
	clientVideos := make([]*VideoInfo2Client, len(videos))
	for i := 0; i < len(videos); i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = config.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.PicUrl = config.HOST + "/" + "assets/videoPic/" + videos[i].PicFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		islike := db.IsUserLikeVideo(userId, videos[i].ID)
		v2c.IsLike = islike
		clientVideos[i] = v2c
	}
	c.JSON(200, clientVideos)
}

func GetPreVideo(c *gin.Context){
	sinceVideoId := c.Query("sinceVideoId")
	userId := c.Query("uid")

	video := db.GetPreVideo(sinceVideoId)

	//转换struct
	v2c := VideoInfo2Client{}
	v2c.ID = video.ID
	v2c.VideoTitle = video.VideoTitle
	v2c.VideoUrl = config.HOST + "/" + "assets/videos/" + video.VideoFileName
	v2c.PicUrl = config.HOST + "/" + "assets/videoPic/" + video.PicFileName
	v2c.CommentNum = video.CommentNum
	v2c.LikeNum = video.LikeNum
	islike := db.IsUserLikeVideo(userId, video.ID)
	v2c.IsLike = islike

	c.JSON(200, v2c)
}

func GetNextVideo(c *gin.Context){
	sinceVideoId := c.Query("sinceVideoId")
	userId := c.Query("uid")

	video:= db.GetNextVideo(sinceVideoId)

	//转换struct
	v2c := VideoInfo2Client{}
	v2c.ID = video.ID
	v2c.VideoTitle = video.VideoTitle
	v2c.VideoUrl = config.HOST + "/" + "assets/videos/" + video.VideoFileName
	v2c.PicUrl = config.HOST + "/" + "assets/videoPic/" + video.PicFileName
	v2c.CommentNum = video.CommentNum
	v2c.LikeNum = video.LikeNum
	islike := db.IsUserLikeVideo(userId, video.ID)
	v2c.IsLike = islike

	c.JSON(200, v2c)
}

func GetStartVideos(c *gin.Context){
	userId := c.Query("uid")

	videos, err := db.GetVideoNsBySinceId("0", 3)
	if err != nil {
		if e := c.Error(err); e != nil {
			panic(e)
		}
		return
	}

	//转换struct
	clientVideos := make([]*VideoInfo2Client, len(videos))
	for i := 0; i < len(videos); i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = config.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.PicUrl = config.HOST + "/" + "assets/videoPic/" + videos[i].PicFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		islike := db.IsUserLikeVideo(userId, videos[i].ID)
		v2c.IsLike = islike
		clientVideos[i] = v2c
	}
	c.JSON(200, clientVideos)
}


type VideoInfo2Client struct {
	ID        string `gorm:"primarykey"`
	VideoTitle string
	VideoUrl string
	LikeNum int
	IsLike bool
	CommentNum int
	PicUrl string
}
