package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
	"yx.com/videos/config"
	"yx.com/videos/api"
	"yx.com/videos/midWare"
	"yx.com/videos/utils"
)


func main()  {
	//开启debugger
	StartHTTPDebuger()

	config.InitConfig()

	//gin engine
	r := gin.New()

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
			//params.Request.UserAgent(),
			//params.ErrorMessage,
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


	//捕获panic，并写入log
	frc, err := os.Create(config.LOG_DIR + "recover.log")
	if err != nil {
		log.Fatal(err)
	}
	r.Use(gin.RecoveryWithWriter(frc))
	r.Use(midWare.ErrorHandler)

	//test connect
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})
	})


	//InitErrorTest(r)
	//InitAbortTest(r)

	//限制访问数量
	midWare.InitializeRequestLimiter(1)
	r.Use(midWare.RequestLimit)

	r.GET("/testLimiter", func(c *gin.Context) {
		now := time.Now().String()
		fmt.Println("testLimiter : " + now)
		time.Sleep(time.Second * 30)
		c.String(http.StatusOK, "i can work")
	})



	//设置api error logger
	api.Logger = utils.NewPdLogger(config.LOG_DIR+ "apiError/", "", utils.LstdFlags, utils.Debug)

	//指定html文件path
	r.LoadHTMLGlob(config.HTML_DIR + "*")

	//设置静态资源
	r.StaticFS("/assets", http.Dir(config.ASSETS_DIR))

	//注册handlers
	api.RegistApi(r)




	r.GET("/uploadFilePage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "uploadFilePage.tmpl", gin.H{
			"title": "Main website",
		})
	})


	r.Run() // listen and serve on 0.0.0.0:8080 (for windowsConst "localhost:8080")
}

//开启一个http服务，监听":7890"，返回运行信息
func StartHTTPDebuger()  {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: ":7890", Handler: pprofHandler}
	go server.ListenAndServe()
}

