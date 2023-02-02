
## 部署和配置

### 部署服务

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
  name: 拉巴力 # 发送用户名
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
  accounts: [ # 管理员 微信公众号openid 
    "oqV-xxxxx",
    "oqV-xxxxx"
  ] 
```

3. ` go run cmd/main.go`
4. 访问 http://127.0.0.1:8989/ 查看是否启动成功

### 配置外网代理
+ 使用ngrok，配置代理，这样外网才能访问

### 微信公众号配置
+ 到微信公众号后台，配置开发回调链接
    - 设置与开发 -> 基本配置 -> 服务器配置 -> 服务器地址(URL)
      https://xxxx.ngrok.io/labali_msg

## 功能分类

### 基本功能
1. 添加、查看、删除todo事项
2. 定时微信消息提醒（通过发邮件方式实现，微信需要设置开启邮件提醒）
3. 保存发送的图片或视频到服务器，并在第二天汇总发送邮件备份

### 实现要点
1. 对接微信开放平台，接收被动消息，处理并回复
2. 查询并缓存微信公众号的access token，用于主动发送请求到微信公众号服务器的场景
3. 给符合条件的微信号发送邮件通知
4. 敏感配置使用AES加解密
5. 接收微信的图片或视频链接，下载到本地服务器相应的位置

### 运维支持
1. 服务器外网ip定时检查，若有变更发送邮件通知（查询微信公众号的accessToken，需要配置白名单ip）
    - 设置与开发 -> 基本配置 -> 公众号开发信息 -> IP白名单
    - 改功能未完全实现
2. 服务接收信号优雅重启
3. 服务启动成功发送邮件通知，并附带服务相关信息（外网代理地址，外网ip，内网ip等）
4. 服务启动时检查ngrok进程是否运行，否则触发启动
5. 服务部署和更新脚本
    - `./script/make.sh` 打包二进制文件
    - `./script/upload.sh` 上传服务器
    - `./script/deploy.sh` 开始重启
6. 服务器内网ip定时检查，若有变更发送邮件通知，方便在外网通过花生壳链接，进行ssh登录到家庭服务器
7. 服务器重启，配置服务自动启动（centos启动项配置）

## 被动消息场景如何新增指令
+ weixin/wxaction/wxaction.go
```go
registerHandler(cmd, func(openid, content string) string {
return "msg"
}) 
```

## 可以学习
1. 使用gorm的例子
2. 如何对接微信公众号接收被动消息
3. 如何主动通过微信公众号API发送请求查询信息
4. 定时任务github.com/robfig/cron/v3的使用
5. 使用github.com/fvbock/endless对服务进行优雅重启
6. 使用gopkg.in/yaml.v2读取yaml配置文件例子

## todo
1. 用户openid和对应的邮箱放在数据库维护，目前是配置在private_config.yml#mail.receiverList
2. 管理员openid放在数据库维护，目前是配置在private_config.yml#admin.accounts
3. 定时邮件提醒时间可配置，cron/cron.go
4. 自动启动，启动花生壳并发送二维码到邮件，扫码后可以在外网ssh到服务器
5. 合master自动部署脚本，git hook触发服务更新并重启
6. 添加用户、 拉黑用户
7. 指令上下文记忆能力
8. 接入配置中心，比如nacos
9. 保存的图片，支持根据图片信息生成经纬度城市时间水印

## 说明
+ 代码目前的实现，只考虑单机服务
+ 若使用sqllite，还可以考虑定时备份到git，common/backup/backup.go

---

1. releases
2. packages
3. new branch
4. 





