package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//从url抓取文件并存储
func FetchFileAndSave(fileUrl, savePath string) error{
	fmt.Println("want fetch file ")
	fmt.Println("From : " + fileUrl)
	fmt.Println("To : " + savePath)

	client := http.DefaultClient

	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		fmt.Println("Fail : " + err.Error())
		return err
	}
	//照着谷歌浏览器f12中的信息写,冒充谷歌浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	//client执行这个request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail : " + err.Error())
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("response status is " + resp.Status)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

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

	n, err := f.Write(body)
	if err != nil {
		return err
	}

	fmt.Printf("write %d bytes to :", n)

	return nil
}


