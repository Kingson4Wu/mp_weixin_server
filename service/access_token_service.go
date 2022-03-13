package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/kingson4wu/weixin-app/config"
)

//https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxda7a1cb0644cb4cd&secret=43af4a789c581f80b6a6df511ef44d2d
//配置文件TODO

type weinxinAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type accessTokenStore struct {
	AccessToken       string `json:"access_token"`
	ExpireTimeSeconds int    `json:"expire_time_seconds"`
}

func Exists(path string) bool {

	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GetAccessToken() string {

	storePath := "./work/access_token_store.json"
	fmt.Printf("config.wexin: %s\n", "====")

	if Exists(storePath) {
		dat, err := ioutil.ReadFile(storePath)
		check(err)
		tokenStore := accessTokenStore{}
		json.Unmarshal(dat, &tokenStore)

		if tokenStore.ExpireTimeSeconds > int(time.Now().Unix()) {
			fmt.Printf("read accessToken from file success\n")
			return tokenStore.AccessToken
		}
	}

	fmt.Printf("read accessToken from file remote\n")
	resp := getRemoteAccessToken()

	store := accessTokenStore{AccessToken: resp.AccessToken, ExpireTimeSeconds: int(time.Now().Unix()) + resp.ExpiresIn}

	storeJson, _ := json.Marshal(store)
	storeJsonStr := string(storeJson)

	fmt.Printf("config.storeJsonStr: %s\n", storeJsonStr)

	if !Exists("./work") {
		os.Mkdir("./work", os.ModePerm)
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

func getRemoteAccessToken() *weinxinAccessTokenResp {

	config := config.GetWeixinConfig()
	//fmt.Printf("weixin config:%#v", config)

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", config.AppId, config.AppSecret)

	response, err := http.Get(url)

	if err != nil {
		//错误处理
		return nil
	} else {
		if response.StatusCode == 200 {

			defer response.Body.Close()

			s, _ := ioutil.ReadAll(response.Body)

			res := weinxinAccessTokenResp{}
			json.Unmarshal(s, &res)

			if res.AccessToken == "" {
				fmt.Println(s)
			} else {
				return &res
			}

		}
	}

	return nil

}
