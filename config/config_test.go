package config_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/config"
	"testing"
)

func TestGetDatabaseConfig(t *testing.T) {

	fmt.Println(config.GetWeixinConfig().Appid)
	fmt.Println(config.GetMailConfig().MailAddress)
}
