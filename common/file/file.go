package file

import (
	"os"
)

func Exists(path string) bool {

	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return false
}

func GetAllFile(pathname string) ([]string, error) {
	rd, err := os.ReadDir(pathname)
	var filePaths []string
	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFile(pathname + fi.Name())
		} else {
			filePaths = append(filePaths, pathname+string(os.PathSeparator)+fi.Name())
		}
	}
	return filePaths, err
}
