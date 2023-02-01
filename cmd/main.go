package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/bash"
	"github.com/kingson4wu/mp_weixin_server/common/ip"
	"github.com/kingson4wu/mp_weixin_server/common/proc"
	"github.com/kingson4wu/mp_weixin_server/global"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxaction"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxmail"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/cron"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"github.com/kingson4wu/mp_weixin_server/mail"
	"github.com/kingson4wu/mp_weixin_server/timingwheel"

	"github.com/fvbock/endless"
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

			receviceMsg := wxaction.WXMsgReceive(context)

			wxaction.HandleMsg(receviceMsg, context)

		} else {
			context.String(http.StatusOK, "")
			log.Println("replyText failure")
		}

	})

	initConfig()
	initLogger()
	gorm.InitDB()
	job.CronInit()

	//gorm.SelectPhotos("oqV-XjlEcZZcA4pCwoaiLtnFF0XQ")

	//initExtranetIpCheck()
	initTimer()

	/** 系统启动后的处理 */
	//systemBootHandle()
	loopCheck()

	//gorm.AddExtranetIp("120.230.98.231")
	//gorm.ExistExtranetIp("120.230.98.231")
	//gorm.AddPhoto("http://mmbiz.qpic.cn/mmbiz_jpg/jRPicmoSEZ5UvQshXWvAZuzSn6Kl4ySXlISdL6iacaKSicxtDdS3lCWUMj78mlu8qKiam7F1m1yRL3mzpRNYaXUX5Q/0", "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ")

	//log.Println("started finished ...")
	//r.Run(":8989")

	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(":8989", r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")

}

func systemBootHandle() bool {

	//bootTime := common.SystemBootTime()
	//log.Println("system boot timestamp : " + strconv.FormatUint(bootTime, 10))

	//if uint64(time.Now().Unix())-bootTime < 10*60 {
	/** 系统原型少于10分钟 */

	intranetIp := ip.GetIntranetIp()
	if intranetIp == "" {
		return false
	}

	log.Println("intranetIp is :" + intranetIp)

	/**  查看ngrok是否启动 */
	if proc.ExistProcName("ngrok") {
		log.Println("ngrok is running")
	} else {

		log.Println("ngrok is not running")
		bash.ExecShellCmd("sed -i '/web_addr:/cweb_addr: " + intranetIp + ":4040'  ~/.ngrok2/ngrok.yml")
		bash.ExecShellCmd("cd /home/labali/software/ && sh ./ngrok_start.sh")

		//https://ngrok.com/docs/ngrok-agent/api
		//curl http://192.168.10.11:4040/api/tunnels
		//获取外网映射地址
	}

	time.Sleep(time.Second * 3)

	ngrokInfo := bash.ExecShellCmd("curl http://" + intranetIp + ":4040/api/tunnels")

	//todo 解析json

	log.Println("ngrok info:" + ngrokInfo)

	/** 启动花生壳并发送二维码 */
	//https://service.oray.com/question/11644.html 好麻烦，回家扫好了。。。
	//TODO

	attachments := []mail.Attachment{}

	if proc.ExistProcName("phtunnel") {
		log.Println("phtunnel is running")

		///home/labali/aarch64-rpi3-linux-gnu
		//attachments = append(attachments, mail.Attachment{FilePath: "/home/labali/aarch64-rpi3-linux-gnu/phtunnel.log", Name: "phtunnel.log"})
	} else {

	}

	/** 发送邮件 */

	account := "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ"

	content := "内网ip地址：" + intranetIp + "<br/>"
	content += "外网地址信息：" + ngrokInfo + "<br/>"
	content += "weixin_app：8989, weixin_page:8787<br/>"

	wxmail.SendMail(account, "服务重启", content, attachments)

	//TODO
	//定时检查内网ip是否变更，作出相应处理，比如网络断了重连或网线被拔了重连
	//服务器启动快于网络连接，如何处理，定期检查，直到网络连接成功

	//}

	return true

}

func loopCheck() {
	//创建定时器，每隔1秒后，定时器就会给channel发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 5)

	go func() {
		for { //循环
			<-ticker.C

			if systemBootHandle() {
				ticker.Stop() //停止定时器
				break
			}
		}
	}() //别忘了()
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
	extranetIp := ip.GetExtranetIp()
	log.Println("extranetIp: " + extranetIp)
	//存数据库， 定时监测

	ticker := time.NewTicker(3600 * time.Second)
	for {
		<-ticker.C
		extranetIp := ip.GetExtranetIp()
		log.Println("extranetIp: " + extranetIp)

		if !gorm.ExistExtranetIp(extranetIp) {
			log.Println("extranetIp not exist notify ...")
			global.MailSender.SendMail([]string{"819966354@qq.com"}, "白名单不存在，请配置", extranetIp)
		}

	}
	//TODO 主动查外网ip的接口
}

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

func initLogger() {
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

//Gin还有很多功能，比如路由分组，自定义中间件，自动Crash处理等

func initConfig() {
	//yamlFile, err := ioutil.ReadFile("./config/config.yml")

	configPath := file.CurrentUserDir() + "/.weixin_app/config/config.yml"

	exist, err := file.PathExists(configPath)
	if err != nil {
		panic(err)
	}
	if !exist {
		log.Println(configPath + " is not exist")
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Println(err.Error())
	}
	var _config *config.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		log.Println(err.Error())
	}

}
