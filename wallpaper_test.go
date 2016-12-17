package wallpaper

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func SetFile(path string) {
	absolute, err := filepath.Abs(path)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = SetWallpaper(absolute)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestSet(t *testing.T) {
	SetFile("test1.jpg")
	time.Sleep(1000 * time.Millisecond)
	SetFile("test2.jpg")
}
