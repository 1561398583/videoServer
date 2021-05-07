package db

func IsUserLikeVideo(userId, videoId string)  bool{
	ulv := UsersLikeVideo{}
	r := DB.Where("user_id=? AND video_id=?", userId, videoId).Find(&ulv)
	if r.Error != nil {
		panic(r.Error)
	}
	//没找到
	if ulv.ID == 0 {
		return false
	}
	return true
}

type UsersLikeVideo struct {
	ID int64
	UserId string
	VideoId string
}
