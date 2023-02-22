package minio

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	CreateBucket("tiktoktest")
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile("tiktoktest", "test.mp4", "./test.mp4", "video/mp4")
	fmt.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()
	fi, _ := os.Stat("./test.mp4")
	err := UploadFile("tiktoktest", "ceshi2", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := GetFileUrl("tiktoktest", "test.mp4", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
