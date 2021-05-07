package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"yx.com/videos/ServerConst"
	"yx.com/videos/db"
)

func GetVideos(c *gin.Context){
	sinceVideoId := c.Query("sinceVideoId")
	userId := c.Query("uid")
	numStr := c.Query("num")
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		_ = c.Error(errors.New("num param is error"))
		c.Abort()
		return
	}
	videos, err := db.GetVideosBySinceId(sinceVideoId, int(num))

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	if len(videos) == 0 {	//说明到达最后一个video了，那么就从头开始
		videos, err = db.GetVideosBySinceId(sinceVideoId, 5)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}
	}
	//转换struct
	clientVideos := make([]*VideoInfo2Client, len(videos))
	for i := 0; i < len(videos); i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = ServerConst.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		v2c.VideoSeconds = videos[i].VideoSeconds
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
	v2c.VideoUrl = ServerConst.HOST + "/" + "assets/videos/" + video.VideoFileName
	v2c.CommentNum = video.CommentNum
	v2c.LikeNum = video.LikeNum
	v2c.VideoSeconds = video.VideoSeconds
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
	v2c.VideoUrl = ServerConst.HOST + "/" + "assets/videos/" + video.VideoFileName
	v2c.CommentNum = video.CommentNum
	v2c.LikeNum = video.LikeNum
	v2c.VideoSeconds = video.VideoSeconds
	islike := db.IsUserLikeVideo(userId, video.ID)
	v2c.IsLike = islike

	c.JSON(200, v2c)
}

func GetStartVideos(c *gin.Context){
	userId := c.Query("uid")

	videos := make([]*db.Video, 3)
	video0 := db.GetFirstVideo()
	video1 := db.GetNextVideo(video0.ID)
	video2 := db.GetPreVideo(video0.ID)
	videos[0] = video0
	videos[1] = video1
	videos[2] = video2

	//转换struct
	clientVideos := make([]*VideoInfo2Client, 3)
	for i := 0; i < 3; i++ {
		v2c := new(VideoInfo2Client)
		v2c.ID = videos[i].ID
		v2c.VideoTitle = videos[i].VideoTitle
		v2c.VideoUrl = ServerConst.HOST + "/" + "assets/videos/" + videos[i].VideoFileName
		v2c.CommentNum = videos[i].CommentNum
		v2c.LikeNum = videos[i].LikeNum
		v2c.VideoSeconds = videos[i].VideoSeconds
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
	VideoSeconds float32
}
