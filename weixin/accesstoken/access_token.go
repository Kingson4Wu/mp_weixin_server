package accesstoken

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kingson4wu/go-common-lib/file"
)

type weixinAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type accessTokenStore struct {
	AccessToken       string `json:"access_token"`
	ExpireTimeSeconds int    `json:"expire_time_seconds"`
}

func exists(path string) bool {

	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

type AccessToken struct {
	appid     string
	appSecret string
}

func New(appid, appSecret string) *AccessToken {
	return &AccessToken{appid: appid, appSecret: appSecret}
}

func (ac *AccessToken) Get() string {

	storeDirPath := strings.Join([]string{file.CurrentUserDir(), ".weixin_app", "work", ac.appid}, fmt.Sprintf("%c", filepath.Separator))
	storePath := strings.Join([]string{storeDirPath, "access_token_store.json"}, fmt.Sprintf("%c", filepath.Separator))

	if exists(storePath) {
		dat, err := os.ReadFile(storePath)
		check(err)
		tokenStore := accessTokenStore{}
		json.Unmarshal(dat, &tokenStore)

		if tokenStore.ExpireTimeSeconds > int(time.Now().Unix()) {
			fmt.Printf("read accessToken from file success\n")
			return tokenStore.AccessToken
		}
	}

	fmt.Printf("read accessToken from file remote\n")
	resp := getRemoteAccessToken(ac.appid, ac.appSecret)

	if resp == nil {
		fmt.Printf("getRemoteAccessToken error !\n")
		return ""
	}

	store := accessTokenStore{AccessToken: resp.AccessToken, ExpireTimeSeconds: int(time.Now().Unix()) + resp.ExpiresIn}

	storeJson, _ := json.Marshal(store)
	storeJsonStr := string(storeJson)

	fmt.Printf("config.storeJsonStr: %s\n", storeJsonStr)

	if !exists(storeDirPath) {
		os.Mkdir(storeDirPath, os.ModePerm)
	}

	f, err := os.Create(storePath)
	check(err)

	defer f.Close()

	n3, err := f.WriteString(storeJsonStr)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	return resp.AccessToken

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRemoteAccessToken(appid, appSecret string) *weixinAccessTokenResp {

	//config := config.GetWeixinConfig()
	//fmt.Printf("weixin config:%#v", config)

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appSecret)

	response, err := http.Get(url)

	if err != nil {
		//错误处理
		return nil
	} else {
		if response.StatusCode == 200 {

			defer response.Body.Close()

			s, _ := io.ReadAll(response.Body)

			res := weixinAccessTokenResp{}

			/**
			{
			errcode: 40164,
			errmsg: "invalid ip 120.235.19.138 ipv6 ::ffff:120.235.19.138, not in whitelist rid: 6235ac2a-7959a184-75304f31"
			}
			*/

			if json.Unmarshal(s, &res) == nil {
				if res.AccessToken == "" {
					fmt.Println(string(s))
				} else {
					fmt.Println("access token json.Unmarshal success ...")
					return &res
				}
			}

		}
	}

	return nil

}
