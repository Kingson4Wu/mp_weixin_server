package accesstoken_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/weixin/accesstoken"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	weixinConfig := config.GetWeixinConfig()
	weixinAccessToken := accesstoken.New(weixinConfig.Appid, weixinConfig.AppSecret)
	fmt.Println(weixinAccessToken.Get())
}
