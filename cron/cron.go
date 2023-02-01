package job

import (
	"github.com/kingson4wu/mp_weixin_server/admin"
	file2 "github.com/kingson4wu/mp_weixin_server/common/file"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxmail"
	"log"
	"strconv"
	"time"

	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"github.com/kingson4wu/mp_weixin_server/mail"
	"github.com/robfig/cron/v3"
)

func InitCron() {

	c := cron.New()
	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("photos notify ... ")

		body := ""
		for _, account := range admin.Accounts() {
			photoList := gorm.SelectPhotos(account, time.Now().AddDate(0, 0, -1))
			for _, photo := range photoList {
				//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
				body = body + "<img src='" + photo + "'/><br/>"
			}
		}
		if body == "" {
			return
		}

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

		for _, account := range admin.Accounts() {
			wxmail.SendMail(account, "时光机", "来了！<br/>"+body, attachments)
		}

	})
	// 开始执行（每个任务会在自己的 goroutine 中执行）

	c.AddFunc("TZ=Asia/Shanghai 0 10 * * *", func() {

		log.Println("daily task notify ... ")

		for _, account := range admin.Accounts() {
			todoList := gorm.SelectTodoList(account)
			todoTaskNotify(account, &todoList, "")
		}

		/*tail := "<br/><br/><br/><br/><br/><br/><br/>" +
		"今天的我比昨天更优秀吗?<br/>" +
		"来次施德明吧<br/>"*/

		//TODO共同的任务，新增admin字段

	})

	c.Start()

	log.Println("cron task start ... ")

}

func todoTaskNotify(account string, todoList *[]gorm.TodoItem, tail string) {
	var content string
	if len(*todoList) > 0 {

		body := ""
		for i, item := range *todoList {
			body = body + strconv.Itoa(i) + "、" + item.Content + "<br/>"
		}

		content = body + tail
	}
	if content != "" {
		wxmail.SendMail(account, "每日任务", content, []mail.Attachment{})
	}

}
