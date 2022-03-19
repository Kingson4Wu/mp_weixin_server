package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetExtranetIp() string {
	var realIP string
	// Get real IP
	resp, err := http.Get("http://ifconfig.me") // the ip discover service, choose a nearby one
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		realIP = string(body)
	} else {
		log.Println("Fail to get ip from ifconfig.me")
	}
	return realIP
}
