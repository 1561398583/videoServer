package utils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNewPdWriter(t *testing.T) {
	pdw := NewPdWriter("/home/yx/Videos/log/" + "pdlog/")
	pdw.Write([]byte("hello word"))
}

func TestNewFile(t *testing.T) {
	//获取日期
	//timeStr := time.Now().String()
	//获取"year/month/day"
	//dateStr := strings.Split(timeStr, " ")[0]
	filePath := "/home/yx/Videos/log/" + "pdlog/" + "20210215"
	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		log.Fatal(err)
	}
	n,err := f.Write([]byte("hello word"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}
