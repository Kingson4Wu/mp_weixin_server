package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"github.com/kingson4wu/weixin-app/config"
	"github.com/kingson4wu/weixin-app/gorm"
	"github.com/kingson4wu/weixin-app/mail"
	"github.com/kingson4wu/weixin-app/service"
	"github.com/kingson4wu/weixin-app/timingwheel"

	"github.com/kingson4wu/weixin-app/common"
	"github.com/kingson4wu/weixin-app/weixin"
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

//Go语言标准库之log : https://www.cnblogs.com/nickchen121/p/11517450.html TODO

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

			receviceMsg := weixin.WXMsgReceive(context)

			weixin.HandleMsg(receviceMsg, context)

		} else {
			context.String(http.StatusOK, "")
			log.Println("replyText failure")
		}

	})

	initConfig()
	initLogger()
	initWeixinAccessToken()
	initExtranetIpCheck()
	initTimer()

	//gorm.AddExtranetIp("120.230.98.231")
	//gorm.AddPhoto("http://mmbiz.qpic.cn/mmbiz_jpg/jRPicmoSEZ5UvQshXWvAZuzSn6Kl4ySXlISdL6iacaKSicxtDdS3lCWUMj78mlu8qKiam7F1m1yRL3mzpRNYaXUX5Q/0", "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ")

	log.Println("started finished ...")
	r.Run(":8989")

}

func initTimer() {
	//初始化一个tick是1s，wheelSize是32的时间轮：
	tw := timingwheel.NewTimingWheel(time.Second, 32)
	tw.Start()
	// 添加任务
	//通过AfterFunc方法添加一个15s的定时任务，如果到期了，那么执行传入的函数。
	tw.AfterFunc(time.Second*15, func() {
		fmt.Println("The timer fires")

	})

}

func initExtranetIpCheck() {
	go extranetIpCheck()
}

func extranetIpCheck() {
	extranetIp := service.GetExtranetIp()
	log.Println("extranetIp: " + extranetIp)
	//存数据库， 定时监测

	ticker := time.NewTicker(3600 * time.Second)
	for {
		<-ticker.C
		extranetIp := service.GetExtranetIp()
		log.Println("extranetIp: " + extranetIp)

		if !gorm.ExistExtranetIp(extranetIp) {
			log.Println("extranetIp not exist notify ...")
			mail.SendMail([]string{"819966354@qq.com"}, "白名单不存在，请配置", extranetIp)
		}

	}
	//TODO 主动查外网ip的接口
}

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

func initWeixinAccessToken() {
	service.GetAccessToken()
}

func initLogger() {
	// 创建、追加、读写，777，所有权限
	logPath := common.CurrentUserDir() + "/.weixin_app/work/log.log"

	if !common.Exists(logPath) {
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
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

//Gin还有很多功能，比如路由分组，自定义中间件，自动Crash处理等

func initConfig() {
	//yamlFile, err := ioutil.ReadFile("./config/config.yml")

	configPath := common.CurrentUserDir() + "/.weixin_app/config/config.yml"

	exist, err := common.PathExists(configPath)
	if err != nil {
		panic(err)
	}
	if !exist {
		log.Println(configPath + " is not exist")
	}

	yamlFile, err := ioutil.ReadFile(configPath)
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
