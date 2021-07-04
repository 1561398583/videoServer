package client

import (
	"fmt"
	"testing"
	"time"
)

func TestTakeOut(t *testing.T) {
	s1 := "牛逼hehe<span><image src='https://www.abc.com'>hehe</image></span>abc"
	s := TakeOut(s1)
	fmt.Println(s)

	s2 := "牛逼hehe"
	s = TakeOut(s2)
	fmt.Println(s)
}

func TestModfiyComments(t *testing.T) {
	start := time.Now()
	offset := 0

	for {
		getNum := ModifyComments(offset)
		if getNum < 100 {
			break
		}
		offset += getNum
	}

	spendTime := time.Since(start)
	fmt.Println("all finish spend time : " + spendTime.String())

}