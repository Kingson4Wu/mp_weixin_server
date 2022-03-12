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
