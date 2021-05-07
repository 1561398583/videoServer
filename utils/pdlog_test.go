package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

//多个协程同时写一个loggo，看看是否出问题。
//因为loggo.Log是有锁的，所以应该没问题
func TestLog(t *testing.T)  {
	start := time.Now()
	logger := NewPdLogger("E:\\log\\test\\", "", LstdFlags, Debug)
	var wg sync.WaitGroup
	wg.Add(10)
	go writeLog("AAAAAAAAAA", logger, &wg)
	go writeLog("BBBBBBBBBB", logger, &wg)
	go writeLog("CCCCCCCCCC", logger, &wg)
	go writeLog("DDDDDDDDDD", logger, &wg)
	go writeLog("EEEEEEEEEE", logger, &wg)
	go writeLog("FFFFFFFFFF", logger, &wg)
	go writeLog("GGGGGGGGGG", logger, &wg)
	go writeLog("HHHHHHHHHH", logger, &wg)
	go writeLog("IIIIIIIIII", logger, &wg)
	go writeLog("JJJJJJJJJJ", logger, &wg)
	fmt.Println("begin 10 gouroutine")
	(&wg).Wait()
	fmt.Println("spend ", time.Since(start).String())
}

func writeLog(s string, pl *PdLog, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i < 1000; i++{
		pl.Debug(s)
	}
}


func TestCheckFile(t *testing.T)  {
	f, err := os.Open("E:\\log\\test\\2020-11-07")
	if err != nil {
		log.Fatal(err)
	}
	checkFile(f)
}

func checkFile(f *os.File)  {
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}
	s := string(bs)
	ss := strings.Split(s, "\n")
	for i := 0; i < len(ss); i++ {
		if len(ss[i]) < 10 {
			continue
		}
		c := isAllSame(strings.Split(ss[i], " ")[2])
		if c != 0 {
			fmt.Println(strconv.FormatInt(int64(i+1), 10), " row ",strconv.FormatInt(int64(c+1), 10), " culom")
			return
		}
	}
	fmt.Println("all ok")
}

func isAllSame(s string)  int{
	for i := 1; i < len(s); i++ {
		if s[i] != s[0]{
			return i
		}
	}
	return 0
}