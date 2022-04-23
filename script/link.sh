src=`ls -lt ~/.weixin_app/pkg | grep weixinapp | head -n 1 |awk '{print $9}'`

ln -snf ~/.weixin_app/pkg/$src ~/.weixin_app/weixinapp