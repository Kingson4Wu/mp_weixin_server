package common

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string, destDir string, fileName string) {

	if !Exists(destDir) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			log.Println("Download MkdirAll error:" + err.Error())
			panic(err)
		}
		log.Println("Download MkdirAll success: " + destDir)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Download error:" + err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(filepath.Join(destDir, fileName))
	if err != nil {
		log.Println("Download Create error:" + err.Error())
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Download Copy error:" + err.Error())
		panic(err)
	}
}
