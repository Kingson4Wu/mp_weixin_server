env GOOS=linux GOARCH=arm64 go build  -o ~/Downloads/weixinapp ./cmd/main.go

#ps -ef|grep 'weixinapp'|grep -v 'grep'|awk '{ print $2}'|xargs kill -15
