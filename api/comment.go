package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yx.com/videos/config"
	"yx.com/videos/db"
)

type RespComment struct {
	ID        string `gorm:"primarykey"`
	CreatedAt string

	Level int

	RootId string	//上一级评论，如果是1级评论，那么rootid==id
	VideoId string
	Comment string
	UserId string
	UserName string
	UserFaceImg string
	ToUserId string
	ChildNum int
	LikeNum int
	IsLike bool
}

func AddL1Comment(c *gin.Context){
	videoId := c.PostForm("videoId")
	userId := c.PostForm("uid")
	content := c.PostForm("content")
	fmt.Println(videoId)
	fmt.Println(userId)
	fmt.Println(content)
	newId, err := db.GetNewId()
	if err != nil {
		c.Error(err)
		return
	}
	comment := db.Comment{
		VideoId: videoId,
		UserId: userId,
		Comment: content,
		CreatedAt: time.Now().Format("2006-01-02"),
		Level: 1,
		ID: newId,
		RootId: newId,
		ChildNum: 0,
		LikeNum: 0,
	}
	err = db.AddAComment(&comment)
	if err != nil {
		c.Error(err)
		return
	}
	user, err := db.GetUserById(userId)
	if err != nil {
		c.Error(err)
		return
	}
	respC := RespComment{
		ID:comment.ID,
		CreatedAt : comment.CreatedAt,
		Level : 1,
		RootId : comment.ID,
		VideoId : comment.VideoId,
		Comment : comment.Comment,
		UserId : comment.UserId,
		UserName : user.UserName,
		UserFaceImg : config.HOST  + "/assets/faceImg/" + user.FaceImage,
		ToUserId : userId,
		ChildNum : comment.ChildNum,
		LikeNum : comment.LikeNum,
	}
	c.JSON(http.StatusOK, respC)

}

//获取10条评论，从第num个开始
func GetLevel1Comments(c *gin.Context){
	videoId := c.Query("videoId")
	numStr := c.Query("num")
	userId := c.Query("userId")

	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.Error(errors.New("num is error"))
		return
	}

	cs, err := db.GetLevel1CommentsByVideoId(videoId, num)
	if err != nil {
		c.Error(err)
		return
	}

	//添加用户信息
	ids := getUserIdsFromComments(cs)

	us, err := db.GetUsers(ids)
	if err != nil {
		c.Error(err)
		return
	}

	userMap := make(map[string]*db.User)

	for _, u := range us{
		userMap[u.ID] = u
	}

	//获取当前user的like信息
	cids := getCommentsId(cs)
	ulcs, err := db.UserLikeComments(userId, cids)
	if err != nil {
		c.Error(err)
		return
	}

	ulcMap := make(map[string]bool)
	for _, ulc := range ulcs{
		ulcMap[ulc.CommentId] = true
	}


	resp := make([]RespComment, len(cs))
	for i, c := range cs{
		resp[i].ID = c.ID
		resp[i].LikeNum = c.LikeNum
		resp[i].Comment = c.Comment
		resp[i].VideoId = c.VideoId
		resp[i].ChildNum = c.ChildNum
		resp[i].CreatedAt = c.CreatedAt
		resp[i].Level = c.Level
		resp[i].RootId = c.RootId
		resp[i].ToUserId = c.ToUserId

		user, ok := userMap[c.UserId]
		if !ok {	//user不存在
			user = &db.User{}
			user.UserName = "已注销"
			user.FaceImage = config.HOST  + "/assets/faceImg/" + config.DEFAULT_USER_FACE_IMG
		}

		resp[i].UserId = user.ID
		resp[i].UserName = user.UserName
		resp[i].UserFaceImg = config.HOST  + "/assets/faceImg/" + user.FaceImage

		if _,ok := ulcMap[c.ID]; ok {
			resp[i].IsLike = true
		}else {
			resp[i].IsLike = false
		}
	}

	c.JSON(http.StatusOK, resp)

}

func GetLevel2Comments(c *gin.Context){
	videoId := c.Query("rootId")
	numStr := c.Query("num")

	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.Error(errors.New("num is error"))
		return
	}

	cs, err := db.GetLevel2CommentsByRootId(videoId, num)
	if err != nil {
		c.Error(err)
		return
	}

	//添加用户信息
	ids := getUserIdsFromComments(cs)

	us, err := db.GetUsers(ids)
	if err != nil {
		c.Error(err)
		return
	}

	userMap := make(map[string]*db.User)

	for _, u := range us{
		userMap[u.ID] = u
	}

	resp := make([]RespComment, len(cs))
	for i, c := range cs{
		resp[i].ID = c.ID
		resp[i].LikeNum = c.LikeNum
		resp[i].Comment = c.Comment
		resp[i].VideoId = c.VideoId
		resp[i].ChildNum = c.ChildNum
		resp[i].CreatedAt = c.CreatedAt
		resp[i].Level = c.Level
		resp[i].RootId = c.RootId
		resp[i].ToUserId = c.ToUserId

		user, ok := userMap[c.UserId]
		if !ok {	//user不存在
			user = &db.User{}
			user.UserName = "已注销"
			user.FaceImage = config.HOST  + "/assets/faceImg/" + config.DEFAULT_USER_FACE_IMG
		}

		resp[i].UserId = user.ID
		resp[i].UserName = user.UserName
		resp[i].UserFaceImg = config.HOST  + "/assets/faceImg/" + user.FaceImage
	}

	c.JSON(http.StatusOK, resp)

}

//从所有评论中获取不重复的userId数组
func getUserIdsFromComments(cs []*db.Comment)  []string{
	//获取所有user的id,用map去重
	userIdMap := make(map[string]bool)
	for _, c := range cs{
		userIdMap[c.UserId] = true
	}

	ids := make([]string, len(userIdMap))
	i := 0
	for id := range userIdMap{
		ids[i] = id
		i++
	}

	return ids
}

func getCommentsId(cs []*db.Comment)  []string{
	ids := make([]string, len(cs))
	for i := 0; i < len(cs); i++ {
		ids[i] = cs[i].ID
	}
	return ids
}

func AddUserLikeComment(c *gin.Context){
	userId := c.Query("userId")
	commentId := c.Query("commentId")

	err := db.AddUserLikeComment(userId, commentId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func DeleteUserLikeComment(c *gin.Context){
	userId := c.Query("userId")
	commentId := c.Query("commentId")

	err := db.DeleteUserLikeComment(userId, commentId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}


