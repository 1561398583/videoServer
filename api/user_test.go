package api

import (
	"fmt"
	"testing"
)

func TestGetOpenId(t *testing.T)  {
	code := "093PmQ000QgvTL1O0X000GK7GV2PmQ00"
	id,  err := getOpenid(code)
	if err != nil {
		t.Error("error " + err.Error())
	}
	//oEn2F4k7H8npoA2LjydjhCE7UUsY
	fmt.Println("get openId : " + id)
}
