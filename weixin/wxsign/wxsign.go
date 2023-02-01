package wxsign

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/kingson4wu/mp_weixin_server/config"
	"sort"
)

func Check(signature string, timestamp string, nonce string) bool {
	//1）将token、timestamp、nonce三个参数进行字典序排序

	config := config.GetWeixinConfig()
	token := config.Token

	//将token、timestamp、nonce三个参数进行字典序排序
	var tempArray = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))

	//fmt.Println("token:" + token)
	//fmt.Println("timestamp:" + timestamp)
	//fmt.Println("nonce:" + nonce)
	//fmt.Println("nsha1once:" + sha1String)
	//fmt.Println("signature:" + signature)

	//获得加密后的字符串可与signature对比
	return sha1String == signature
}
