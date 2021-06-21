package db

import (
	"fmt"
	"testing"
)

func TestIsUserLikeVideo(t *testing.T) {
	items := []struct{
		userId, videoId string
		r bool
	}{
		{"user1", "video1", true},
		{"123", "123", false},
		{"test_user_id1", "test_video_id1", false},
	}

	for _, item := range items{
		r := IsUserLikeVideo(item.userId, item.videoId)
		if r != item.r {
			t.Errorf("userId=%s ; videoId=%s, exception %t but %t\n", item.userId, item.videoId, item.r, r)
		}
	}
}

func TestAddLike(t *testing.T) {
	uid := "user1"
	videoId := "video2"
	err := AddLike(uid, videoId)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLike(t *testing.T) {
	uid := "user1"
	videoId := "video1"
	err := DeleteLike(uid, videoId)
	if err != nil {
		fmt.Printf("%#v \n", err)
	}
}
