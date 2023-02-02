cat /proc/version
Linux version 4.4.194.pdnas.rk3328.258 (PD996@pdbolt.com) (gcc version 8.3.0 (GCC) ) #1 SMP Sat Sep 18 10:50:43 CST 2021
[labali@centos ~]$ uname -a
Linux centos 4.4.194.pdnas.rk3328.258 #1 SMP Sat Sep 18 10:50:43 CST 2021 aarch64 aarch64 aarch64 GNU/Linux

### docker

`docker run --rm -it -v ~/Personal/go-src/mp_weixin_server/:/go/src/app  -v ~/Downloads/:/go/output gobuilder:1.17.7-stretch`

### 开机启
+ vi /etc/rc.d/rc.local
+ /bin/su -labali -c  "/home/labali/.weixin_app/start.sh"

## 机器重启处理
+ http://wifi.cmcc/ 路由查看服务器ip

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