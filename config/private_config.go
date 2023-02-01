package config

import (
	"github.com/kingson4wu/mp_weixin_server/common/aes"
	"log"
	"os"
	"sync"

	"github.com/kingson4wu/go-common-lib/file"
	"gopkg.in/yaml.v2"
)

type PrivateConfig struct {
	WeixinConfig *WeixinConfig `yaml:"weixin"`
	MailConfig   *MailConfig   `yaml:"mail"`
	Encrypt      *Encrypt      `yaml:"encrypt"`
	AdminConfig  *AdminConfig  `yaml:"admin"`
}

type Encrypt struct {
	Key string `yaml:"key"`
}

type WeixinConfig struct {
	Appid     string `yaml:"appid"`
	AppSecret string `yaml:"appSecret"`
	Token     string `yaml:"token"`
}

type MailConfig struct {
	MailAddress   string         `yaml:"address"`
	MailName      string         `yaml:"name"`
	MailPass      string         `yaml:"pass"`
	SmtpHost      string         `yaml:"smtpHost"`
	UserMailInfos []UserMailInfo `yaml:"receiverList"`
}

type UserMailInfo struct {
	OpenId  string `yaml:"openId"`
	Address string `yaml:"address"`
}

type AdminConfig struct {
	Accounts []string `yaml:"accounts"`
}

var (
	privateConfig *PrivateConfig
	privateOnce   = new(sync.Once)
)

func getConfig() *PrivateConfig {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("getConfig error: %s \n", err)
			once = new(sync.Once)
		}
	}()

	if privateConfig != nil {
		return privateConfig
	}

	privateOnce.Do(func() {
		b := getYamlFileData()
		var _config *PrivateConfig
		err := yaml.Unmarshal(b, &_config)
		if err != nil {
			panic(err)
		}
		privateConfig = _config

	})

	if privateConfig == nil {
		privateOnce = new(sync.Once)
	}

	return privateConfig
}

func getYamlFileData() []byte {

	configPath := file.CurrentUserDir() + "/.weixin_app/config/private_config.yml"
	log.Printf("read file : %s\n", configPath)

	exist, err := file.PathExists(configPath)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic(configPath + " is not exist")
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	return yamlFile
}

func GetWeixinConfig() *WeixinConfig {
	_config := getConfig()

	_weixin := _config.WeixinConfig
	sss := _config.Encrypt.Key
	_weixin.Appid, _ = aes.DecryptByAesWithKey(_weixin.Appid, sss)
	_weixin.AppSecret, _ = aes.DecryptByAesWithKey(_weixin.AppSecret, sss)
	_weixin.Token, _ = aes.DecryptByAesWithKey(_weixin.Token, sss)

	return _weixin
}

func GetMailConfig() *MailConfig {
	_config := getConfig()

	_mail := _config.MailConfig
	sss := _config.Encrypt.Key
	_mail.MailAddress, _ = aes.DecryptByAesWithKey(_mail.MailAddress, sss)
	_mail.MailPass, _ = aes.DecryptByAesWithKey(_mail.MailPass, sss)

	return _mail
}

func GetAdminConfig() *AdminConfig {
	_config := getConfig()

	return _config.AdminConfig
}
