# MP Weixin Server

A comprehensive WeChat public account server implemented in Go, providing various features such as todo management, scheduled notifications, and media backup with external network access via ngrok.

## Architecture

![](https://raw.githubusercontent.com/Kingson4Wu/mp_weixin_server/main/docs/image/weixin_app_architecture.drawio.png)

1. Service deployed on a home network small server (including Golang program, MySQL, etc.)
2. External network access provided through ngrok proxy for WeChat public account, uni-app, etc.
3. Message notifications pushed to WeChat via email sending

## Deployment Flow

![](https://raw.githubusercontent.com/Kingson4Wu/mp_weixin_server/main/docs/image/weixin_app_deploy.drawio.png)

1. Cross-compile to generate arm64 Linux binary package on development machine
2. Transfer via SCP to small server
3. Send signal to restart service

### Development in External Network Environment
+ Using花生壳 (Phddns) software to provide mapping services
    1. Map MySQL to allow development machine to start service normally
    2. Map small server SSH port to allow normal login and file transfer

## Deployment and Configuration

### Deploy Service

1. `vi ~/.weixin_app/config/config.yml` - General configuration (MySQL database configuration)
```yml
database:
  username: user
  password: pass
  host: 192.168.33.174
  port: 3306
  dbname: weixin_app
  timeout: 10s
```

2. `vi ~/.weixin_app/config/private_config.yml` - Private configuration
```yml
encrypt: # For encrypting/decrypting sensitive information
  key: 123456... # (16 characters)
weixin: # WeChat Open Platform configuration
  appid: xxxx # Fill in after AES encryption using encrypt.key
  appSecret: xxxx # Fill in after AES encryption using encrypt.key
  token: xxxx # Fill in after AES encryption using encrypt.key
mail: # For sending notifications
  address: fffff # Fill in after AES encryption using encrypt.key
  pass: fff # Fill in after AES encryption using encrypt.key (email password or authorization code)
  name: Labali # Sender user name
  smtpHost: smtp.qq.com
  receiverList: [ # Find corresponding email by WeChat openid for sending email notifications
    openId: "xxx", # (WeChat public account openid)
    address: "xxx@qq.com"
  ,{
    openId: "ddd",
    address: "dd@qq.com"
  }
  ]
admin:
  accounts: [ # Administrator WeChat public account openid 
    "oqV-xxxxx",
    "oqV-xxxxx"
  ] 
```

3. `go run cmd/main.go`
4. Access http://127.0.0.1:8989/ to check if startup was successful

### Configure External Network Proxy
+ Use ngrok to configure proxy for external network access

### WeChat Public Account Configuration
+ Go to WeChat public account backend, configure development callback link
    - Settings & Development -> Basic Configuration -> Server Configuration -> Server Address(URL)
      https://xxxx.ngrok.io/labali_msg

## Features

### Basic Features
1. Add, view, delete todo items
2. Scheduled WeChat message reminders (implemented via email sending, WeChat needs to have email reminders enabled)
3. Save sent images or videos to server, and summarize and send email backup the next day

### Implementation Highlights
1. Connect with WeChat Open Platform, receive passive messages, process and reply
2. Query and cache WeChat public account access token for scenarios where active requests are sent to WeChat public account server
3. Send email notifications to eligible WeChat accounts
4. Sensitive configurations use AES encryption/decryption
5. Receive WeChat image or video links, download to corresponding local server location

### Operations Support
1. Server external IP periodic check, send email notification if changed (querying WeChat public account accessToken requires configuring whitelist IP)
    - Settings & Development -> Basic Configuration -> Public Account Development Information -> IP Whitelist
    - This feature is not fully implemented
2. Service graceful restart on signal reception
3. Send email notification when service starts successfully, with service information (external proxy address, external IP, internal IP, etc.)
4. Check if ngrok process is running when service starts, trigger start if not
5. Service deployment and update scripts
    - `./script/make.sh` - Package binary files
    - `./script/upload.sh` - Upload to server
    - `./script/deploy.sh` - Restart service
6. Periodic check of server internal IP, send email notification if changed, for external network connection via Peanut Shell for SSH to home server
7. Server restart, configure automatic service startup (centos startup item configuration)

## Add New Commands for Passive Message Scenarios
+ weixin/wxaction/wxaction.go
```go
registerHandler(cmd, func(openid, content string) string {
return "msg"
}) 
```

## Learning Points
1. Examples of using gorm
2. How to connect with WeChat public account and receive passive messages
3. How to actively query information through WeChat public account API
4. Usage of scheduled task github.com/robfig/cron/v3
5. Using github.com/fvbock/endless for graceful service restart
6. Examples of using gopkg.in/yaml.v2 for reading YAML configuration files

## Todo List
1. Store user openid and corresponding email in database, currently configured in private_config.yml#mail.receiverList
2. Store administrator openid in database, currently configured in private_config.yml#admin.accounts
3. Make scheduled email reminder times configurable, cron/cron.go
4. Auto-start, start Peanut Shell and send QR code to email, scan to SSH to server from external network
5. Auto-deployment script on merging to master, git hook triggers service update and restart
6. Add user, block user, send message notifications to administrator account id
7. Command context memory capability
8. Integrate configuration center, such as nacos
9. Saved images, support generating geolocation city time watermark based on image information

## Notes
+ Current code implementation only considers single-server service
+ If using sqlite, consider periodic backup to git, common/backup/backup.go

## License

This project is licensed under the terms of the LICENSE file.

---

For more information, visit [deepwiki](https://deepwiki.com/Kingson4Wu/mp_weixin_server)