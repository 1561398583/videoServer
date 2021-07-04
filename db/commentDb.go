package db

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Comment struct {
	ID        string `gorm:"primarykey"`
	CreatedAt string

	Level int

	RootId string	//上一级评论，如果是1级评论，那么rootid==id
	VideoId string
	Comment string
	UserId string

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
	if comment.ID == ""{
		return errors.New("db.AddAComment error : id can not be empty")
	}
	if comment.UserId == ""{
		return errors.New("db.AddAComment error : userId can not be empty")
	}
	if comment.VideoId == "" {
		return errors.New("db.AddAComment error : videoId can not be empty")
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

//获取10条level1评论，从num号开始，按like数排序
func GetLevel1CommentsByVideoId(vId string, num int) ([]*Comment, error) {
	 var cs []*Comment
	 r := DB.Where("video_id = ? AND level=1", vId).Offset(num).Limit(10).Order("like_num desc").Find(&cs)
	 if r.Error != nil {
		 //未知错误
	 	panic(r.Error)
	 }
	 return cs, nil
}

//获取10条2级评论，从num号开始
func GetLevel2CommentsByRootId(rootId string, num int) ([]*Comment, error) {
	var cs []*Comment
	r := DB.Where("root_id = ? AND level=2", rootId).Offset(num).Limit(10).Order("like_num desc").Find(&cs)
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


func GetNewId()  (string,error){
	comment := Comment{}
	tx := DB.Last(&comment)
	if tx.Error != nil {
		return "", tx.Error
	}
	lastIdStr := comment.ID
	lastId, err := strconv.ParseInt(lastIdStr, 10, 64)
	if err != nil {
		return "", err
	}
	newId := lastId + 1
	newIdStr := strconv.FormatInt(newId, 10)
	return newIdStr, nil
}

func IsUserLikeComment(userId, commentId string) bool {
	ulc := UserLikeComment{}
	r := DB.Where("user_id=? AND comment_id=?", userId, commentId).First(&ulc)
	//fmt.Printf("%#v \n", ulv)
	if r.Error != nil {
		//fmt.Printf("%#v \n", r.Error)
		if r.Error == gorm.ErrRecordNotFound {
			return false
		}

		panic(r.Error)	//未知error
	}

	return true
}

func AddUserLikeComment(userId, commentId string)  error{
	ulc := UserLikeComment{
		UserId: userId,
		CommentId: commentId,
	}
	r := DB.Create(&ulc)

	if r.Error != nil {
		//fmt.Printf("%#v", r.Error)	//查看error类型
		if e, ok := r.Error.(*mysql.MySQLError); ok {
			if e.Number == 0x426{	//重复,已经添加like了，直接返回即可
				return nil
			}else {
				panic(r.Error)	//未知error
			}
		}else {
			panic(r.Error)	//未知error
		}

	}

	//comment likeNum + 1
	tx := DB.Exec("UPDATE comments SET like_num=like_num+1 WHERE id = ?", commentId)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DeleteUserLikeComment(userId, commentId string) error {
	r := DB.Where("user_id=? AND comment_id=?", userId, commentId).Delete(UserLikeComment{})

	if r.Error != nil {
		//fmt.Printf("%#v", r.Error)	//查看error类型
		panic(r.Error)	//未知error
	}

	//comment likeNum - 1
	tx := DB.Exec("UPDATE comments SET like_num=like_num-1 WHERE id = ?", commentId)
	if tx.Error != nil {
		return tx.Error
	}

	return nil

}

func UserLikeComments(userId string, commentsId []string)  ([]*UserLikeComment, error){
	var ulcs []*UserLikeComment
	tx := DB.Where("user_id = ? AND comment_id IN ?", userId, commentsId).Find(&ulcs)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return ulcs, nil
}



type UserLikeComment struct {
	Id int64
	UserId string
	CommentId string
}


//获取100条level1评论，从offset号开始
func GetLevel1Comments(offset int) ([]*Comment, error) {
	var cs []*Comment
	r := DB.Where("level=1").Offset(offset).Limit(100).Find(&cs)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	return cs, nil
}

func UpdateComment(c *Comment)  error{
	tx := DB.Model(&Comment{}).Where("id = ?", c.ID).Update("comment", c.Comment)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}


