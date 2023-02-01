
+ 只考虑单机服务

## 功能点

### weixin access token
+ 当需要主动操作微信公众号时，需要调用公众号提供的API，调用API需要使用accessToken
+ weixin/accesstoken/access_token.go

### 发送邮件
+ mail/mail.go

### 数据存储

+ 若使用sqllite，还可以考虑定时备份到git，common/backup/backup.go



+ 读取yaml格式配置