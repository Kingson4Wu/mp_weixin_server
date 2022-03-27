package common

import (
	"log"
	"os"
	"os/user"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CurrentUserDir() string {

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return u.HomeDir

	//fmt.Println("Home dir:", u.HomeDir)
}

func AppDataDir() string {
	return CurrentUserDir() + "/.weixin_app"
}
