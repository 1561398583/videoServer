package db

import (
	"testing"
)

func TestIsUserLikeVideo(t *testing.T) {
	items := []struct{
		userId, videoId string
		r bool
	}{
		{"123", "123", false},
		{"test_user_id1", "test_video_id1", true},
	}

	for _, item := range items{
		r := IsUserLikeVideo(item.userId, item.videoId)
		if r != item.r {
			t.Errorf("userId=%s ; videoId=%s, exception %t but %t\n", item.userId, item.videoId, item.r, r)
		}
	}
}
