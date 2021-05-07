package creeper

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
获取url中的文件名
例如：从
https://f.video.weibocdn.com/lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4?label=mp4_hd&template=540x960.24.0
截取出
lAVblciAlx07L20y92Ok01041200cf3R0E010.mp4
*/
func GetFileNameFromUrl(url string)  string{
	us1 := strings.Split(url, "?")
	if len(us1) == 0 {
		return ""
	}
	us2 := strings.Split(us1[0], "/")
	if len(us2) == 0 {
		return ""
	}
	return us2[len(us2) - 1]
}


//从url抓取文件并存储
func FetchFileAndSave(fileUrl, savePath string) error{
	client := &http.Client{}

	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		return err
	}
	//照着谷歌浏览器f12中的信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	if resp.StatusCode != 200 {
		return errors.New("response status is " + resp.Status)
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	f, err := os.Create(savePath)
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return nil
}

func UrlAddParams(url string, params map[string]string) string {
	if len(params) == 0 {
		return url
	}
	url += "?"
	for k, v := range params {
		kv := k + "=" + v
		url = url + kv + "&"
	}
	return url[:len(url) - 1]
}

func ReadFile(path string)  ([]byte, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("ReadFile open file error : " + err.Error())
	}else {
		defer func() {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}()
	}
	//读取file
	fInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fSize := fInfo.Size()
	bs := make([]byte, fSize)
	_, err = f.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs, nil
}


func ShowResponse(resp *http.Response)  {
	fmt.Println("Status Code : " + resp.Status)

	for k, h := range resp.Header {
		fmt.Print(k + " : ")
		fmt.Println(h)
	}

	fmt.Println()

	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
}

func ShowResponseHeader(resp *http.Response)  {
	fmt.Println("Status Code : " + resp.Status)

	for k, h := range resp.Header {
		fmt.Print(k + " : ")
		fmt.Println(h)
	}
}

func GetResponse(reqUrl string, urlParams map[string]string, heads map[string]string)  (*http.Response,error){
	client := http.DefaultClient

	//组装url
	reqUrl= UrlAddParams(reqUrl, urlParams)

	req,err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}

	/*
		组装head
	*/
	for k, v := range heads {
		req.Header.Set(k, v)
	}

	//照着谷歌浏览器f12中的User-Agent信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	cookie := CookieMap2Str(cookieMap)
	req.Header.Set("cookie", cookie)

	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetResponse error : " + err.Error())
	}

	//更新cookie
	if cookies, ok := resp.Header["Set-Cookie"]; ok{
		UpdateCookie(cookies)
	}

	//如果body是gzip压缩过的就解压
	if ce,ok := resp.Header["Content-Encoding"];ok {
		if ce[0] == "gzip" {
			gr, err := gzip.NewReader(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body = gr
		}
	}

	if resp.StatusCode == 200 || resp.StatusCode == 206{
		return resp, nil
	}else {
		ShowResponse(resp)
		return nil, errors.New("GetResponse response status is " + resp.Status)
	}

}

func PostResponse(client *http.Client, reqUrl string, heads map[string]string, body string)  *http.Response{

	req,err := http.NewRequest("POST", reqUrl, bytes.NewReader([]byte(body)))
	if err != nil {
		panic(err)
	}

	/*
		组装head
	*/
	for k, v := range heads {
		req.Header.Set(k, v)
	}

	//照着谷歌浏览器f12中的User-Agent信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	cookie := CookieMap2Str(cookieMap)
	req.Header.Set("cookie", cookie)

	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//更新cookie
	if cookies, ok := resp.Header["Set-Cookie"]; ok{
		UpdateCookie(cookies)
	}

	return resp
}

//key1=value1&key2=value2分解为map
func String2Map(s string) map[string]string{
	//以&为分隔符
	ss := strings.Split(s, "&")
	m := make(map[string]string)
	//每一个都是key=value
	for _, info := range ss {
		kv := strings.Split(info, "=")
		m[kv[0]] = kv[1]
	}
	return m
}







