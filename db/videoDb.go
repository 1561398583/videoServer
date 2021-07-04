package db

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)


type VideoN struct {
	ID        string `gorm:"primarykey"`
	CreateTime string
	VideoTitle string
	VideoFileName string
	LikeNum int
	CommentNum int
	PicFileName string
}

func AddVideoN(video *VideoN) error {
	r := DB.Create(video)
	if r.Error != nil {
		if mysqlErr, ok := r.Error.(*mysql.MySQLError); ok{
			//id重复
			if mysqlErr.Number == 1062 {
				return mysqlErr
			}
		}
		//未知错误
		panic(r.Error)
	}
	return nil
}


func GetVideoNById(id string) (*VideoN, error) {
	video := VideoN{}
	r := DB.First(&video, "id = ?", id)
	if r.Error != nil {
		//video 不存在
		if r.Error == gorm.ErrRecordNotFound{
			return nil, NotFindErr
		}
		//未知错误
		panic(r.Error)
	}
	return &video, nil
}


func GetVideoNByIds(ids []string) ([]*VideoN, error) {
	var videos []*VideoN
	r := DB.Where("id IN ?", ids).Find(&videos)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	return videos, nil
}



//返回从sinceId开始的num个video
func GetVideoNsBySinceId(sinceId string, num int) ([]*VideoN, error) {
	var videos []*VideoN
	r := DB.Where("id > ?", sinceId).Limit(num).Find(&videos)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	if len(videos) < num {	//说明到达最后一个video了，那么就从第一个video循环
		var videos1 []*VideoN
		r := DB.Limit(num - len(videos)).Offset(0).Find(&videos1)
		if r.Error != nil {
			//未知错误
			panic(r.Error)
		}
		if len(videos1) != (num - len(videos)) {
			panic("get video num error")
		}
		videos = append(videos, videos1...)
	}

	return videos, nil
}




//Id上一个video
func GetPreVideo(id string) *VideoN {
	var video VideoN
	r := DB.Where("id < ?", id).Order("id DESC").Limit(1).Find(&video)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	//没找到，很可能id是第一个id，返回最后一个，循环
	if video.ID == "" {
		vp := GetLastVideo()
		return vp
	}
	return &video
}

//Id下一个video
func GetNextVideo(id string) *VideoN {
	v := VideoN{}
	r := DB.Where("id > ?", id).Order("id ASC").Limit(1).Find(&v)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	//没找到，很可能id是最后一个，那么就返回第一个，用以循环
	if v.ID == "" {
		vp := GetFirstVideo()
		return vp
	}

	return &v
}

func GetFirstVideo()  *VideoN{
	video := VideoN{}
	r := DB.Order("id").Limit(1).Find(&video)
	if r.Error != nil {
		panic(r.Error)
	}
	if video.ID == "" {
		panic("video id is enpty")
	}
	return &video
}

func GetLastVideo()  *VideoN{
	var video VideoN
	r := DB.Order("id DESC").Limit(1).Find(&video)
	if r.Error != nil {
		panic(r.Error)
	}
	if video.ID == "" {
		panic("video id is enpty")
	}
	return &video
}

//likeNum加1
func AddVideoLikeNum(videoId string) error {
	video, err := GetVideoNById(videoId)
	if err != nil {
		return err
	}

	video.LikeNum += 1

	r := DB.Save(&video)
	if r.Error != nil {
		panic(err)
	}

	return nil
}

//likeNum减1
func MinusVideoLikeNum(videoId string) error {
	video, err := GetVideoNById(videoId)
	if err != nil {
		return err
	}

	if video.LikeNum > 0 {
		video.LikeNum -= 1
	}

	r := DB.Save(&video)
	if r.Error != nil {
		panic(err)
	}

	return nil
}