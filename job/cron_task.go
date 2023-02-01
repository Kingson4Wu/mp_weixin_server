package job

import (
	file2 "github.com/kingson4wu/mp_weixin_server/common/file"
	"log"
	"strconv"
	"time"

	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"github.com/kingson4wu/mp_weixin_server/mail"
	"github.com/kingson4wu/mp_weixin_server/weixin"
	"github.com/robfig/cron/v3"
)

func CronInit() {

	c := cron.New()
	// 添加一个任务，每 30s 执行一次
	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("photos notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"
		account2 := "oqV-Xju6thOVtzvi0FrTWHaB5So4"
		photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))
		photoList2 := gorm.SelectPhotos(account2, time.Now().AddDate(0, 0, -1))

		if len(photoList) == 0 && len(photoList2) == 0 {
			return
		}

		body := ""
		for _, photo := range photoList {
			//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
			body = body + "<img src='" + photo + "'/><br/>"
		}

		for _, photo := range photoList2 {
			body = body + "<img src='" + photo + "'/><br/>"
		}

		//log.Println(body)

		dateTime := time.Now().AddDate(0, 0, -1)

		storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + dateTime.Format("2006_01_02")

		filePaths, err := file2.GetAllFile(storeDirPath)
		if err != nil {
			panic(err)
		}
		attachments := make([]mail.Attachment, len(filePaths))
		for i, filePath := range filePaths {
			attachments[i] = mail.Attachment{FilePath: filePath, Name: ""}
		}

		weixin.SendMail(account, "时光机", "来了！<br/>"+body, attachments)
		weixin.SendMail(account2, "时光机", "来了！<br/>"+body, attachments)

		// fmt.Println("Every hour on the half hour")
	})
	// 开始执行（每个任务会在自己的 goroutine 中执行）

	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("daily task notify ... ")

		account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"

		content := "<br/><br/><br/><br/><br/><br/><br/>" +
			"今天的我比昨天更优秀吗?<br/>" +
			"来次施德明吧<br/>"

		todoList := gorm.SelectTodoList(account)
		todoTaskNotify(account, &todoList, content)

		account2 := "oqV-Xju6thOVtzvi0FrTWHaB5So4"
		todoList2 := gorm.SelectTodoList(account2)
		if len(todoList2) > 0 {
			todoTaskNotify(account2, &todoList2, "")
		}

		//TODO共同的任务，新增admin字段

		// fmt.Println("Every hour on the half hour")
	})

	c.Start()

	log.Println("cron task start ... ")

}

func todoTaskNotify(account string, todoList *[]gorm.TodoItem, content string) {
	if len(*todoList) > 0 {

		body := ""
		for i, item := range *todoList {
			body = body + strconv.Itoa(i) + "、" + item.Content + "<br/>"
		}

		//log.Println(body)
		content = body + content
	}
	weixin.SendMail(account, "每日任务", content, []mail.Attachment{})
}
