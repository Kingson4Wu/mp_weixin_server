
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


+ 读取yaml格式配置