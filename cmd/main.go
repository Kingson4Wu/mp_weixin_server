package main

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/bash"
	"github.com/kingson4wu/mp_weixin_server/common/ip"
	"github.com/kingson4wu/mp_weixin_server/common/proc"
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/global"
	"github.com/kingson4wu/mp_weixin_server/logger"
	"github.com/kingson4wu/mp_weixin_server/ngrok"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxgin"
	"log"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kingson4wu/mp_weixin_server/cron"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"github.com/kingson4wu/mp_weixin_server/mail"
)

func init() {
	logger.InitLogger()
	gorm.InitDB()
	job.InitCron()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"labali": "koo~",
			"fox":    "shit",
		})
	})
	wxgin.Handle(r)

	/** 系统启动后的处理 */
	loopCheck()

	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(":8989", r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exit.")

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

	ngrokInfo := bash.ExecShellCmd(fmt.Sprintf("curl -s http://%s:4040/api/tunnels", intranetIp))
	log.Println("ngrok info:" + ngrokInfo)

	ngrokText := ngrokInfo
	ngrokText = ngrok.Parse(ngrokInfo)
	log.Printf("ngrok parse result:%s\n", ngrokText)

	/** 启动花生壳并发送二维码 */
	//https://service.oray.com/question/11644.html 好麻烦，回家扫好了。。。
	//TODO

	var attachments []mail.Attachment

	if proc.ExistProcName("phtunnel") {
		log.Println("phtunnel is running")

		///home/labali/aarch64-rpi3-linux-gnu
		//attachments = append(attachments, mail.Attachment{FilePath: "/home/labali/aarch64-rpi3-linux-gnu/phtunnel.log", Name: "phtunnel.log"})
	} else {

	}

	/** 发送邮件 */
	log.Println("server is started ...")

	content := fmt.Sprintf("内网ip地址：%s<br/>", intranetIp)
	content += fmt.Sprintf("外网ip地址：%s<br/>", ip.GetExtranetIp())
	content += fmt.Sprintf("外网地址信息：%s<br/>", ngrokText)
	content += "weixin_app：8989, weixin_page:8787<br/>"

	global.MailSender.SendMailWithAttachment([]string{config.GetMailConfig().MailAddress}, "服务重启", content, attachments)

	log.Println("send start server email ...")

	//TODO
	//定时检查内网ip是否变更，作出相应处理，比如网络断了重连或网线被拔了重连
	//服务器启动快于网络连接，如何处理，定期检查，直到网络连接成功

	return true
}

func loopCheck() {
	//创建定时器，每隔1秒后，定时器就会给channel发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 5)

	go func() {
		for {
			<-ticker.C

			if systemBootHandle() {
				ticker.Stop() //停止定时器
				break
			}
		}
	}()
}
