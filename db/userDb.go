package db

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`

	UserName string
	PassWord string
	FaceImage string
}

func CreateUser(user *User) error{
	result := DB.Create(user)
	if result.Error != nil {
		//用户已经存在
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok{	//error interface转为具体的MysqlError type
			if mysqlErr.Number == 1062 {
				return errors.New("user has existed")
			}
		}

		//未知错误
		errMsg := "db create user error : " + result.Error.Error()
		panic(errMsg)
	}
	return nil
}

func GetUserById(id string)  (*User,error){
	user := User{}
	tx := DB.First(&user, "id = ?", id)
	if tx.Error != nil {
		//用户不存在
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, tx.Error
		}

		//未知错误
		panic(tx.Error)
	}
	return &user, nil
}

func GetUserByUserName(userName string)  (*User,error){
	user := User{}
	tx := DB.First(&user, "user_name = ?", userName)
	if tx.Error != nil {
		//用户不存在
		if mysqlErr, ok := tx.Error.(*mysql.MySQLError); ok{	//error interface转为具体的MysqlError type
			if mysqlErr.Number == 1054 {
				return nil, errors.New("user not found")
			}
		}

		if tx.Error == gorm.ErrRecordNotFound {
			return nil, tx.Error
		}

		//未知错误
		errMsg := "db find user error : " + tx.Error.Error()
		panic(errMsg)
	}
	return &user, nil
}

func UpdateUser(user *User)  error{
	tx := DB.Save(user)
	if tx.Error != nil {
		//未知错误
		errMsg := "db update user error : " + tx.Error.Error()
		panic(errMsg)
	}
	return nil
}

func DeleteUserById(id string) {
	tx := DB.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		//未知错误
		errMsg := "db delete user error : " + tx.Error.Error()
		panic(errMsg)
	}
}


