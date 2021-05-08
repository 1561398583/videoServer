package creeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"yx.com/videos/config"
	"yx.com/videos/db"
)

//下面这些变量从浏览器获取，访问主页。
var(
	//用户id
	uid = "2640113513"
	t = "0"
	//主页容器id
	containerid = "1005052640113513"

)


//登录需要手机验证，代码实现太麻烦，所以先用浏览器登录，再把cookie拿来用
//访问被拒绝才需要调用这个func
//一般直接运行即可，cookie被保存在cookie文件中，直接访问是从cookie文件中读取
func InitCookie()  {
	originalCookieStr := "WEIBOCN_WM=3349; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WhTzBIOG_4RuJlkY5Bwk6eq5NHD95Qfeoz0eh271hnEWs4DqcjMi--NiK.Xi-2Ri--ciKnRi-zNSKzEe05pehnRentt; WEIBOCN_FROM=1110106030; SCF=AuWDb1t_wLpDoWn1Hwt_z37r6O_pXqvOLwuWdI6mN2mEF2zUA5JKVWGGiRmII3DFb0fJ27Av7a3Xmt2ZdQyzX9k.; SUB=_2A25NaMd6DeRhGeNM6VER8S7Ewz6IHXVukukyrDV6PUJbktANLVTQkW1NTiOp8Vab4DpZ3ELCtMSP6aKsNeVd1ICE; SSOLoginState=1617737514; _T_WM=72862240042; MLOGIN=1; XSRF-TOKEN=d4e984; M_WEIBOCN_PARAMS=luicode%3D10000011%26lfid%3D100103type%253D17%2526q%253D%25E6%2590%259E%25E7%25AC%2591%2526t%253D0%26fid%3D1076032640113513%26uicode%3D10000011"
	newCookies := []string{
		"XSRF-TOKEN=deleted; expires=Thu, 01-Jan-1970 00:00:01 GMT; Max-Age=0; path=/; domain=.weibo.cn; HttpOnly",
		"XSRF-TOKEN=d4e984; expires=Fri, 23-Apr-2021 13:52:00 GMT; Max-Age=1200; path=/; domain=m.weibo.cn",
		"MLOGIN=1; expires=Fri, 23-Apr-2021 14:32:00 GMT; Max-Age=3600; path=/; domain=.weibo.cn",
		"M_WEIBOCN_PARAMS=luicode%3D10000011%26lfid%3D100103type%253D17%2526q%253D%25E6%2590%259E%25E7%25AC%2591%2526t%253D0%26fid%3D2315672640113513%26uicode%3D10000011; expires=Fri, 23-Apr-2021 13:42:01 GMT; Max-Age=600; path=/; domain=.weibo.cn; HttpOnly",
	}
	CookieStr2Map(originalCookieStr)
	SetCookie(newCookies)
}

func FetchVideos()  {
	videoContainerId, err := QueryMainPage()
	if err != nil {
		fmt.Println(err)
		return
	}
	GetUserVideosAndComments(videoContainerId)
}

//访问主页
func QueryMainPage() (string, error) {
	url := "https://m.weibo.cn/api/container/getIndex?type=uid&value=2640113513&containerid=1005052640113513"
	urlParams := make(map[string]string)
	urlParams["type"] = "uid"
	urlParams["value"] = uid
	urlParams["containerid"] = containerid

	headers := make(map[string]string)
	headers["accept"] = "application/json, text/plain, */*"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["mweibo-pwa"] = "1"
	refererUrl := "https://m.weibo.cn/u/" + uid
	headers["referer"] = refererUrl
	headers["x-requested-with"] = "XMLHttpRequest"
	headers["x-xsrf-token"] = cookieMap["xsrf-token"]

	resp, err := GetResponse(url, urlParams, headers)
	if err != nil {
		return "", errors.New("QueryMainPage : " + err.Error())
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	mainPage := MainPage{}
	err = json.Unmarshal(bs, &mainPage)
	if err != nil {
		ShowResponse(resp)
		panic(err)
	}
	videoContainerId := ""
	for _, tab := range mainPage.Data.TabsInfo.Tabs {
		if tab.TabType == "video" {
			videoContainerId = tab.Containerid
			break
		}
	}
	return videoContainerId, nil
}

func GetUserVideosAndComments(videoContainerId string) {
	since_id := ""
	/*
	get first page index
	*/
	urlParams := make(map[string]string)
	urlParams["uid"] = uid
	urlParams["t"] = t
	urlParams["luicode"] = getMWeiboCnParams("luicode")
	urlParams["lfid"] = getMWeiboCnParams("lfid")
	urlParams["type"] = "uid"
	urlParams["value"] = uid
	urlParams["containerid"] = videoContainerId


	firstPageIndex, err := getOnePageIndex(urlParams)
	if err != nil {
		fmt.Println("firstPageIndex error : " + err.Error())
		return
	}
	if firstPageIndex.OK != 1 {
		fmt.Println("GetUserVideos error : first page response.ok is not 1")
		fmt.Printf("%#v", firstPageIndex)
		return
	}

	videoTotal := firstPageIndex.Data.CardlistInfo.Total
	getTotal := 0

	time.Sleep(5 * time.Second)

	for _, card := range firstPageIndex.Data.Cards {
		getAVideoAndComments(&card)
		getTotal += 1
		fmt.Println("get video finish : " + strconv.FormatInt(int64(getTotal), 10) + "/" + strconv.FormatInt(int64(videoTotal), 10))
	}

	time.Sleep(10 * time.Second)


	since_id = strconv.FormatInt(firstPageIndex.Data.CardlistInfo.Since_id, 10)

	for getTotal < videoTotal  {
		urlParams := make(map[string]string)
		urlParams["uid"] = uid
		urlParams["t"] = t
		urlParams["luicode"] = getMWeiboCnParams("luicode")
		urlParams["lfid"] = getMWeiboCnParams("lfid")
		urlParams["type"] = "uid"
		urlParams["value"] = uid
		urlParams["containerid"] = videoContainerId
		urlParams["since_id"] = since_id

		onePageIndex, err := getOnePageIndex(urlParams)
		if err != nil {
			fmt.Println("get onePageIndex error" + err.Error())
			break
		}


		since_id = strconv.FormatInt(onePageIndex.Data.CardlistInfo.Since_id, 10)

		for _, card := range onePageIndex.Data.Cards {
			getAVideoAndComments(&card)
			getTotal += 1
			fmt.Println("get video finish : " + strconv.FormatInt(int64(getTotal), 10) + "/" + strconv.FormatInt(int64(videoTotal), 10))
		}

		if onePageIndex == nil || onePageIndex.Data.CardlistInfo.Since_id == 0 {
			fmt.Println("sinceId == 0, break")
			break
		}

		time.Sleep(time.Second * 10)
	}

	fmt.Printf("get %d video, over", getTotal)
}

func getOnePageIndex(urlParams map[string]string)  (*OnePageIndex, error){
	headers := make(map[string]string)
	headers["accept"] = "application/json, text/plain, */*"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["mweibo-pwa"] = "1"
	refererUrl := "https://m.weibo.cn/u/" + urlParams["uid"]
	headers["referer"] = refererUrl
	headers["x-requested-with"] = "XMLHttpRequest"
	headers["x-xsrf-token"] = cookieMap["xsrf-token"]

	resp, err := GetResponse("https://m.weibo.cn/api/container/getIndex", urlParams, headers)
	if err != nil {
		return nil, errors.New("getOnePageIndex error : " + err.Error())
	}

	/*
	//gzip解压
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	bs, err := ioutil.ReadAll(gr)
	 */
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	//解析json
	onePageIndex := OnePageIndex{}

	err = json.Unmarshal(bs, &onePageIndex)
	if err != nil {
		panic(err)
	}

	if onePageIndex.OK != 1 {
		ShowResponse(resp)
		return nil, errors.New("getOnePageIndex error : response json.ok is not 1")
	}

	return &onePageIndex, nil
}

var VideoExisted = errors.New("video is existed")
var NotAVideo = errors.New("not a video")
var VideoUrlEmpty = errors.New("video url is empty")


func getAVideoAndComments(card *Card)  {
	fmt.Println("start get " + card.Mblog.Id)
	if card.Mblog.PageInfo.Type != "video" {
		fmt.Println("not a video")
		return
	}
	//video是否已经存在
	video, err := db.GetVideoById(card.Mblog.Id)
	if err == nil && video != nil {
		fmt.Println("video is existed")
		return
	}

	var videoUrl string
	if card.Mblog.PageInfo.MediaInfo.StreamUrl != "" {
		videoUrl = card.Mblog.PageInfo.MediaInfo.StreamUrl
	}else if card.Mblog.PageInfo.MediaInfo.StreamUrlHd != "" {
		videoUrl = card.Mblog.PageInfo.MediaInfo.StreamUrlHd
	}else {
		fmt.Println("video url is empty")
		return
	}

	videoName := GetFileNameFromUrl(videoUrl)
	savePath := config.VIDEO_DIR + videoName
	err = downLoadOneVideoAndSave(videoUrl, savePath)
	if err != nil {
		fmt.Println("fetchVideoAndSave error : " + err.Error())
		return
	}else {
		fmt.Println("get video " + videoName + " sucess")
	}

	time.Sleep(5 * time.Second)

	/*
		获取评论
	*/
	wbId := card.Mblog.Id
	fmt.Println("start get " + wbId + " comments")
	cNum, err := GetVideoComments(wbId)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("get comments sucess %d / %d \n", cNum, card.Mblog.CommentsCount)
	}

	//添加video
	newVideo := db.Video{
		ID:            card.Mblog.Id,
		CreateTime:   card.Mblog.CreatedAt,
		VideoTitle:    card.Mblog.PageInfo.Title,
		VideoFileName: videoName,
		VideoSeconds:  card.Mblog.PageInfo.MediaInfo.Duration,
		LikeNum:       card.Mblog.AttitudesCount,
		CommentNum:    card.Mblog.CommentsCount,
	}
	err = db.AddVideo(&newVideo)
	if err != nil {
		fmt.Println("add video error : " + err.Error())
		return
	}
	fmt.Println("get " + card.Mblog.Id + " sucess")

	time.Sleep(10 * time.Second)
}


func downLoadOneVideoAndSave(videoUrl, savePath string) error{
	headers := make(map[string]string)
	headers["Accept"] = "*/*"
	headers["Accept-Encoding"] = "identity;q=1, *;q=0"
	headers["Accept-Language"] = "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"
	headers["Connection"] = "keep-alive"
	headers["Host"] = "f.video.weibocdn.com"
	headers["Range"] = "bytes=0-"
	headers["Referer"] = "https://m.weibo.cn/"

	resp,err := GetResponse(videoUrl, nil, headers)
	if err != nil {
		return errors.New("downloadOneVideoAndSave error : " + err.Error())
	}

	if resp.Header["Content-Type"][0] != "video/mp4" {
		ShowResponse(resp)
		return errors.New("downLoad video error : is not a mp4")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	n, err := f.Write(body)
	if err != nil {
		panic(err)
	}

	if n == 0 {
		return errors.New("fetchVideoAndSave write 0 bytes")
	}

	return nil
}

