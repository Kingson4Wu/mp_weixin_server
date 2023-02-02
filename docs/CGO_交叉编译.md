
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC="/usr/local/bin/x86_64-linux-musl-gcc" CGO_LDFLAGS="-static" go build -o ~/Downloads/weixinapp ./cmd/main.go
没成功

CGO_ENABLED=1 GOOS=linux CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ go build -a -installsuffix cgo -o ~/Downloads/weixinapp ./cmd/main.go
没成功，不支持arm64


详细介绍Go 交叉汇编 ARM:<https://zhuanlan.zhihu.com/p/319682047>
