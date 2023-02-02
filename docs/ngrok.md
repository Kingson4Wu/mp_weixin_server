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
authtoken: 26HiG7H8RHRx3e
web_addr: 192.168.10.2:4040
tunnels:
  first:
    addr: 8989
    proto: http
  second:
    addr: 8787
    proto: http
    bind_tls: true
  third:
    addr: 8988
    proto: http
    bind_tls: true
```
+ `sed -i "/web_addr:/cweb_addr: 192.168.10.7:4041"  ~/.ngrok2/ngrok.yml` (linux执行不会报错)

+ 最多只能配四个通道, 包括http、https, 已经配了4个了

+ ngrok start --all
+ nohup ./ngrok start --all >ngrok.log 2>&1 &

+ ngrok 中文文档: https://www.jishuchi.com/read/ngrok/5107

+ nohup ./ngrok http 8989 >/dev/null 2>&1 &