package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	"time"
	"yx.com/videos/api"
	"yx.com/videos/config"
	"yx.com/videos/midWare"
	"yx.com/videos/utils"
)


func main()  {
	//开启debugger
	StartHTTPDebuger()

	config.InitConfig()

	//gin engine
	r := gin.Default()
	//default默认开启recover，并且panic信息写入DefaultErrorWriter(默认是Stderr)，这里设置为写入文件中
	recoverFile, err := os.OpenFile(config.LOG_DIR+ "/recover.log", os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultErrorWriter = recoverFile
	//gin log写入DefaultWriter（默认Stdout），改为写入文件
	ginLogFile, err := os.OpenFile(config.LOG_DIR+ "/gin.log", os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = ginLogFile

	//统一错误处理
	r.Use(midWare.ErrorHandler)

	//test connect
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})
	})
	//test panic
	r.GET("/panic", func(c *gin.Context) {
		panic("get a panic")
	})


	//InitErrorTest(r)
	//InitAbortTest(r)

	//限制访问数量
	midWare.InitializeRequestLimiter(config.REQUEST_MAX_NUM)
	r.Use(midWare.RequestLimit)

	r.GET("/testLimiter", func(c *gin.Context) {
		now := time.Now().String()
		fmt.Println("testLimiter : " + now)
		time.Sleep(time.Second * 30)
		c.String(http.StatusOK, "i can work")
	})



	//设置api error logger
	api.Logger = utils.NewPdLogger(config.LOG_DIR+ "/apiError/", "", utils.LstdFlags, utils.Debug)

	//指定html文件path
	r.LoadHTMLGlob(config.HTML_DIR + "/*")

	//设置静态资源
	r.StaticFS("/assets", http.Dir(config.ASSETS_DIR))

	//注册handlers
	api.RegistApi(r)

	goos := runtime.GOOS
	if goos == "windows" {
		r.Run() // listen and serve on 0.0.0.0:8080 (for windowsConst "localhost:8080")
	}else {	//linux
		r.RunTLS(":443", config.TLS_DIR + "/xiaoyxiao.cn.pem", config.TLS_DIR + "/xiaoyxiao.cn.key")
	}



	/*
	fmt.Println("start server 2021.6.14 15:34")

	err = r.RunTLS(":443", config.TLS_DIR + "/xiaoyxiao.cn.pem", config.TLS_DIR + "/xiaoyxiao.cn.key")
	if err != nil {
		panic(err)
	}
	*/
}

//开启一个http服务，监听":7890"，返回运行信息
func StartHTTPDebuger()  {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: ":7890", Handler: pprofHandler}
	go server.ListenAndServe()
}

func setLog(r *gin.Engine) {
	//设置日志
	logFormatterFunc := func(params gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}
	logWriter, err := os.OpenFile(config.LOG_DIR+ "gin.log", os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		log.Fatal(err)
	}
	loggerConfig := gin.LoggerConfig{
		Formatter: logFormatterFunc,
		Output: logWriter,
	}
	logger := gin.LoggerWithConfig(loggerConfig)
	r.Use(logger)
}

