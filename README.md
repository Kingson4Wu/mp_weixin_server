
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
