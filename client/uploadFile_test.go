package client

import "testing"

func TestUploadFile(t *testing.T) {
	UploadFile("http://localhost:8080/uploadFile", "E:\\videoProject\\server\\assets\\faceImg\\3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg","E:\\videoProject\\server\\3bf77dbcly8gfhtx0sfbjj20ro0rodgx.jpg")

}
