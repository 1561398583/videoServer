package creeper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchPage(url string) (*goquery.Document, error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,err
	}
	//照着谷歌浏览器f12中的User-Agent信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")
	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	if resp.StatusCode != 200 {
		return nil,errors.New("response status is " + resp.Status)
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil,err
	}

	return doc, nil
}




func fetchVideoAndSave(videoUrl, savePath string) error{
	client := &http.Client{}
	req, err := http.NewRequest("GET", videoUrl, nil)
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
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("fetchVideoAndSave : response status is " + resp.Status)
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fPath := savePath
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(body)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("fetchVideoAndSave write 0 bytes")
	}

	return nil
}

