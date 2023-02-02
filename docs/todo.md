
## TODO
+ 错误统一处理，日志打印
  ./gooooooooo/docs/grammar/error_handle.md
+ 远程调试
  ./gooooooooo/docs/golang_base/remote_debugging.md
+ pprof
  ./gooooooooo/docs/golang_base/pprof.md
+ 特别是一些 recover 之后什么都不做的代码，这种代码简直是毒瘤！当然，崩溃，可以是早一些向上传递 error，不一定就是 panic。同时，我要求大家不要在没有充分的必要性的时候 panic，应该更多地使用向上传递 error，做好 metrics 监控。合格的 golang 程序员，都不会在没有必要的时候无视 error，会妥善地做好 error 处理、向上传递、监控。一个死掉的程序，通常比一个瘫痪的程序，造成的损害要小得多。
+ 写测试用例，才发现自己隐藏的bug或者可以优化的地方


### 开机启动

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
  /usr/sbin/runuser -l labali -c  "/home/labali/.weixin_app/start.sh"
+ nohup /home/labali/.weixin_app/weixinapp >/dev/null 2>&1 &

vi weixin_app
```shell
#!/bin/bash
# chkconfig: 2345 10 90
# description: weixin_app ....
sleep 15
/usr/sbin/runuser -l labali -c  "/home/labali/.weixin_app/start.sh"
```


## TODO
1. 工具：房贷计算，球队对战
2. 支持根据图片信息生成经纬度城市时间水印图片
3. go 加密工具 