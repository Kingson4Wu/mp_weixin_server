#!/bin/bash 

ls_date=`date +%Y%m%d%H%M%S`
dest='labali@192.168.10.44:/home/labali/.weixin_app/pkg/weixinapp_'$ls_date
src=`ls ~/Downloads/weixinapp`
pass=611264 

expect -c "
    spawn scp -P 30022 $src $dest
    expect {
        \"*assword\" {set timeout 300; send \"$pass\r\"; exp_continue;}
        \"yes/no\" {send \"yes\r\";exp_continue } 
    }
"

