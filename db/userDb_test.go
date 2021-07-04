package db

import (
	"fmt"
	"testing"
)

func TestDbUser(t *testing.T)  {
	//now := time.Now().Format("2006_01_02_15_04_05")
	userId  := "123"
	user := User{ID:userId, UserName: "test", FaceImage: "abc.jpg"}
	err := CreateUser(&user)
	if err != nil {
		t.Errorf("TestCreateUser error : %v \n", err)
	}
	//create 已存在的id会报错
	err = CreateUser(&user)
	if err != nil && err.Error() != "user has existed" {
		t.Errorf("TestCreateUser except error : user has excisted but : %v \n", err)
	}

	//query
	user2, err := GetUserById(userId)
	if err != nil {
		t.Errorf("query user error : %v \n", err)
	}
	if user2.UserName != "test" || user2.FaceImage != "abc.jpg" {
		t.Errorf("query user error \nexcept : %v \nbut %v \n", user, user2)
	}

	//query一个不存在的user
	userNotExist, err := GetUserById("0")
	if err == nil || err.Error() != "record not found" || userNotExist != nil{
		t.Errorf("query userNotExist expect record not found but %s \n", err.Error())
	}

	//update
	user.UserName = "has update"
	err = UpdateUser(&user)
	if err != nil {
		t.Errorf("update user error : %v \n", err)
	}
	user3, err := GetUserById(userId)
	if user3 == nil || user3.UserName != "has update" {
		t.Errorf("update user error \nexcept userName: has update \nbut %s \n", user3.UserName)
	}

	//delete
	DeleteUserById(userId)
	user4, err := GetUserById(userId)
	if err == nil || user4 != nil {
		t.Errorf("delete user fail \n")
	}
}

//看没找到会返回什么error
func TestNotFound(t *testing.T)  {
	userId := "test_notFound"
	_, err := GetUserById(userId)
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateUser(t *testing.T) {
	user := User{ID:"test1", UserName: "赵信", FaceImage: "test1.jpg"}
	err := UpdateUser(&user)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	ids := []string{
		"1006075324a",
		"1009593475",
		"1015841314",
	}
	us, err := GetUsers(ids)
	if err != nil {
		t.Error(err)
	}

	for _, u := range us{
		fmt.Println(u.UserName)
	}
}
