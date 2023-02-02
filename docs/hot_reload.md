### hot reload
+ go install github.com/cosmtrek/air@latest
+ air init
+ vi .air.toml
+   cmd = "go build -o ./tmp/main ./cmd/main.go"
+ air
+ crtl + c
+ ps -ef |grep weixin
+ attach process id
+ debug launch.json:
```json
{
            "name": "Attach to Process",
            "type": "go",
            "request": "attach",
            "mode": "local",
            "processId": 64088
        }
```
+ 未试！