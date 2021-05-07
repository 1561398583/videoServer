package creeper

import (
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var cookieFilePath = "E:/go_project/videoProject/server/video/creeper/cookie"

var cookieMap map[string]string

func init()  {
	cookieMap = make(map[string]string)
	CookieFile2Map()
}

func UpdateCookie(cookies []string)  {
	//更新cookieMap
	SetCookie(cookies)
	//更新cookie file
	cookieStr := CookieMap2Str(cookieMap)
	cookieF, err := os.OpenFile(cookieFilePath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cookieF.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = cookieF.Write([]byte(cookieStr))
	if err != nil {
		panic(err)
	}
}

func getMWeiboCnParams(name string) string{
	mWeiboCnParams := cookieMap["m_weibocn_params"]
	mWeiboCnParams, err := url.QueryUnescape(mWeiboCnParams)
	if err != nil {
		panic(err)
	}
	m := String2Map(mWeiboCnParams)
	return m[name]
}

func CookieFile2Map()  {
	cookieF, err := os.OpenFile(cookieFilePath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cookieF.Close(); err != nil {
			panic(err)
		}
	}()

	bs, err := ioutil.ReadAll(cookieF)
	if err != nil {
		panic(err)
	}
	CookieStr2Map(string(bs))
}

func CookieMap2Str(cookieMap map[string]string)  string{
	r := ""
	for k, v := range cookieMap {
		r = r + k + "=" + v +"; "
	}
	return r
}

func CookieStr2Map(cookieStr string)  {
	if cookieStr == "" {
		return
	}
	kvStrS := strings.Split(cookieStr, "; ")
	for _, kvStr := range kvStrS {
		kv := strings.Split(kvStr, "=")
		if len(kv) != 2 {
			continue
		}
		k := kv[0]
		v := kv[1]
		cookieMap[k] = v
	}
}

//每一条set-cookie的格式： Set－Cookie: NAME=VALUE；Expires=DATE；Path=PATH；Domain=DOMAIN_NAME；SECURE，
//这里只需要NAME=VALUE
func SetCookie(sc []string)  {
	for _, c := range sc {
		nameValue := strings.Split(c, "; ")[0]
		nv := strings.Split(nameValue, "=")
		name := strings.ToLower(nv[0])
		value := nv[1]
		cookieMap[name] = value
	}
}


