package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PrivateConfig struct {
	WeixinConfig *WeixinConfig `yaml:"weixin"`
	MailConfig   *MailConfig   `yaml:"mail"`
}

type WeixinConfig struct {
	Appid     string `yaml:"appid"`
	Appsecret string `yaml:"appsecret"`
	Token     string `yaml:"token"`
}

type MailConfig struct {
	User          string         `yaml:"user"`
	Pass          string         `yaml:"pass"`
	UserMailInfos []UserMailInfo `yaml:"receiverList"`
}

type UserMailInfo struct {
	OpenId  string `yaml:"openId"`
	Address string `yaml:"address"`
}

func GetWeixinConfig() *WeixinConfig {
	yamlFile, err := ioutil.ReadFile("./config/private_config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *PrivateConfig
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _config.WeixinConfig
}

func GetMailConfig() *MailConfig {
	yamlFile, err := ioutil.ReadFile("./config/private_config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *PrivateConfig
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _config.MailConfig
}
