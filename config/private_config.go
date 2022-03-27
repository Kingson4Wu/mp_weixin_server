package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kingson4wu/weixin-app/common"
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

func getYamlFileData() []byte {

	configPath := common.CurrentUserDir() + "/.weixin_app/config/private_config.yml"

	exist, err := common.PathExists(configPath)
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

	return _config.WeixinConfig
}

func GetMailConfig() *MailConfig {
	yamlFile := getYamlFileData()
	var _config *PrivateConfig
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return _config.MailConfig
}
