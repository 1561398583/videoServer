package db

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func IsUserLikeVideo(userId, videoId string)  bool{
	ulv := UserLikeVideo{}
	r := DB.Where("user_id=? AND video_id=?", userId, videoId).First(&ulv)
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

func AddLike(userId, videoId string)  error{
	like := IsUserLikeVideo(userId, videoId)
	if like {	//已经like了，就不用再添加
		return nil
	}

	//添加一个like， 并且video的likeNum加1
	ulv := UserLikeVideo{}
	ulv.VideoId = videoId
	ulv.UserId = userId

	r := DB.Create(&ulv)

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

	//video likeNum 加1
	err := AddVideoLikeNum(videoId)
	if err != nil {
		panic(err)
	}

	return nil
}

func DeleteLike(userId, videoId string)  error{
	like := IsUserLikeVideo(userId, videoId)
	if !like {	//本来就不like了，就不用再删除了
		return nil
	}

	r := DB.Where("user_id=? AND video_id=?", userId, videoId).Delete(UserLikeVideo{})

	if r.Error != nil {
		//fmt.Printf("%#v", r.Error)	//查看error类型
		panic(r.Error)	//未知error
	}

	//video likeNum 减1
	err := MinusVideoLikeNum(videoId)
	if err != nil {
		panic(err)
	}

	return nil
}

func GetUserLikeVideosInfo(userId string, offset int)  ([]*UserLikeVideo,error){
	var ulvs []*UserLikeVideo
	r := DB.Where("user_id = ?", userId, ).Limit(10).Offset(offset).Find(&ulvs)
	if r.Error != nil {
		return nil, r.Error
	}
	return ulvs, nil
}



type UserLikeVideo struct {
	ID int64
	UserId string
	VideoId string
}
