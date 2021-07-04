package client

import (
	"fmt"
	"strings"
	"time"
	"yx.com/videos/db"
)

//去掉<xxx>
func TakeOut(s string) string {
	index1 := strings.Index(s, "<")
	index2 := strings.LastIndex(s, ">")
	if index1 == -1 {
		return s
	}


	ns := ""
	ns += s[0:index1] + s[index2+1:]

	return ns
}

func ModifyComments(offset int)  (getCommentNum int){
	startTime := time.Now()
	cs, err := db.GetLevel1Comments(offset)
	if err != nil {
		panic(err)
	}
	getCommentNum = len(cs)
	//原本每个comment的length
	ls := make([]int, len(cs))
	for i := 0; i < len(cs); i++ {
		ls[i] = len(cs[i].Comment)
	}

	//修改comment
	for i := 0; i < len(cs); i++ {
		ns := TakeOut(cs[i].Comment)
		if ns == "" {
			ns = "哈哈"
		}
		/*
		if len(ns) != len(cs[i].Comment) {
			fmt.Println(cs[i].Comment)
			fmt.Println("=>")
			fmt.Println(ns)
		}

		 */
		cs[i].Comment = ns
	}


	modifyNum := 0
	//数据库修改
	for i := 0; i < len(cs); i++ {
		if len(cs[i].Comment) != ls[i] {
			err := db.UpdateComment(cs[i])
			if err != nil {
				panic(err)
			}
			modifyNum ++
		}
	}

	spendTime := time.Since(startTime)

	fmt.Printf("modify %d / %d\n", modifyNum, len(cs))
	fmt.Println("spend time : " + spendTime.String())

	return
}
