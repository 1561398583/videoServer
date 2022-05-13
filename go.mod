module yx.com/videos

go 1.15

require (
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/validator/v10 v10.2.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/unrolled/secure v1.0.9
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/text v0.3.2
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.21.4
)

//防火墙原因，这些包从github下载
replace (
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 => github.com/golang/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/text/encoding v0.3.2 => github.com/golang/text v0.3.2
	google.golang.org/protobuf v1.27.1 => github.com/protocolbuffers/protobuf-go v1.27.1
)
