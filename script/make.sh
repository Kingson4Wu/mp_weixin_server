env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"  -o ~/Downloads/weixinapp ./cmd/main.go && upx -9 ~/Downloads/weixinapp

# brew install upx

#ps -ef|grep 'weixinapp'|grep -v 'grep'|awk '{ print $2}'|xargs kill -15
