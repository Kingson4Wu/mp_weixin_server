package logger

import (
	"github.com/kingson4wu/go-common-lib/file"
	"log"
	"os"
)

//Go语言标准库之log : https://www.cnblogs.com/nickchen121/p/11517450.html TODO

func InitLogger() {
	// 创建、追加、读写，777，所有权限
	logPath := file.CurrentUserDir() + "/.weixin_app/work/log.log"

	if !file.Exists(logPath) {
		log.Println("create log file ... ")
		_, err := os.Create(logPath)
		if err != nil {
			panic(err)
		}
		//f.Sync()
		//f.Close()
	}

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[labali]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ltime)
}
