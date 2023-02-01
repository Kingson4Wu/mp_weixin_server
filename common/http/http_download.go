package http

import (
	"github.com/kingson4wu/mp_weixin_server/common/file"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string, destDir string, fileName string) {

	if !file.Exists(destDir) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			log.Println("http.Download MkdirAll error:" + err.Error())
			panic(err)
		}
		log.Println("http.Download MkdirAll success: " + destDir)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("http.Download error:" + err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join(destDir, fileName))
	if err != nil {
		log.Println("http.Download Create error:" + err.Error())
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Download Copy error:" + err.Error())
		panic(err)
	}
}
