package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kingson4wu/go-common-lib/file"
	"github.com/kingson4wu/mp_weixin_server/common"
	"gopkg.in/yaml.v2"
)

type PrivateConfig struct {
	WeixinConfig *WeixinConfig `yaml:"weixin"`
	MailConfig   *MailConfig   `yaml:"mail"`
	Labali       *Labali       `yaml:"labali"`
}

type Labali struct {
	Sss string `yaml:"sss"`
}

type WeixinConfig struct {
	Appid     string `yaml:"appid"`
	Appsecret string `yaml:"appsecret"`
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

func getYamlFileData() []byte {

	configPath := file.CurrentUserDir() + "/.weixin_app/config/private_config.yml"

	exist, err := file.PathExists(configPath)
	if err != nil {
		panic(err)
	}
	if !exist {
		log.Println(configPath + " is not exist")
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err.Error())
	}

	return yamlFile
}

func GetWeixinConfig() *WeixinConfig {
	yamlFile := getYamlFileData()
	var _config *PrivateConfig
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	_weixin := _config.WeixinConfig
	sss := _config.Labali.Sss
	_weixin.Appid, _ = common.DecryptByAesWithKey(_weixin.Appid, sss)
	_weixin.Appsecret, _ = common.DecryptByAesWithKey(_weixin.Appsecret, sss)
	_weixin.Token, _ = common.DecryptByAesWithKey(_weixin.Token, sss)

	return _weixin
}

func GetMailConfig() *MailConfig {
	yamlFile := getYamlFileData()
	var _config *PrivateConfig
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	_mail := _config.MailConfig
	sss := _config.Labali.Sss
	_mail.MailAddress, _ = common.DecryptByAesWithKey(_mail.MailAddress, sss)
	_mail.MailPass, _ = common.DecryptByAesWithKey(_mail.MailPass, sss)

	return _mail
}
