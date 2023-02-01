package common

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
https://www.cnblogs.com/xuweiqiang/p/15560119.html

golang中获取公网ip、查看内网ip、检测ip类型、校验ip区间、ip地址string和int转换、根据ip判断地区国家运营商等
https://www.itdaan.com/blog/2017/09/22/3d8db64be58b1cc54e627f1cb66b068b.html
*/

/*
端口检测
*/
func ScanPort(protocol string, hostname string, port int) bool {
	fmt.Printf("scanning port %d \n", port)
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// 传入查询的端口号
// 返回端口号对应的进程PID，若没有找到相关进程，返回-1
func GetPidWithPort(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %d", portNumber)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			res = -1
		} else {
			res = pid
		}
	}
	return res
}
