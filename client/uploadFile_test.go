package client

import (
	"fmt"
	"io/ioutil"
	"testing"
)


func TestUploadFile1(t *testing.T) {
	UploadFile("http://121.5.72.78:8080/uploadFile", "E:\\videoProject\\server\\assets\\faceImg\\3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg","/usr/videoProject/assets/faceImg/3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg")

}

func TestUploadFile2(t *testing.T) {
	UploadFile("http://121.5.72.78:8080/uploadFile", "E:\\videoProject\\server\\assets\\videos\\000iwJWagx07HnRM3WTC010412009FhP0E010.mp4","/usr/videoProject/assets/videos/000iwJWagx07HnRM3WTC010412009FhP0E010.mp4")

}

func TestUploadFile3(t *testing.T) {
	UploadFile("http://localhost:8080/uploadFile", "E:\\videoProject\\server\\assets\\faceImg\\3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg","E:\\videoProjectA\\3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg")

}

func TestUploadFaceImg(t *testing.T) {
	fileInfos, err := ioutil.ReadDir("E:\\videoProject\\server\\assets\\faceImg")
	if err != nil {
		t.Error(err)
	}
	sucessCount := 0
	for _, fileInfo := range fileInfos {
		//如果服务器上已经有这个文件，就不用上传了
		b, err := IsFileExist("http://121.5.72.78:8080/isFileExist", "/usr/videoProject/assets/faceImg/"+fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		if b {
			fmt.Println("/usr/videoProject/assets/faceImg/"+fileInfo.Name() + " has excist")
			sucessCount ++
			continue
		}
		err = UploadFile("http://121.5.72.78:8080/uploadFile", "E:\\videoProject\\server\\assets\\faceImg\\"+fileInfo.Name(),"/usr/videoProject/assets/faceImg/"+fileInfo.Name())
		if err != nil {
			fmt.Println("upload " + fileInfo.Name() + "err : " + err.Error())
			continue
		}
		sucessCount ++
		fmt.Printf("upload sucess %d / %d \n", sucessCount, len(fileInfos))
	}
}

func TestUploadVideoPic(t *testing.T) {
	localdir := "E:\\videoProject\\server\\assets\\videoPic\\"
	removDir := "/usr/videoProject/assets/videoPic/"
	fileInfos, err := ioutil.ReadDir(localdir)
	if err != nil {
		t.Error(err)
	}
	sucessCount := 0
	for _, fileInfo := range fileInfos {
		//如果服务器上已经有这个文件，就不用上传了
		b, err := IsFileExist("http://121.5.72.78:8080/isFileExist", removDir + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		if b {
			fmt.Println(removDir + fileInfo.Name() + " has excist")
			sucessCount ++
			continue
		}
		err = UploadFile("http://121.5.72.78:8080/uploadFile", localdir + fileInfo.Name(), removDir + fileInfo.Name())
		if err != nil {
			fmt.Println("upload " + fileInfo.Name() + "err : " + err.Error())
			continue
		}
		sucessCount ++
		fmt.Printf("upload sucess %d / %d \n", sucessCount, len(fileInfos))
	}
}



func TestUploadVideos(t *testing.T) {
	localdir := "E:\\videoProject\\server\\assets\\videos\\"
	removDir := "/usr/videoProject/assets/videos/"
	fileInfos, err := ioutil.ReadDir(localdir)
	if err != nil {
		t.Error(err)
	}
	sucessCount := 0
	for _, fileInfo := range fileInfos {
		//如果服务器上已经有这个文件，就不用上传了
		b, err := IsFileExist("http://121.5.72.78:8080/isFileExist", removDir + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		if b {
			fmt.Println(removDir + fileInfo.Name() + " has excist")
			sucessCount ++
			continue
		}
		err = UploadFile("http://121.5.72.78:8080/uploadFile", localdir + fileInfo.Name(), removDir + fileInfo.Name())
		if err != nil {
			fmt.Println("upload " + fileInfo.Name() + "err : " + err.Error())
			continue
		}
		sucessCount ++
		fmt.Printf("upload sucess %d / %d \n", sucessCount, len(fileInfos))
	}
}

func TestIsFileExist(t *testing.T) {
	r, err := IsFileExist("http://localhost:8080/isFileExist", "E:\\videoProject\\server\\videoServer\\main\\main1.go")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}
