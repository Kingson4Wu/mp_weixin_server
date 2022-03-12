package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})

	r.GET("/access_check_signature", func(context *gin.Context) {

		//http://127.0.0.1:8989/access_check_signature?signature=4654fdg&timestamp=3534&nonce=35fdgf
		signature := context.Query("signature")
		timestamp := context.Query("timestamp")
		nonce := context.Query("nonce")
		echostr := context.Query("echostr")

		fmt.Println(signature)
		fmt.Println(timestamp)
		fmt.Println(nonce)
		fmt.Println(echostr)

		//74aac4f5140093d830c3c11b70fdfae86d689b4a
		//1647091983
		//2084438825
		//6200237875618768367

		//1）将token、timestamp、nonce三个参数进行字典序排序

		token := "123456"
		myList := []string{token, timestamp, nonce}

		fmt.Printf("Before: %v\n", myList)

		// Pass in our list and a func to compare values
		sort.Slice(myList, func(i, j int) bool {
			numA, _ := strconv.Atoi(myList[i])
			numB, _ := strconv.Atoi(myList[j])
			return numA < numB
		})
		sb := ""
		for _, str := range myList {
			sb += str
		}

		//2.1 sha1加密
		sha1 := SHA1(sb)
		fmt.Println("sha1:" + sha1)
		//3）开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
		if sha1 == signature {
			context.String(http.StatusOK, echostr)
		} else {
			context.String(http.StatusOK, "")
		}

	})

	r.Run(":8989")
}

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

//Gin还有很多功能，比如路由分组，自定义中间件，自动Crash处理等
