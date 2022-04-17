package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kingson4wu/go-common-lib/file"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App      *App      `yaml:"app"`
	Log      *Log      `yaml:"log"`
	Database *Database `yaml:"database"`
}

type App struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}

type Database struct {
	Username string `yaml:"username"` //账号
	Password string `yaml:"password"` //密码
	Host     string `yaml:"host"`     //数据库地址，可以是Ip或者域名
	Port     int    `yaml:"port"`     //数据库端口
	Dbname   string `yaml:"dbname"`   //数据库名
	Timeout  string `yaml:"timeout"`  //连接超时，10秒

}

/*type Weixin struct {
	AppId     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
}*/

//https://zhuanlan.zhihu.com/p/261030657

func getYamlFileConfigData() []byte {

	configPath := file.CurrentUserDir() + "/.weixin_app/config/config.yml"

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

func GetDatabaseConfig() *Database {
	yamlFile := getYamlFileConfigData()
	var _config *Config
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	_database := _config.Database

	return _database
}
