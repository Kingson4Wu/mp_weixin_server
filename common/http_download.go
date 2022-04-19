package common

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string, destDir string, fileName string) {

	if !Exists(destDir) {
		os.Mkdir(destDir, os.ModePerm)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(filepath.Join(destDir, fileName))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
