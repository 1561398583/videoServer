package db

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Video struct {
	ID        string `gorm:"primarykey"`
	CreateTime string
	VideoTitle string
	VideoFileName string
	VideoSeconds float32
	VideoWidth int
	VideoHeight int
	LikeNum int
	Status int
	CommentNum int
}

func AddVideo(video *Video) error {
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

func GetVideoById(id string) (*Video, error) {
	video := Video{}
	r := DB.First(&video, "id = ?", id)
	if r.Error != nil {
		//video 不存在
		if r.Error == gorm.ErrRecordNotFound{
			return nil, r.Error
		}
		//未知错误
		panic(r.Error)
	}
	return &video, nil
}

//返回从sinceId开始的num个video
func GetVideosBySinceId(sinceId string, num int) ([]*Video, error) {
	videos := make([]*Video, num)
	r := DB.Where("id > ?", sinceId).Limit(num).Find(&videos)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	getVideoNum := 0
	for _, v := range videos {
		if v != nil {
			getVideoNum ++
		}else {
			return nil, errors.New("not find enough video")
		}
	}
	return videos, nil
}

//Id上一个video
func GetPreVideo(id string) *Video {
	v := Video{}
	r := DB.Where("id < ?", id).Order("id DESC").Limit(1).Find(&v)
	if r.Error != nil {
		//未知错误
		panic(r.Error)
	}
	//没找到，很可能id是第一个id，返回最后一个，循环
	if v.ID == "" {
		vp := GetLastVideo()
		return vp
	}
	return &v
}

//Id下一个video
func GetNextVideo(id string) *Video {
	v := Video{}
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

func GetFirstVideo()  *Video{
	video := Video{}
	r := DB.Order("id").Limit(1).Find(&video)
	if r.Error != nil {
		panic(r.Error)
	}
	return &video
}

func GetLastVideo()  *Video{
	video := Video{}
	r := DB.Order("id DESC").Limit(1).Find(&video)
	if r.Error != nil {
		panic(r.Error)
	}
	return &video
}


