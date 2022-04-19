package job

import (
	"log"
	"time"

	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/weixin-app/common"
	"github.com/kingson4wu/weixin-app/gorm"
	"github.com/kingson4wu/weixin-app/mail"
	"github.com/kingson4wu/weixin-app/weixin"
	"github.com/robfig/cron/v3"
)

func CronInit() {

	c := cron.New()
	// 添加一个任务，每 30s 执行一次
	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("photos notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"
		photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))

		body := ""
		for _, photo := range photoList {
			//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
			body = body + "<img src='" + photo + "'/><br/>"
		}

		//log.Println(body)

		dateTime := time.Now().AddDate(0, 0, -1)

		storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + dateTime.Format("2006_01_02")

		filePaths, err := common.GetAllFile(storeDirPath)
		if err != nil {
			panic(err)
		}
		attachments := make([]mail.MailAttachment, len(filePaths))
		for i, filePath := range filePaths {
			attachments[i] = mail.MailAttachment{FilePath: filePath, Name: ""}
		}

		weixin.SendMail(account, "时光机", "来了！<br/>"+body, attachments)

		// fmt.Println("Every hour on the half hour")
	})
	// 开始执行（每个任务会在自己的 goroutine 中执行）

	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("daily task notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"
		photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))

		body := ""
		for _, photo := range photoList {
			//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
			body = body + "<img src='" + photo + "'/><br/>"
		}

		//log.Println(body)

		content := "1、Go算法一道；<br/>" +
			"2、Go文章一篇；<br/>" +
			"3、整理kugouNotes直到能写一篇文章；<br/>" +
			"4、学习几篇收集的技术文章(印象笔记)；<br/>" +
			"<br/><br/><br/><br/><br/><br/><br/>" +
			"今天的我比昨天更优秀吗?<br/>" +
			"来次施德明吧<br/>"

		weixin.SendMail(account, "每日任务", content, []mail.MailAttachment{})

		// fmt.Println("Every hour on the half hour")
	})

	c.Start()

	log.Println("cron task start ... ")

}
