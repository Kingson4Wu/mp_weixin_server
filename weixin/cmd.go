package weixin

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/weixin-app/admin"
	"github.com/kingson4wu/weixin-app/common"
	"github.com/kingson4wu/weixin-app/config"
	"github.com/kingson4wu/weixin-app/gorm"
	"github.com/kingson4wu/weixin-app/mail"
	"github.com/kingson4wu/weixin-app/service"
)

//https://studygolang.com/articles/2212

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
	PicUrl       string
	MediaId      string
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func value2CDATA(v string) CDATAText {
	return CDATAText{v}
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

//https://juejin.cn/post/6844904114707496973
// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) *WXTextMsg {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return nil
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
	return &textMsg

}

func HandleMsg(receviceMsg *WXTextMsg, context *gin.Context) {

	msg := "【1】[添加外网ip白名单]\n" +
		"【2】[查看外网ip]\n" +
		"【3】[发送邮件]\n" +
		"【4】[查看图片地址]\n"

	if admin.IsAdminstrator(receviceMsg.FromUserName) {
		if strings.HasPrefix(receviceMsg.Content, "[添加外网ip白名单]") {
			extranetIp := strings.Replace(receviceMsg.Content, "[添加外网ip白名单]", "", 1)
			log.Println("add extranetIp to white list : " + extranetIp)
			gorm.AddExtranetIp(extranetIp)

			msg = "添加成功"
		}

		if strings.HasPrefix(receviceMsg.Content, "[查看外网ip]") {
			extranetIp := service.GetExtranetIp()
			msg = extranetIp
		}

		if strings.HasPrefix(receviceMsg.Content, "[发送邮件]") {

			dateTime := time.Now().AddDate(0, 0, -1)

			photoList := gorm.SelectPhotos(receviceMsg.FromUserName, dateTime)

			if len(photoList) > 0 {

				body := ""
				for _, photo := range photoList {
					//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
					body = body + "<img src='" + photo + "'/><br/>"
				}

				//log.Println(body)

				//dateTime := time.Now()

				storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + dateTime.Format("2006_01_02")

				filePaths, err := common.GetAllFile(storeDirPath)
				if err != nil {
					panic(err)
				}
				attachments := make([]mail.MailAttachment, len(filePaths))
				for i, filePath := range filePaths {
					attachments[i] = mail.MailAttachment{FilePath: filePath, Name: ""}
				}

				SendMail(receviceMsg.FromUserName, "时光机", "来了！<br/>"+body, attachments)
				msg = "发送成功"
			} else {
				msg = "没有图片"
			}

		}

		if strings.HasPrefix(receviceMsg.Content, "[查看图片地址]") {

			photoList := gorm.SelectTodayPhotos(receviceMsg.FromUserName)

			if len(photoList) > 0 {

				body := ""
				for _, photo := range photoList {
					body = body + photo + "\n"
				}

				//log.Println(body)
				msg = body
			} else {
				msg = "没有图片"
			}

		}

		//fmt.Println("receviceMsg.MsgType:" + receviceMsg.MsgType)
		if receviceMsg.MsgType == "image" {
			log.Println("receviceMsg.PicUrl:" + receviceMsg.PicUrl)
			gorm.AddPhoto(receviceMsg.PicUrl, receviceMsg.FromUserName)

			currentTime := time.Now()

			storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + currentTime.Format("2006_01_02")

			fileName := currentTime.Format("2006_01_02_15_04_05_000000") + ".jpg"

			log.Println("image storeDirPath :" + storeDirPath)
			log.Println("image fileName :" + fileName)

			common.Download(receviceMsg.PicUrl, storeDirPath, fileName)

			msg = "保存成功"
		}

		if receviceMsg.MsgType == "video" {
			//56_cB3HghFQd5R-Vh44zlwfpHkGSmO7E1YvK-V188XkxcfIRogYmG_qx6nFrmrS--EkiGqqd_E1RFoAc_jGOHZl992Yd1kHmNoE2DYcbrNxVHYw9VScrgCwtahtIPSDTMm_4PiFmETbJapAED5BKHYjAHAAXO
			//curl -I -G "https://api.weixin.qq.com/cgi-bin/media/get?access_token=56_cB3HghFQd5R-Vh44zlwfpHkGSmO7E1YvK-V188XkxcfIRogYmG_qx6nFrmrS--EkiGqqd_E1RFoAc_jGOHZl992Yd1kHmNoE2DYcbrNxVHYw9VScrgCwtahtIPSDTMm_4PiFmETbJapAED5BKHYjAHAAXO&media_id=StR6zYhR2A5JRAmcK-Sc4-jsJp3nJ31VLN41xhu9HZ0TD1gp_pSiQJdhQaOH0951IV-BWF-ATa8ahrYj-PR8Jg"

			log.Println("receviceMsg.Video medialId:" + receviceMsg.MediaId)

			currentTime := time.Now()

			storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_video" + "/" + currentTime.Format("2006_01_02")

			fileName := currentTime.Format("2006_01_02_15_04_05_000000") + ".mp4"

			log.Println("video storeDirPath :" + storeDirPath)
			log.Println("video fileName :" + fileName)

			accessToken := service.GetAccessToken()
			url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", accessToken, receviceMsg.MediaId)

			common.Download(url, storeDirPath, fileName)

			msg = "保存成功"
		}
	}

	context.Header("Content-Type", "text/xml; charset=utf-8")

	text, _ := makeTextResponseBody(receviceMsg.ToUserName, receviceMsg.FromUserName, msg)
	replyText := string(text)

	log.Println(replyText)

	context.String(http.StatusOK, replyText)

	log.Println("replyText success")

	//sendMail(receviceMsg.FromUserName, "通知", "Hello World !")

}

func SendMail(account string, subject string, body string, attachements []mail.MailAttachment) {

	mailConfig := config.GetMailConfig()
	elements := mailConfig.UserMailInfos
	elementMap := make(map[string]string)
	for _, data := range elements {
		elementMap[data.OpenId] = data.Address
	}

	if v, ok := elementMap[account]; ok {

		///--------

		// 邮件接收方
		mailTo := []string{
			//可以是多个接收人
			//"xxx@163.com",
			v,
		}

		if len(attachements) > 0 {
			err := mail.SendMailWithAttachment(mailTo, subject, body, attachements)
			if err != nil {
				fmt.Println("Send fail! - ", err)
				return
			}
		} else {
			err := mail.SendMail(mailTo, subject, body)
			if err != nil {
				fmt.Println("Send fail! - ", err)
				return
			}
		}

	}
}

/**

<xml>
    <ToUserName>
        <![CDATA[gh_66ad12244999]]>
    </ToUserName>
    <FromUserName>
        <![CDATA[oqV-XjlEcZZcA4pCwoaiLtnFF0XQ]]>
    </FromUserName>
    <CreateTime>1648481096</CreateTime>
    <MsgType>
        <![CDATA[image]]>
    </MsgType>
    <PicUrl>
        <![CDATA[http://mmbiz.qpic.cn/mmbiz_jpg/jRPicmoSEZ5UvQshXWvAZuzSn6Kl4ySXlPpjok5eMAexqfYOD9duNx4aUIHyg9QXAPAu3RU1xriamEwI6e4l9wPw/0]]>
    </PicUrl>
    <MsgId>23600573381408707</MsgId>
    <MediaId>
        <![CDATA[dYcHDmyFCVADtJ798uvkQ-Bl4tVBMpjDWBXr4YvW_qxWHRnA7TxXxgHd1S6iayUX]]>
    </MediaId>
</xml>

*/
