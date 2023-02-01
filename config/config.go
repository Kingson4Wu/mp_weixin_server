package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/kingson4wu/go-common-lib/file"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Log      *Log      `yaml:"log"`
	Database *Database `yaml:"database"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Timeout  string `yaml:"timeout"`
}

//https://zhuanlan.zhihu.com/p/261030657

var (
	yamlFileData []byte
	once         = new(sync.Once)
)

func getYamlFileConfigData() []byte {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("getYamlFileConfigData error: %s \n", err)
			once = new(sync.Once)
		}
	}()

	if yamlFileData != nil {
		return yamlFileData
	}

	once.Do(func() {

		configPath := file.CurrentUserDir() + "/.weixin_app/config/config.yml"
		log.Printf("read file : %s\n", configPath)
		
		exist, err := file.PathExists(configPath)
		if err != nil {
			panic(err)
		}
		if !exist {
			panic(configPath + " is not exist")
		}

		b, err := os.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		yamlFileData = b
	})

	if yamlFileData == nil {
		once = new(sync.Once)
	}

	return yamlFileData
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
