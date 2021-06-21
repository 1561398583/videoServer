package config

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_path(t *testing.T)  {
	fmt.Println("GOOS : " + runtime.GOOS)
}