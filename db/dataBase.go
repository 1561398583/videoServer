package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"yx.com/videos/config"
)

var dsn = "wddlzh123:wddlmm123@Mysql@tcp(rm-wz96qo32w042vtyf2bo.mysql.rds.aliyuncs.com:3306)/video_info?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func init()  {
	//设置gorm log
	logFile, err := os.OpenFile(config.LOG_DIR+ "gorm_log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * 10,   // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	DB = db
}



