package config

type Config struct {
	App *App `yaml:"app"`
	Log *Log `yaml:"log"`
	//Weixin *Weixin `yaml:"weixin"`
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

/*type Weixin struct {
	AppId     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
}*/

//https://zhuanlan.zhihu.com/p/261030657
