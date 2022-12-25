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


+ vi ~/.ngrok2/ngrok.yml
```yml
authtoken: 26HiG7HktnlGuez6dzvXcgECyaD_55UEwnDYC4P6pS8RHRx3e
web_addr: 192.168.10.5:4040
tunnels:
  first:
    addr: 8989
    proto: http
  second:
    addr: 8787
    proto: http

```

+ ngrok start --all
+ nohup ./ngrok start --all >ngrok.log 2>&1 &

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

env GOOS=linux GOARCH=arm64 go build  -o ~/Downloads/weixinapp ./cmd/main.go
这个才行！


CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC="/usr/local/bin/x86_64-linux-musl-gcc" CGO_LDFLAGS="-static" go build -o ~/Downloads/weixinapp ./cmd/main.go
没成功

CGO_ENABLED=1 GOOS=linux CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ go build -a -installsuffix cgo -o ~/Downloads/weixinapp ./cmd/main.go
没成功，不支持arm64


scp -P 30022 ~/Downloads/weixinapp labali@192.168.10.8:/home/labali 
nohup ./weixinapp >/dev/null 2>&1 &

ps -ef|grep 'weixinapp'|grep -v 'grep'|awk '{ print $2}'|xargs kill -15

详细介绍Go 交叉汇编 ARM:<https://zhuanlan.zhihu.com/p/319682047>

scp -P 30022 config/private_config.yml labali@192.168.10.8:/home/labali 

nohup ./main >/dev/null 2>&1 &

nohup ./ngrok http 8989 >/dev/null 2>&1 &

`~/soft/sqlite-tools-osx-x86-3380100/sqlite3 ~/.weixin_app/db/wexin_app.db`

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


---

`docker run --rm -it -v ~/Personal/go-src/weixin-app/:/go/src/app  -v ~/Downloads/:/go/output gobuilder:1.17.7-stretch`

---

### 开机启动
+ vi /etc/rc.d/rc.local
+ /bin/su -labali -c  "/home/labali/.weixin_app/start.sh"

---

### vscode-debug
+ launch.json
```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceRoot}/cmd/main.go"
        }
    ]
}
```

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

### 花生壳
+ ` scp -P 30022 ~/Downloads/phddns_5.2.0_amd64.rpm labali@192.168.10.5:/home/labali`
+ install
<pre>
[labali@centos ~]$ su root
Password: 
[root@centos labali]# sudo yum localinstall /home/labali/phddns_5.2.0_amd64.rpm 
Loaded plugins: fastestmirror
Examining /home/labali/phddns_5.2.0_amd64.rpm: phddns-5.2.0-1.amd64
Cannot add package /home/labali/phddns_5.2.0_amd64.rpm to transaction. Not a compatible architecture: amd64
Nothing to do
</pre>

### aarch64架构如何安装花生壳？
+ https://www.zhihu.com/question/409997368
+ https://download.oray.com/peanuthull/sdk/latest/phtunnel-raspberry.tar.gz
+ https://hsk.oray.com/download 花生壳管理APP
+ cd /home/labali/aarch64-rpi3-linux-gnu
+ ./phtunnel
+ 手机app扫码登录
+ ctrl + z (暂停)
+ jobs (查询暂停的任务号)
+ bg %1 (将该任务号放入后台)
+ https://blog.csdn.net/u014609263/article/details/120703016

### natapp
+ `scp -P 30022 ~/Downloads/natapp  labali@192.168.10.5:/home/labali`


---

### centos

sudo vim /etc/ssh/sshd_config
systemctl restart sshd
netstat -tunlp | grep "ssh"

getconf LONG_BIT
64

mysql.root.611264

mysql -h127.0.0.1 -uroot -p611264
mysql -h192.168.10.8 -uroot -p611264

Linux centos 4.4.194.pdnas.rk3328.258 #1 SMP Sat Sep 18 10:50:43 CST 2021 aarch64 aarch64 aarch64 GNU/Linux

192.168.10.8 9202
root/labali611264
labali/611264

ssh labali@192.168.10.5 -p 30022

ssh labali@500i08756s.zicp.vip -p 11408

<pre>
[root@centos /]# df -h
Filesystem       Size  Used Avail Use% Mounted on
devtmpfs         962M     0  962M   0% /dev
tmpfs            980M     0  980M   0% /dev/shm
tmpfs            980M   26M  954M   3% /run
tmpfs            980M     0  980M   0% /sys/fs/cgroup
/dev/mmcblk2p17   13G  4.9G  7.0G  42% /
tmpfs            980M  4.0K  980M   1% /tmp
/dev/mmcblk2p16  511M  200M  312M  40% /boot
tmpfs            196M     0  196M   0% /run/user/0


2.5G	./www
1.9G	./usr
267M	./var
200M	./boot
117M	./root
39M	./etc
26M	./run
24K	./home
4.0K	./tmp
4.0K	./srv
</pre>

---



## TODO
+ 错误统一处理，日志打印
./gooooooooo/docs/grammar/error_handle.md
+ 远程调试
./gooooooooo/docs/golang_base/remote_debugging.md
+ pprof
./gooooooooo/docs/golang_base/pprof.md
+ 特别是一些 recover 之后什么都不做的代码，这种代码简直是毒瘤！当然，崩溃，可以是早一些向上传递 error，不一定就是 panic。同时，我要求大家不要在没有充分的必要性的时候 panic，应该更多地使用向上传递 error，做好 metrics 监控。合格的 golang 程序员，都不会在没有必要的时候无视 error，会妥善地做好 error 处理、向上传递、监控。一个死掉的程序，通常比一个瘫痪的程序，造成的损害要小得多。
+ 写测试用例，才发现自己隐藏的bug或者可以优化的地方


----

## 机器重启处理
+ http://wifi.cmcc/ 路由查看服务器ip


+ ./phtunnel
如何将二维码发送到邮件？？？

+ 将ip发送到邮件，花生壳app修改映射ip

+ 开机启动

+ 启动程序
1. 
cd /home/labali/.weixin_app
./start.sh
2. 
vi ~/.ngrok2/ngrok.yml
改当前ip
cd /home/labali/software
./ngrok_start.sh

3. https://dashboard.ngrok.com/tunnels/agents/ts_2BghCk6T37k8TWAwVGYiSU4Szok
找到外网链接
4. 微信公众号设置
https://mp.weixin.qq.com/advanced/advanced?action=interface&t=advanced/interface&token=1964219243&lang=zh_CN


----

### 自动化解决
1. 发送邮件：
（1）内网ip（用于花生壳设置）
（2）花生壳登录二维码，程序启动后重新登录了才能用
（3）外网地址（用于设置公众号）
（4）检查其他是否启动ok（可选）

```yml
authtoken: 26HiG7HktnlGuez6dzvXcgECyaD_55UEwnDYC4P6pS8RHRx3e
web_addr: 192.168.10.11:4040
tunnels:
  first:
    addr: 8989
    proto: http
  second:
    addr: 8787
    proto: http
```

+ `sed -i "/web_addr:/cweb_addr: 192.168.10.7:4041"  ~/.ngrok2/ngrok.yml` (linux执行不会报错)

+ 数据库地址

+ centos开机启动
+ 开机启动 weixin_app
+ vi /etc/rc.d/rc.local
/bin/su -labali -c  "/home/labali/.weixin_app/start.sh"
+ nohup /home/labali/.weixin_app/weixinapp >/dev/null 2>&1 &


---


### apk文件下载
```conf
location ^~ /labali/apk/download/ {
        alias /home/labali/.apk/;
 
        if ($request_filename ~* ^.*?\.(apk)$) {
            add_header Content-Disposition attachment;
            add_header Content-Type application/octet-stream;
        }
            sendfile on;   # 开启高效文件传输模式
            autoindex on;  # 开启目录文件列表
            autoindex_exact_size on;  # 显示出文件的确切大小，单位是bytes
            autoindex_localtime on;  # 显示的文件时间为文件的服务器时间
            charset utf-8,gbk;  # 避免中文乱码
      }
```

+ `http://192.168.10.11:8787/labali/apk/download/__UNI__B7FE310__20220731142411.apk`

---


## TODO
1. 工具：房贷计算，球队对战
2. 支持根据图片信息生成经纬度城市时间水印图片
3. go 加密工具  

