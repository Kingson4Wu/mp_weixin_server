package wxaction

import (
	"encoding/xml"
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/admin"
	file2 "github.com/kingson4wu/mp_weixin_server/common/file"
	http2 "github.com/kingson4wu/mp_weixin_server/common/http"
	"github.com/kingson4wu/mp_weixin_server/common/ip"
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/global"
	"github.com/kingson4wu/mp_weixin_server/weixin/accesstoken"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxmail"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"github.com/kingson4wu/mp_weixin_server/mail"
)

var weixinAccessToken *accesstoken.AccessToken

func init() {
	weixinConfig := config.GetWeixinConfig()
	weixinAccessToken = accesstoken.New(weixinConfig.Appid, weixinConfig.AppSecret)
	weixinAccessToken.Get()

}

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

// WXMsgReceive https://juejin.cn/post/6844904114707496973
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

type commandInfo struct {
	id       int
	desc     string
	usage    string
	cmd      Command
	userType global.UserType
}

type Command int

const (
	_                     = iota
	AddExtranetIp Command = iota
	QueryExtranetIp
	QueryIntranetIp
	SendMail
	QueryPhotoAddress
	AddTODO
	QueryTODO
	FinishTODO
	DeleteTODO
	AddShareTODO
	QueryShareTODO
	FinishShareTODO
	DeleteShareTODO
	NBADraft
	SetMemo
	QueryMemo
)

var (
	defaultMessage     string
	cmdIdToCommandInfo map[string]*commandInfo
	handlers           = make(map[Command]func(openid, content string) string)
)

func registerHandler(cmd Command, f func(openid, content string) string) {
	handlers[cmd] = f
}

func init() {
	cmds := []*commandInfo{
		{cmd: AddExtranetIp, desc: "添加外网ip白名单", usage: "192.168.33.174", userType: global.Admin},
		{cmd: QueryExtranetIp, desc: "查看外网ip", userType: global.Admin},
		{cmd: QueryIntranetIp, desc: "查看内网ip", userType: global.Admin},
		{cmd: SendMail, desc: "发送邮件", userType: global.USER},
		{cmd: QueryPhotoAddress, desc: "查看图片地址", userType: global.USER},
		{cmd: AddTODO, desc: "添加todo", usage: "[优先级数字，越大优先级越高]+[todo]", userType: global.USER},
		{cmd: QueryTODO, desc: "查看todo", userType: global.USER},
		{cmd: FinishTODO, desc: "完成todo", usage: "[todo id]", userType: global.USER},
		{cmd: DeleteTODO, desc: "删除todo", usage: "[todo id]", userType: global.USER},
		{cmd: AddShareTODO, desc: "添加todo[共同]", usage: "[优先级数字，越大优先级越高]+[todo]", userType: global.Admin},
		{cmd: QueryShareTODO, desc: "查看todo[共同]", userType: global.Admin},
		{cmd: FinishShareTODO, desc: "完成todo[共同]", usage: "[todo id]", userType: global.Admin},
		{cmd: DeleteShareTODO, desc: "删除todo[共同]", usage: "[todo id]", userType: global.Admin},
		{cmd: NBADraft, desc: "NBA选秀", userType: global.COMMON},
		{cmd: SetMemo, desc: "设置备忘", usage: "[memo]", userType: global.USER},
		{cmd: QueryMemo, desc: "查看备忘", userType: global.USER},
	}
	for i, cmd := range cmds {
		cmd.id = i + 1
		if cmd.usage == "" {
			cmd.usage = strconv.Itoa(cmd.id)
		} else {
			cmd.usage = strconv.Itoa(cmd.id) + "+" + cmd.usage
		}
	}

	var msg string
	for _, cmd := range cmds {
		item := fmt.Sprintf("【%v】%s\n【Usage】%s\n", cmd.id, cmd.desc, cmd.usage)
		msg += item
	}
	defaultMessage = msg

	cmdIdToCommandInfo = make(map[string]*commandInfo)
	for _, cmd := range cmds {
		cmdIdToCommandInfo[strconv.Itoa(cmd.id)] = cmd
	}

	registerHandler(AddExtranetIp, func(openid, content string) string {
		extranetIp := content
		log.Println("add extranetIp to white list : " + extranetIp)
		gorm.AddExtranetIp(extranetIp)
		return "添加成功"
	})
	registerHandler(QueryExtranetIp, func(openid, content string) string {
		return ip.GetExtranetIp()
	})
	registerHandler(QueryIntranetIp, func(openid, content string) string {
		return ip.GetIntranetIp()
	})
	registerHandler(SendMail, func(openid, content string) string {
		dateTime := time.Now().AddDate(0, 0, -1)

		photoList := gorm.SelectPhotos(openid, dateTime)

		if len(photoList) > 0 {

			body := ""
			for _, photo := range photoList {
				//body = body + "<img src='data:image/png;base64," + base64Photo + "'/><br/>"
				body = body + "<img src='" + photo + "'/><br/>"
			}

			storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + dateTime.Format("2006_01_02")

			filePaths, err := file2.GetAllFile(storeDirPath)
			if err != nil {
				panic(err)
			}
			attachments := make([]mail.Attachment, len(filePaths))
			for i, filePath := range filePaths {
				attachments[i] = mail.Attachment{FilePath: filePath, Name: ""}
			}

			wxmail.SendMail(openid, "时光机", "来了！<br/>"+body, attachments)
			return "发送成功"
		} else {
			return "没有图片"
		}
	})
	registerHandler(QueryPhotoAddress, func(openid, content string) string {
		photoList := gorm.SelectTodayPhotos(openid)

		if len(photoList) > 0 {

			body := ""
			for _, photo := range photoList {
				body = body + photo + "\n"
			}
			return body
		} else {
			return "没有图片"
		}
	})
	registerHandler(AddTODO, func(openid, content string) string {
		log.Println("add todo list : " + content)
		endIndex := strings.Index(content, "+")
		if endIndex > 0 {
			sort := content[0:endIndex]
			if v, err := strconv.Atoi(sort); err == nil {
				gorm.AddTodoItem(content[endIndex+1:], v, openid)
			}
		}
		return "添加成功"
	})
	registerHandler(QueryTODO, func(openid, content string) string {
		todoList := gorm.SelectTodoList(openid)

		if len(todoList) > 0 {

			body := ""
			for i, item := range todoList {
				body = body + strconv.Itoa(i) + "、[sort-" + strconv.Itoa(item.Sort) + "]" + "[id-" + strconv.Itoa(int(item.ID)) + "]--" + item.Content + "\n"
			}
			return body
		} else {
			return "没有todolist"
		}
	})

	registerHandler(FinishTODO, func(openid, content string) string {
		log.Println("finish todo list : " + content)
		if v, err := strconv.Atoi(content); err == nil {
			gorm.CompleteTodoItem(v)
		}
		return "完成成功"
	})

	registerHandler(DeleteTODO, func(openid, content string) string {
		log.Println("delete todo list : " + content)
		if v, err := strconv.Atoi(content); err == nil {
			gorm.DeleteTodoItem(v)
		}
		return "删除成功"
	})

	registerHandler(NBADraft, func(openid, content string) string {
		return "https://cc24-120-230-98-139.ngrok.io/"
	})

	registerHandler(AddShareTODO, func(openid, content string) string {
		log.Println("add todo list : " + content)
		endIndex := strings.Index(content, "+")
		if endIndex > 0 {
			sort := content[0:endIndex]
			if v, err := strconv.Atoi(sort); err == nil {
				gorm.AddGroupTodoItem(content[endIndex+1:], v, openid, "labali")
			}
		}
		return "添加成功"
	})
	registerHandler(QueryShareTODO, func(openid, content string) string {
		todoList := gorm.SelectGroupTodoList("labali")

		if len(todoList) > 0 {

			body := ""
			for i, item := range todoList {
				body = body + strconv.Itoa(i) + "、[sort-" + strconv.Itoa(item.Sort) + "]" + "[id-" + strconv.Itoa(int(item.ID)) + "]--" + item.Content + "\n"
			}
			return body
		} else {
			return "没有todolist"
		}
	})
	registerHandler(FinishShareTODO, func(openid, content string) string {
		log.Println("finish todo list : " + content)
		if v, err := strconv.Atoi(content); err == nil {
			gorm.CompleteGroupTodoItem(v)
		}
		return "完成成功"
	})
	registerHandler(DeleteShareTODO, func(openid, content string) string {
		log.Println("delete todo list : " + content)
		if v, err := strconv.Atoi(content); err == nil {
			gorm.DeleteGroupTodoItem(v)
		}
		return "删除成功"
	})

	registerHandler(SetMemo, func(openid, content string) string {
		log.Println("set memo : " + content)

		if content == "" {
			return "内容不能为空"
		}

		gorm.AddMemo(content, openid)
		return "设置成功"
	})

	registerHandler(QueryMemo, func(openid, content string) string {
		memo := gorm.SelectLatestMemo(openid)

		if memo != nil && memo.Content != "" {
			return memo.Content
		}
		return "没有备忘"
	})
}

func HandleMsg(receiveMsg *WXTextMsg, context *gin.Context) {

	msg := defaultMessage

	if receiveMsg.Content == "链接" {
		msg = "<![CDATA[labali天地：<a href='https://6fa8-120-235-19-241.ngrok.io/weixin_page/'>点击进入</a>]]"
	}
	log.Printf("receive content: %s\n", receiveMsg.Content)
	c := strings.SplitN(receiveMsg.Content, "+", 2)
	if len(c) >= 1 {
		cmdId := c[0]

		var content string
		if len(c) == 2 {
			content = c[1]
		}
		log.Printf("cmdId:%s, content:%s", cmdId, content)

		if cmdInfo, ok := cmdIdToCommandInfo[cmdId]; ok {
			if cmdInfo.userType == global.Admin && !admin.IsAdministrator(receiveMsg.FromUserName) {
				msg = "没有权限"
				reply(receiveMsg, context, msg)
				return
			}
			if cmdInfo.userType == global.USER {
				if _, ok := global.OpenidToMail[receiveMsg.FromUserName]; !ok {
					msg = "没有权限"
					reply(receiveMsg, context, msg)
					return
				}
			}
			if f, ok := handlers[cmdInfo.cmd]; ok {
				reply(receiveMsg, context, f(receiveMsg.FromUserName, content))
				return
			}
		}
	}

	if _, ok := global.OpenidToMail[receiveMsg.FromUserName]; ok {

		if receiveMsg.MsgType == "image" {
			log.Println("receiveMsg.PicUrl:" + receiveMsg.PicUrl)
			gorm.AddPhoto(receiveMsg.PicUrl, receiveMsg.FromUserName)

			currentTime := time.Now()

			storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_image" + "/" + currentTime.Format("2006_01_02")

			fileName := currentTime.Format("2006_01_02_15_04_05_000000") + ".jpg"

			log.Println("image storeDirPath :" + storeDirPath)
			log.Println("image fileName :" + fileName)

			http2.Download(receiveMsg.PicUrl, storeDirPath, fileName)

			//增加水印：image/draw库
			//https://www.cnblogs.com/ExMan/p/13158662.html
			//https://zhuanlan.zhihu.com/p/387753099
			//https://blog.csdn.net/diandianxiyu_geek/article/details/119382334

			msg = "保存成功"
		}

		if receiveMsg.MsgType == "video" {
			//56_cB3HghFQd5R-Vh44zlwfpHkGSmO7E1YvK-V188XkxcfIRogYmG_qx6nFrmrS--EkiGqqd_E1RFoAc_jGOHZl992Yd1kHmNoE2DYcbrNxVHYw9VScrgCwtahtIPSDTMm_4PiFmETbJapAED5BKHYjAHAAXO
			//curl -I -G "https://api.weixin.qq.com/cgi-bin/media/get?access_token=56_cB3HghFQd5R-Vh44zlwfpHkGSmO7E1YvK-V188XkxcfIRogYmG_qx6nFrmrS--EkiGqqd_E1RFoAc_jGOHZl992Yd1kHmNoE2DYcbrNxVHYw9VScrgCwtahtIPSDTMm_4PiFmETbJapAED5BKHYjAHAAXO&media_id=StR6zYhR2A5JRAmcK-Sc4-jsJp3nJ31VLN41xhu9HZ0TD1gp_pSiQJdhQaOH0951IV-BWF-ATa8ahrYj-PR8Jg"

			log.Println("receiveMsg.Video medialId:" + receiveMsg.MediaId)

			currentTime := time.Now()

			storeDirPath := file.CurrentUserDir() + "/.weixin_app/upload_video" + "/" + currentTime.Format("2006_01_02")

			fileName := currentTime.Format("2006_01_02_15_04_05_000000") + ".mp4"

			log.Println("video storeDirPath :" + storeDirPath)
			log.Println("video fileName :" + fileName)

			accessToken := weixinAccessToken.Get()
			url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", accessToken, receiveMsg.MediaId)

			http2.Download(url, storeDirPath, fileName)

			msg = "保存成功"
		}
	}

	reply(receiveMsg, context, msg)
}

func reply(receiveMsg *WXTextMsg, context *gin.Context, msg string) {
	context.Header("Content-Type", "text/xml; charset=utf-8")

	text, _ := makeTextResponseBody(receiveMsg.ToUserName, receiveMsg.FromUserName, msg)
	replyText := string(text)

	log.Println(replyText)

	context.String(http.StatusOK, replyText)

	log.Println("replyText success")
}
