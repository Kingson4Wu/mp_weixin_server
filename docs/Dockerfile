FROM golang:1.17.7-stretch
MAINTAINER author <email@xxx.com>

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    wget \
    vim \
    htop \
    curl \
    sudo \
    git \
    net-tools \
    tzdata \
    gcc-aarch64-linux-gnu \
    libc6-dev \
    libc6-dev-i386 \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && rm -rf /var/lib/apt/lists/* \
    && ln -s /usr/include/asm-generic /usr/include/asm \
    && sed 's/-Werror/ /g' /usr/local/go/src/runtime/cgo/cgo.go
ENV CFLAGS="-Wno-error"
ENV CXXFLAGS="-Wno-error"
ENV GOBIN=$GOPATH/bin
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
ENV TZ=Asia/Shanghai

WORKDIR /go/src/app

ENTRYPOINT ["env", "CGO_LDFLAGS='-static'", "CGO_ENABLED=1", "CC=aarch64-linux-gnu-gcc", "GOOS=linux", "GOARCH=arm64", "go", "build", "-o", "/go/output/weixin-app", "./cmd/main.go"]