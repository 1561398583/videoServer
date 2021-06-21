package db

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primarykey"`

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
				return nil
			}
		}

		//未知错误
		panic(result.Error)
	}
	return nil
}

func GetUserById(id string)  (*User,error){
	user := User{}
	tx := DB.First(&user, "id = ?", id)
	if tx.Error != nil {
		//fmt.Printf("%#v", tx.Error)	//查看error类型
		//用户不存在
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, NotFindErr
		}

		//未知错误
		panic(tx.Error)
	}
	return &user, nil
}


func UpdateUser(user *User)  error{
	tx := DB.Save(user)
	if tx.Error != nil {
		//fmt.Printf("%#v", tx.Error)	//查看error类型
		//未知错误
		panic(tx.Error)
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


