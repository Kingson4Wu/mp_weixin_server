go get -u github.com/gin-gonic/gin 下载gin，然后import导入即可。

go mod edit -require github.com/gin-gonic/gin@latest
go mod vendor

http://127.0.0.1:8989/



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