package common

import (
	"github.com/shirou/gopsutil/host"
)

func SystemBootTime() uint64 {
	timestamp, _ := host.BootTime()
	//t := time.Unix(int64(timestamp), 0)
	//fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
	return timestamp
}
