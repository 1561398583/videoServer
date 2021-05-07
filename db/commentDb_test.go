package db

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAddAComment(t *testing.T) {
	c := Comment{
		Comment:"试一试",
		ID: "123",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		VideoId: "video123",
		UserId: int64(123),
	}
	err := AddAComment(&c)
	if err != nil {
		fmt.Printf("error type : %T \n", err)
		t.Errorf("%+v", err)
	}
}

func TestAddComments(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := strconv.FormatInt(int64(i), 10)
		likeNum := rand.Intn(100)
		c := Comment{
			Comment:"试一试",
			ID: id,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			VideoId: "video123",
			UserId: int64(123),
			LikeNum:likeNum,
		}
		err := AddAComment(&c)
		if err != nil {
			fmt.Printf("error type : %T \n", err)
			t.Errorf("%+v", err)
		}
	}

}

func TestFindComments(t *testing.T)  {
	cs, err := GetCommentsByVideoId("video123", 0)
	if err != nil {
		t.Errorf("%#v", err)
	}
	for _, c := range cs {
		fmt.Printf("id : %s, likeNum :%d \n",c.ID, c.LikeNum)
	}
	fmt.Printf("========================================\n")
	cs, err = GetCommentsByVideoId("video123", 1)
	if err != nil {
		t.Errorf("%#v", err)
	}
	for _, c := range cs {
		fmt.Printf("id : %s, likeNum :%d \n",c.ID, c.LikeNum)
	}
}

func TestDeleteComment(t *testing.T)  {
	r := DB.Where("id != ?", "123").Delete(&Comment{})
	if r.Error != nil {
		t.Error(r.Error)
	}
}
