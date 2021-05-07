module yx.com/videos

go 1.15

require (
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/now v1.1.2 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/text v0.3.2
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.21.4
	github.com/google/uuid v1.1.2
)

//防火墙原因，这些包从github下载
replace (
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 => github.com/golang/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/text/encoding v0.3.2 => github.com/golang/text v0.3.2
)
