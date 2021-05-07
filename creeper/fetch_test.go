package creeper

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)




func TestJson(t *testing.T)  {
	rb := WBRequestPostBody1{}
	rb.Data.ComponentPlayPlayinfo.Oid = "1034:4618086914129924"
	r := obj2json(&rb)
	fmt.Println(string(r))
}

type WBRequestPostBody1 struct {
	Data struct{
		ComponentPlayPlayinfo struct{
			Oid string	`json:"oid"`
		}	`json:"Component_Play_Playinfo"`
	} `json:"data"`
}

func obj2json(obj interface{}) string {
	r, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("json Marshal error " + err.Error())
		return ""
	}
	return string(r)
}

func TestFetchAUrl(t *testing.T)  {
	url := "https://m.weibo.cn/comments/hotflow"

	//cookie := "SINAGLOBAL=4045166587677.128.1617093181167; ULV=1617274969050:4:3:4:7077268667971.459.1617274969011:1617267845307; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9Whis7Bljw-wJdnwKoN5a8BK5NHD95QNSKn01h.feh5pWs4DqcjiMND0IoqE1Kzc; UOR=,,graph.qq.com; webim_unReadCount=%7B%22time%22%3A1617275246965%2C%22dm_pub_total%22%3A1%2C%22chat_group_client%22%3A0%2C%22chat_group_notice%22%3A0%2C%22allcountNum%22%3A1%2C%22msgbox%22%3A0%7D; SUB=_2A25NYdk5DeThGeFL41EZ-SvMyz2IHXVuredxrDV8PUJbkNANLVHNkW1NfY7MjXyWj3JX8SvGflIdWlD41qp3LWF5; _s_tentry=login.sina.com.cn; Apache=7077268667971.459.1617274969011; login_sid_t=3205c91bfed25a26a68658e96c9edeed; cross_origin_proto=SSL; crossidccode=CODE-yf-1LrVas-24GTq9-yUxqIxoCo7j5NKW0db030; appkey=; WBStorage=8daec78e6a891122|undefined; WBtopGlobal_register_version=2021040119; SSOLoginState=1617275241; wvr=6"

	client := &http.Client{}

	urlParams := make(map[string]string)
	urlParams["id"] = "4619851467064883"
	urlParams["mid"] = "4619851467064883"
	urlParams["max_id"] = "138163372234924"
	urlParams["max_id_type"] = "0"

	url = UrlAddParams(url, urlParams)


	req,err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//在浏览器上先登录，获取cookie信息，再照着浏览器request中的header写
	//照着谷歌浏览器f12中的User-Agent信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Set("referer", "https://m.weibo.cn/detail/4619851467064883")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-XSRF-TOKEN", "3533ee")


	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	gbk := simplifiedchinese.GBK
	ec := gbk.NewDecoder()
	r, err := ec.Bytes(bs)
	if err != nil {
		fmt.Println("decode error " + err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("ERROR response status is " + resp.Status)
	}

	fmt.Println(string(r))
}

func Test1(t *testing.T)  {
	FetchVideos()
}

func Test2(t *testing.T)  {
	fmt.Println("\u6682\u65e0\u6570\u636e")
}

func TestGetAVideoAndComments(t *testing.T)  {
	InitCookie()
	/*
		get first page index
	*/
	urlParams := make(map[string]string)
	urlParams["uid"] = "2640113513"
	urlParams["t"] = "0"
	urlParams["luicode"] = getMWeiboCnParams("luicode")
	urlParams["lfid"] = getMWeiboCnParams("lfid")
	urlParams["type"] = "uid"
	urlParams["value"] = uid
	urlParams["containerid"] = getMWeiboCnParams("fid")


	firstPageIndex, err := getOnePageIndex(urlParams)
	if err != nil {
		fmt.Println(err)
	}
	card := firstPageIndex.Data.Cards[3]
	getAVideoAndComments(&card)

}

func TestGetIndex(t *testing.T)  {
	InitCookie()

	since_id := ""
	/*
		get first page index
	*/
	urlParams := make(map[string]string)
	urlParams["uid"] = uid
	urlParams["t"] = "0"
	urlParams["luicode"] = getMWeiboCnParams("luicode")
	urlParams["lfid"] = getMWeiboCnParams("lfid")
	urlParams["type"] = "uid"
	urlParams["value"] = uid
	urlParams["containerid"] = getMWeiboCnParams("fid")


	firstPageIndex, err := getOnePageIndex(urlParams)
	if err != nil {
		fmt.Println("GetUserVideos error : " + err.Error())
	}
	if firstPageIndex.OK != 1 {
		fmt.Println("GetUserVideos error : first page response.ok is not 1")
		fmt.Printf("%#v", firstPageIndex)
	}


	since_id = strconv.FormatInt(firstPageIndex.Data.CardlistInfo.Since_id, 10)

	for   {
		urlParams := make(map[string]string)
		urlParams["uid"] = uid
		urlParams["t"] = "0"
		urlParams["luicode"] = getMWeiboCnParams("luicode")
		urlParams["lfid"] = getMWeiboCnParams("lfid")
		urlParams["type"] = "uid"
		urlParams["value"] = uid
		urlParams["containerid"] = getMWeiboCnParams("fid")
		urlParams["since_id"] = since_id

		onePageIndex, err := getOnePageIndex(urlParams)
		if err != nil {
			fmt.Println("get onePageIndex error")
			break
		}
		if onePageIndex == nil || onePageIndex.Data.CardlistInfo.Since_id == 0 {
			fmt.Println("sinceId == 0, break")
			break
		}

	}
}