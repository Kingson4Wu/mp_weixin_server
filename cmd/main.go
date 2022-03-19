package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"github.com/kingson4wu/weixin-app/config"
	"github.com/kingson4wu/weixin-app/mail"
	"github.com/kingson4wu/weixin-app/service"
)

func checkSign(signature string, timestamp string, nonce string) bool {
	//1）将token、timestamp、nonce三个参数进行字典序排序

	config := config.GetWeixinConfig()
	token := config.Token

	//将token、timestamp、nonce三个参数进行字典序排序
	var tempArray = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))

	fmt.Println("token:" + token)
	fmt.Println("timestamp:" + timestamp)
	fmt.Println("nonce:" + nonce)
	fmt.Println("nsha1once:" + sha1String)
	fmt.Println("signature:" + signature)

	//获得加密后的字符串可与signature对比
	return sha1String == signature
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

//Go语言标准库之log : https://www.cnblogs.com/nickchen121/p/11517450.html TODO

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

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"labali": "koo~",
			"fox":    "shit",
		})
	})

	r.GET("/labali_msg", func(context *gin.Context) {

		//http://127.0.0.1:8989/access_check_signature?signature=4654fdg&timestamp=3534&nonce=35fdgf
		signature := context.Query("signature")
		timestamp := context.Query("timestamp")
		nonce := context.Query("nonce")
		echostr := context.Query("echostr")

		//3）开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
		if checkSign(signature, timestamp, nonce) {
			context.String(http.StatusOK, echostr)
		} else {
			context.String(http.StatusOK, "")
		}

	})

	r.POST("/labali_msg", func(context *gin.Context) {

		//http://127.0.0.1:8989/access_check_signature?signature=4654fdg&timestamp=3534&nonce=35fdgf
		signature := context.Query("signature")
		timestamp := context.Query("timestamp")
		nonce := context.Query("nonce")
		//echostr := context.Query("echostr")
		openid := context.Query("openid")

		fmt.Println(signature)
		fmt.Println(timestamp)
		fmt.Println(nonce)
		fmt.Println(openid)

		if checkSign(signature, timestamp, nonce) {

			//body, _ := ioutil.ReadAll(context.Request.Body)
			//fmt.Println("---body/--- \r\n " + string(body))
			//go orm 保存数据库 TODO

			receviceMsg := WXMsgReceive(context)

			//fmt.Println("receviceMsg.MsgType:" + receviceMsg.MsgType)
			if receviceMsg.MsgType == "image" {
				fmt.Println("receviceMsg.PicUrl:" + receviceMsg.PicUrl)
			}

			replyText := "success"

			//if openid == "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ" {
			fmt.Println("replyText text custom ... ")
			context.Header("Content-Type", "text/xml; charset=utf-8")

			////https://studygolang.com/articles/2212

			//replyText = fmt.Sprintf("<xml><ToUserName><![%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>", "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ", "gh_66ad12244999", time.Now().Unix(), "我来了")

			//text, _ := makeTextResponseBody("gh_66ad12244999", "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ", "来了来了！！")
			text, _ := makeTextResponseBody(receviceMsg.ToUserName, receviceMsg.FromUserName, "来了来了！！")
			replyText = string(text)
			//}

			fmt.Println(replyText)
			//context.String(http.StatusOK, echostr)
			context.String(http.StatusOK, replyText)
			fmt.Println("replyText success")

			mailConfig := config.GetMailConfig()
			elements := mailConfig.UserMailInfos
			elementMap := make(map[string]string)
			for _, data := range elements {
				elementMap[data.OpenId] = data.Address
			}

			if v, ok := elementMap[receviceMsg.FromUserName]; ok {

				///--------

				// 邮件接收方
				mailTo := []string{
					//可以是多个接收人
					//"xxx@163.com",
					v,
				}

				subject := "Hello World!" // 邮件主题
				body := "测试发送邮件"          // 邮件正文

				err := mail.SendMail(mailTo, subject, body)
				if err != nil {
					fmt.Println("Send fail! - ", err)
					return
				}

			}

			fmt.Println("Send successfully!")

		} else {
			context.String(http.StatusOK, "")
			fmt.Println("replyText failure")
		}

	})

	InitConfig()

	service.GetAccessToken()

	go extranetIpCheck()

	r.Run(":8989")
}

func extranetIpCheck() {
	extranetIp := service.GetExtranetIp()
	log.Println("extranetIp: " + extranetIp)
	//存数据库， 定时监测

	ticker := time.NewTicker(3 * time.Second)
	for {
		<-ticker.C
		extranetIp := service.GetExtranetIp()
		log.Println("extranetIp: " + extranetIp)
	}
	//TODO 主动查外网ip的接口
}

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

//Gin还有很多功能，比如路由分组，自定义中间件，自动Crash处理等

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *config.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.app: %#v\n", _config.App)
	fmt.Printf("config.log: %#v\n", _config.Log)

}
