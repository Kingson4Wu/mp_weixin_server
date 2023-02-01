package wxserver

import (
	"github.com/kingson4wu/mp_weixin_server/common/ip"
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/global"
	"github.com/kingson4wu/mp_weixin_server/gorm"
	"log"
	"time"
)

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
			global.MailSender.SendMail([]string{config.GetMailConfig().MailAddress}, "白名单不存在，请配置", extranetIp)
		}

	}
	//TODO 主动查外网ip的接口
}
