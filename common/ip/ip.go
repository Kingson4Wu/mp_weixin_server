package ip

import (
	"fmt"
	"io"
	"log"
	"net"
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

func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
