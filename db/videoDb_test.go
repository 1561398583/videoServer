package db

import (
	"fmt"
	"testing"
)

func TestGetVideosBySinceId(t *testing.T) {
	sinceId := "0"
	videos, err := GetVideosBySinceId(sinceId, 10)
	if err != nil {
		t.Error(err)
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetFirstVideo(t *testing.T) {
	video := GetFirstVideo()
	fmt.Printf("%#v", *video)
}

func TestGetLastVideoVideo(t *testing.T) {
	video := GetLastVideo()
	fmt.Printf("%#v", *video)
}

func TestGetPreVideo(t *testing.T) {
	video := GetPreVideo("4629254341920300")
	fmt.Printf("%#v", *video)
}
