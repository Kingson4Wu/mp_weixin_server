package job

import (
	"log"
	"time"

	"github.com/kingson4wu/weixin-app/gorm"
	"github.com/kingson4wu/weixin-app/weixin"
	"github.com/robfig/cron/v3"
)

func CronInit() {

	c := cron.New()
	// 添加一个任务，每 30s 执行一次
	c.AddFunc("0 15 9 * * ? ", func() {

		log.Println("photos notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"
		photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))

		body := ""
		for _, photo := range photoList {
			//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
			body = body + "<img src='" + photo + "'/><br/>"
		}

		//log.Println(body)

		weixin.SendMail(account, "时光机", "来了！<br/>"+body)

		// fmt.Println("Every hour on the half hour")
	})
	// 开始执行（每个任务会在自己的 goroutine 中执行）

	c.AddFunc("0 15 10 * * ? ", func() {

		log.Println("daily task notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"
		photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))

		body := ""
		for _, photo := range photoList {
			//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
			body = body + "<img src='" + photo + "'/><br/>"
		}

		//log.Println(body)

		content := "1、Go算法一道；\n" +
			"2、Go文章一篇；\n" +
			"3、整理kugouNotes直到能写一篇文章；\n"

		weixin.SendMail(account, "每日任务", content)

		// fmt.Println("Every hour on the half hour")
	})

	c.Start()

	log.Println("cron task start ... ")

}
