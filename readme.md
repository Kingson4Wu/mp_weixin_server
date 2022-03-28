go get -u github.com/gin-gonic/gin 下载gin，然后import导入即可。

go mod edit -require github.com/gin-gonic/gin@latest
go mod vendor

http://127.0.0.1:8989/

---
连接 sqlite
go-gorm-sqlite
https://zhuanlan.zhihu.com/p/388652094

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
gorm 的 sqlite 驱动，底层使用的还是 mattn/go-sqlite3 库。

---

go get gopkg.in/src-d/go-git.v4


---

### ngrok
+ ./ngrok http 8989

<pre>
Account                       kingson4wu@gmail.com (Plan: Free)
Version                       2.3.40
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://c71d-120-230-140-160.ngrok.io -> http://localhost:8989
Forwarding                    https://c71d-120-230-140-160.ngrok.io -> http://localhost:8989
</pre>

+ https://dashboard.ngrok.com/get-started/tutorials

+ https://stackoverflow.com/questions/36018375/how-to-change-ngroks-web-interface-port-address-not-4040
https://ngrok.com/docs#config-location
web_addr: localhost:4040

---

### Go之项目打包部署
打包成二进制文件，可以在linux平台运行

首先，进入到main.go文件目录下，执行以下命令

$ set GOARCH=amd64
$ set GOOS=linux

go build main.go

go build -o weixinapp main.go

复制代码go bulid 后就会在这个目录下生成打包好的Go项目文件了，是linux平台可执行的二进制文件。
将该文件放入linux系统某个文件夹下，chmod 773 [文件名] 赋予文件权限，./xx 命令即可执行文件，不需要go的任何依赖，就可以直接运行了。

作者：RunFromHere
链接：https://juejin.cn/post/6844903843201810446
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

scp -P 30022 weixinapp labali@192.168.10.8:/home/labali

chmod 773 weixinapp

---

env GOOS=linux GOARCH=arm64 go build main.go
这个才行！

详细介绍Go 交叉汇编 ARM:<https://zhuanlan.zhihu.com/p/319682047>

scp -P 30022 config/private_config.yml labali@192.168.10.8:/home/labali 

nohup ./main >/dev/null 2>&1 &

nohup ./ngrok http 8989 >/dev/null 2>&1 &

---

cat /proc/version
Linux version 4.4.194.pdnas.rk3328.258 (PD996@pdbolt.com) (gcc version 8.3.0 (GCC) ) #1 SMP Sat Sep 18 10:50:43 CST 2021
[labali@centos ~]$ uname -a
Linux centos 4.4.194.pdnas.rk3328.258 #1 SMP Sat Sep 18 10:50:43 CST 2021 aarch64 aarch64 aarch64 GNU/Linux

---

### 微信云托管
+ https://cloud.weixin.qq.com/cloudrun/console

---

一般我们都是在某个局域网内部，由于NAT的存在，其IP地址是经过转换的，那么如何得知转换后的公网IP呢？有两个在线工具可以帮你快速知道自己的外网地址，一个是国内的http://ip138.com，一个是国外的http://ifconfig.me。可以通过浏览器访问上面的站点查看，也可以通过curl工具查看：

StelladeMacBook-Air:~ stellazhou$ curl ifconfig.me
61.141.200.149
————————————————
版权声明：本文为CSDN博主「摩西2016」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/MoSee/article/details/77489677

---

+ go使用nacos作为配置中心
+ 【Golang】使用Go语言操作etcd——配置中心：<http://www.randyfield.cn/post/2021-05-04-go-etcd-config-center/>

+ go程序意外停止自动重启
+ linux服务器断电后自动重启服务
+ go程序发布到centos
+ 合master自动部署脚本，githook
+ web shell界面实现、安全校验

