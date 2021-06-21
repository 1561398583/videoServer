package db

import (
	"fmt"
	"math/rand"
	"regexp"
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

func TestBadComments(t *testing.T) {
	var commentNum int
	DB.Raw("SELECT count(*) FROM comments").Scan(&commentNum)

	fmt.Println(commentNum)

	offset := 0
	var comments []*Comment
	for offset < commentNum {
		// SELECT * FROM video_ns OFFSET offset LIMIT 10;
		DB.Limit(100).Offset(offset).Find(&comments)

		printBadComments(comments)

		offset += len(comments)
		comments = nil
	}
}

func printBadComments(comments []*Comment)  {
	for _, comment := range comments{
		if b := match(comment.Comment); b {
			println(comment.Comment)
		}
	}
}

//匹配 @xxx
func match(s string)  bool{
	r1, _ := regexp.MatchString("^@[\\s\\S\u4e00-\u9fa5]+$", s)
	if r1 {
		return true
	}
	r2, _ := regexp.MatchString("^[\\s\\S\u4e00-\u9fa5]*[<][\\s\\S]+[>][\\s\\S\u4e00-\u9fa5]*$", s)
	if r2 {
		return true
	}

	return false
}

func Test1(t *testing.T)  {
	//s1 := "哈哈<a href='http://t.cn/A6UKWauG' data-hide=''>are you ok?"
	//s2 := "@are you ok? '哈哈'"
	s3 := "are you ok?@"
	r := match(s3)
	fmt.Println(r)
}
