package db

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

type Comment struct {
	ID        string `gorm:"primarykey"`
	CreatedAt string

	Level int

	RootId string	//上一级评论，如果是1级评论，那么rootid==id
	VideoId string
	Comment string
	UserId int64
	ToUserId string
	ChildNum int
	LikeNum int
}

func AddComments(cs []*Comment)  error{
	if cs == nil || len(cs) == 0 {
		return nil
	}
	r := DB.Create(cs)
	if r.Error != nil {
		if mysqlErr, ok := r.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil
			}
		}
		panic(r.Error)
	}
	return  nil
}

func AddAComment(comment *Comment)  error{
	if comment.ID == "" || comment.UserId == 0 || comment.VideoId == "" {
		return errors.New("id,userId,videoId can not be empty")
	}

	result := DB.Create(comment)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok{
			//id重复
			if mysqlErr.Number == 1062 {
				return errors.New("id has exist")
			}
		}
		//未知错误
		panic(result.Error)
	}
	return nil

}

//每页10条,按like数排序
func GetCommentsByVideoId(vId string, pageNum int) ([]*Comment, error) {
	 if vId == "" {
	 	return nil, errors.New("videoId can not be nil")
	 }
	 cs := make([]*Comment, 10)
	 r := DB.Where("video_id = ?", vId).Offset(pageNum * 10).Limit(10).Order("like_num desc").Find(&cs)
	 if r.Error != nil {
		 //未知错误
	 	panic(r.Error)
	 }
	 return cs, nil
}

//每页10条,按like数排序
func GetCommentsByFatherCommentId(fId string, pageNum int) ([]*Comment, error) {
	if fId == "" {
		return nil, errors.New("fatherCommentId can not be nil")
	}
	cs := make([]*Comment, 10)
	r := DB.Where("father_comment_id = ?", fId).Offset(pageNum * 10).Limit(10).Order("like_num desc").Find(&cs)
	if r.Error != nil {
		return nil, r.Error
	}
	return cs, nil
}



