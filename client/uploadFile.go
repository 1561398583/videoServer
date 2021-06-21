package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"yx.com/videos/creeper"
)

func UploadFile(url, localFilePath, serverFilePath string)  error{
	file, err := os.Open(localFilePath)
	if err != nil {
		return errors.New("open " + localFilePath + " fail : " + err.Error())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("close file err : " + err.Error())
		}
	}()

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
		return errors.New("request error : " + err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("close resp.body err : " + err.Error())
		}
	}()

	if resp.StatusCode != 200 {
		creeper.ShowResponse(resp)
		fmt.Println()
		return errors.New("status is not 200")
	}

	return nil
}

func IsFileExist(url, filePath string) (bool, error) {
	fp := fp{
		FilePath: filePath,
	}
	data, err := json.Marshal(fp)
	if err != nil {
		return false, errors.New("IsFileExist json.Marshal error : "+err.Error())
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return false, errors.New("IsFileExist http.NewRequest error : "+err.Error())
	}
	//post数据并接收http响应
	resp,err :=http.DefaultClient.Do(request)
	if err!=nil{
		return false, errors.New("IsFileExist request error : "+err.Error())
	}else {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				panic(err)
			}
		}()

		respBody,err :=ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, errors.New("IsFileExist read resp.body error : "+err.Error())
		}

		result := RespJ{}
		err = json.Unmarshal(respBody, &result)
		if err != nil {
			return false, errors.New("IsFileExist json.Unmarshal error : "+err.Error())
		}

		if result.Status != "ok" {
			return false, errors.New("IsFileExist error : "+result.ErrorMsg)
		}

		if result.Data == "true" {
			return true, nil
		}else {
			return false, nil
		}
	}
}

type fp struct {
	FilePath string
}

type RespJ struct {
	Status string //ok; error
	ErrorMsg string
	Data string
}
