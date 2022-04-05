nohup ./weixinapp >/dev/null 2>&1 &

ps -ef|grep 'weixinapp'|grep -v 'grep'|awk '{ print $2}'|xargs kill -15
