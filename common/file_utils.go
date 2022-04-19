package common

import (
	"io/ioutil"
	"os"
)

func Exists(path string) bool {

	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return false
}

func GetAllFile(pathname string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	filePaths := []string{}
	for _, fi := range rd {
		if fi.IsDir() {
			//GetAllFile(pathname + fi.Name())
		} else {
			filePaths = append(filePaths, pathname+string(os.PathSeparator)+fi.Name())
			//fmt.Println(fi.Name())
		}
	}
	return filePaths, err
}
