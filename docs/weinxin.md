curl -d '{
     "button":[
     {	
          "type":"click",
          "name":"查一下",
          "key":"V1001_check"
      },
      {
           "name":"菜单",
           "sub_button":[
           {	
               "type":"view",
               "name":"搜索",
               "url":"http://www.bing.com/"
            },
            {
                 "type":"miniprogram",
                 "name":"wxa",
                 "url":"http://mp.weixin.qq.com",
                 "appid":"wx286b93c14bbf93aa",
                 "pagepath":"pages/lunar/index"
             },
            {
               "type":"click",
               "name":"赞一下我们",
               "key":"V1001_GOOD"
            }]
       }]
 }' 'https://api.weixin.qq.com/cgi-bin/menu/create?access_token=55'

 {"errcode":48001,"errmsg":"api unauthorized rid: 6235eacb-33473b07-03d53927"}%

 没认证，没权限

### 微信云托管
+ https://cloud.weixin.qq.com/cloudrun/console
