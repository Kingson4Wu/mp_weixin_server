package ip

import (
	"io"
	"log"
	"net/http"
)

func GetExtranetIp() string {
	var realIP string
	// Get real IP
	resp, err := http.Get("http://ifconfig.me") // the ip discover service, choose a nearby one
	defer resp.Body.Close()
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		realIP = string(body)
	} else {
		log.Println("Fail to get ip from ifconfig.me")
	}
	return realIP
}
