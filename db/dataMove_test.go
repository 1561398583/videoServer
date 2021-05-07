package db

import "testing"

/**
需要把user表的int类型的id改为string类型的id
 */
func TestDataMove(t *testing.T)  {

}

type NewUser struct {
	ID        string `gorm:"primarykey"`

	UserName string
	PassWord string
	FaceImage string
}
