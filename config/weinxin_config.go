package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type WeixinConfig struct {
	AppId     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
	Token     string `yaml:"token"`
}

func GetWeixinConfig() *WeixinConfig {
	yamlFile, err := ioutil.ReadFile("./config/weixin_config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *WeixinConfig
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Printf("config.wexin: %#v\n", _config)

	return _config
}
