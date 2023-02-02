
## 如何运行
### 配置文件准备
1. `vi ~/.weixin_app/config/config.yml` 通用配置（MySQL数据库配置）
```yml
database:
  username: user
  password: pass
  host: 192.168.33.174
  port: 3306
  dbname: weixin_app
  timeout: 10s
```

2. `vi ~/.weixin_app/config/private_config.yml` 配置私有配置
```yml
encrypt: # 用于加解密敏感信息
  key: 123456... #（16位字符）
weixin: # 微信开放平台的配置
  appid: xxxx #使用encrypt.key AES加密后填写
  appSecret: xxxx #使用encrypt.key AES加密后填写
  token: xxxx #使用encrypt.key AES加密后填写
mail: #用于发送通知的邮件
  address: fffff  #使用encrypt.key AES加密后填写
  pass: fff #使用encrypt.key AES加密后填写 # 邮箱密码或授权码
  name: 拉巴力
  smtpHost: smtp.qq.com
  receiverList: [{ # 通过微信openid找到对应的邮箱，用于发送邮件通知
    openId: "xxx", #（微信公众号openid）
    address: "xxx@qq.com"
  },{
    openId: "ddd",
    address: "dd@qq.com"
  }
  ]
admin:
  accounts: [ # 使用者（具备功能使用权限） 微信公众号openid 
    "oqV-XjlEcZZcA4pCwoaiLtnFF0XQ",
    "oqV-Xju6thOVtzvi0FrTWHaB5So4"
  ] 
```

## todo
1. 用户openid和对应的邮箱放在数据库维护，目前是配置在private_config.yml#mail.receiverList
2. 

http://127.0.0.1:8989/

+ 只考虑单机服务

## 功能点

### weixin access token
+ 当需要主动操作微信公众号时，需要调用公众号提供的API，调用API需要使用accessToken
+ weixin/accesstoken/access_token.go

### 发送邮件
+ mail/mail.go

### 数据存储

+ 若使用sqllite，还可以考虑定时备份到git，common/backup/backup.go

### SQL 结构



### 定时发送昨日的图片到邮件 （时光机）
+ job/cron_task.go

### 定时TODO计划


### 对接微信公众号提供以下功能
1. 记录todo，查询todo
2. 记录照片
3. 查询机器内网外网ip


## 运维相关
### 开机启动




+ 读取yaml格式配置

## TODO
### 提供开放体验入口
+ 输入指令重新申请
+ 添加用户
+ 拉黑用户
+ 指令上下文记忆能力
