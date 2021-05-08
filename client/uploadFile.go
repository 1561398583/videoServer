package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(url, localFilePath, serverFilePath string)  {
	file, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("open " + localFilePath + " fail : " + err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// 实例化multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建multipart 文件字段
	part, err := writer.CreateFormFile("file", "file")
	if err != nil {
		panic(err)
	}

	// 写入文件数据到multipart
	_, err = part.Write(data)
	if err != nil {
		panic(err)
	}

	//将额外参数也写入到multipart
	err = writer.WriteField("dest", serverFilePath)
	if err != nil {
		panic(err)
	}

	err = writer.Close() //添加结尾的boundary
	if err != nil {
		panic(err)
	}

	//创建请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	//不要忘记加上writer.FormDataContentType()，
	//该值等于content-type :multipart/form-data; boundary=xxxxx
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("request error : " + err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp)
		return
	}

	fmt.Println("upload sucess")
}
