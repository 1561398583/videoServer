package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type PdWriter struct {
	mu sync.Mutex
	dir string
	currentFileName string
	f *os.File
}

func NewPdWriter(dir string) *PdWriter {
	//获取日期
	timeStr := time.Now().String()
	//获取"year/month/day"
	dateStr := strings.Split(timeStr, " ")[0]
	filePath := dir + dateStr
	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		panic(err)
	}
	pdw := PdWriter{dir: dir, f: f, currentFileName: dateStr}
	//开启日期检测并替换文件
	go pdw.CheckDateAndNewFile()

	return &pdw
}


func (pdw *PdWriter) Write(p []byte) (n int, err error) {
	pdw.mu.Lock()
	defer pdw.mu.Unlock()
	n, err = pdw.f.Write(p)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//检查日期，如果是新的一天，那么就新建一个log文件，并把log文件设置为新文件
func (pdw *PdWriter) CheckDateAndNewFile()  {
	for  {
		time.Sleep(time.Minute * 1)	//1分钟检测一次
		//获取日期
		timeStr := time.Now().String()
		//获取"year/month/day"
		dateStr := strings.Split(timeStr, " ")[0]
		if dateStr != pdw.currentFileName{	//说明新的一天开始了
			filePath := pdw.dir + dateStr
			f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777 )
			if err != nil {
				panic(err)
			}
			oldFile := pdw.f
			//设置新的文件
			pdw.mu.Lock()
			pdw.f = f
			pdw.mu.Unlock()
			//关闭老的文件
			err = oldFile.Close()
			if err != nil {
				panic(err)
			}
		}

	}
}


