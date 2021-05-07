package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//新开一个协程，每天检查日期，如果是新的一天就新建一个log文件，然后log写入新log文件。
/**
dir:log file 目录，每天在这个目录中产生新的log文件
prefix:每一行log的前缀
level:记录log的级别
 */
func NewPdLogger(dir string, prefix string, flag int,  level LogLevel) *PdLog {
	//获取日期
	timeStr := time.Now().String()
	//获取"year/month/day"
	dateStr := strings.Split(timeStr, " ")[0]
	logFileName := dir + dateStr
	logFile, err := os.OpenFile(logFileName, os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, prefix, flag)
	pl := &PdLog{log: logger, level: level, dir: dir, currentLogFile: logFile, dateStr: dateStr}

	//开启日期检测并替换log文件
	go pl.CheckDateAndNewFile()

	return pl
}


// LogLevel
type LogLevel int

const (
	Debug	LogLevel = iota + 1
	Info
	Warn
	Error
)

//flag
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)


//会有多个协程对log进行写，而*log.Logger是并发安全的（自带锁）
//level只有读操作
//dateStr,logPath,currentLogFile只有CheckDateAndNewFile()一个协程写
//综上，Loggo struct不用加锁
type PdLog struct {
	log *log.Logger
	level LogLevel
	dateStr string	//当前日志文件名（日期）
	dir string	//log文件夹
	currentLogFile *os.File  //当前log文件
}

func (pl *PdLog) Debug(e interface{})  {
	if pl.level <= Debug {
		var s string
		//判断最常用的string和error，免得每次都fmt.Sprint(e)反射得到类型，反射时间消耗多
		switch e := e.(type) {
		case string:
			s = e
		case error:
			s = e.Error()
		default:
			s = fmt.Sprint(e)
		}
		pl.log.Println("\n" + "DEBUG\n" + s + "\n")
	}
}

func (pl *PdLog) Info(e interface{})  {
	if pl.level <= Info {
		var s string
		//判断最常用的string和error，免得每次都fmt.Sprint(e)反射得到类型，反射时间消耗多
		switch e := e.(type) {
		case string:
			s = e
		case error:
			s = e.Error()
		default:
			s = fmt.Sprint(e)
		}
		pl.log.Println("\n" + "INFO\n" + s + "\n")
	}
}

func (pl *PdLog) Warn(e interface{})  {
	if pl.level <= Warn {
		var s string
		//判断最常用的string和error，免得每次都fmt.Sprint(e)反射得到类型，反射时间消耗多
		switch e := e.(type) {
		case string:
			s = e
		case error:
			s = e.Error()
		default:
			s = fmt.Sprint(e)
		}
		pl.log.Println("\n" + "WARN\n" + s + "\n")
	}
}

func (pl *PdLog) Error(e interface{})  {
	if pl.level <= Error {
		var s string
		//判断最常用的string和error，免得每次都fmt.Sprint(e)反射得到类型，反射时间消耗多
		switch e := e.(type) {
		case string:
			s = e
		case error:
			s = e.Error()
		default:
			s = fmt.Sprint(e)
		}
		pl.log.Println("\n" + "ERROR\n" + s + "\n")
	}
}

//检查日期，如果是新的一天，那么就新建一个log文件，并把log文件设置为新文件
func (pl *PdLog) CheckDateAndNewFile()  {
	for  {
		time.Sleep(time.Minute * 1)	//1分钟检测一次
		//获取日期
		timeStr := time.Now().String()
		//获取"year/month/day"
		dateStr := strings.Split(timeStr, " ")[0]
		if dateStr != pl.dateStr{	//说明新的一天开始了
			logFileName := pl.dir + dateStr
			logFile, err := os.OpenFile(logFileName, os.O_APPEND | os.O_CREATE, 0777 )
			if err != nil {
				panic(err)
			}
			//设置新的log文件
			pl.log.SetOutput(logFile)
			//关闭老的log文件
			err = pl.currentLogFile.Close()
			if err != nil {
				panic(err)
			}
			//loggo持有新的log文件
			pl.currentLogFile = logFile
		}

	}
}


