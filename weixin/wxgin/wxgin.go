package wxgin

import (
	"github.com/gin-gonic/gin"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxaction"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxsign"
	"log"
	"net/http"
)

const (
	path = "/labali_msg"
)

func Handle(r *gin.Engine) {

	r.GET(path, func(context *gin.Context) {

		//http://127.0.0.1:8989/access_check_signature?signature=4654fdg&timestamp=3534&nonce=35fdgf
		signature := context.Query("signature")
		timestamp := context.Query("timestamp")
		nonce := context.Query("nonce")
		echostr := context.Query("echostr")

		//3）开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
		if wxsign.Check(signature, timestamp, nonce) {
			context.String(http.StatusOK, echostr)
		} else {
			context.String(http.StatusOK, "")
		}
	})

	r.POST(path, func(context *gin.Context) {

		//http://127.0.0.1:8989/access_check_signature?signature=4654fdg&timestamp=3534&nonce=35fdgf
		signature := context.Query("signature")
		timestamp := context.Query("timestamp")
		nonce := context.Query("nonce")
		//echostr := context.Query("echostr")
		//openid := context.Query("openid")

		//fmt.Println(signature)
		//fmt.Println(timestamp)
		//fmt.Println(nonce)
		//fmt.Println(openid)

		if wxsign.Check(signature, timestamp, nonce) {

			receiveMsg := wxaction.WXMsgReceive(context)
			wxaction.HandleMsg(receiveMsg, context)

		} else {
			context.String(http.StatusOK, "")
			log.Println("replyText failure")
		}
	})

}
