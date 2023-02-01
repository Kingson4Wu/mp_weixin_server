package date

import "time"

/**
Golang时间类型通过自带的 Format 方法进行格式化。

需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go语言的诞生时间 2006-01-02 15:04:05 -0700 MST。

https://zhuanlan.zhihu.com/p/145009400
*/

func GetDayRange(day time.Time) (*time.Time, *time.Time) {

	//date := time.Now().AddDate(0, 0, -1).Local().Format("2006-01-02")
	date := day.Local().Format("2006-01-02")

	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return &startTime, &end
}
