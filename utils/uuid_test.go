package utils

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		s, err := GetUUID()
		if err != nil {
			t.Errorf("getUUID error : %v\n", err)
		}
		if len(s) != 32 {
			t.Errorf("except len is 32, but %d\n", len(s))
		}
		fmt.Println(s)
	}
}
